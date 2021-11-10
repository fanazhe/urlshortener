<template v-slot:shortener>
  <form
    class="relative text-gray-700 flex flex-wrap mt-3"
    spellcheck="false"
    v-on:submit.prevent="submit"
  >
    <input
      class="w-full md:w-10/12 h-14 px-4 mr-0 md:mr-2 text-lg placeholder-gray-400 border rounded-lg focus:ring-2 focus:ring-inset focus:ring-blue-600 selection:bg-blue-300"
      type="text"
      placeholder="Shorten your link"
      v-model="inputUrl"
      @input="onInputChange"
      :class="isCopyMode ? 'text-blue-600' : 'text-black'"
      ref="inputRef"
    />
    <button
      class="w-full h-14 mt-2 md:flex-1 md:mt-0 items-center px-4 font-bold text-white bg-blue-700 rounded-lg hover:bg-blue-600"
      ref="submitRef"
    >{{ isCopyMode ? 'Copy' : 'Shorten' }}</button>
    <transition
      enter-active-class="duration-500 ease-in"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="duration-500 ease-out"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-show="showingAlert"
        class="relative w-full my-2 px-4 py-3 leading-normal rounded-lg"
        :class="isErrorAlert ? 'text-red-700 bg-red-100' : 'text-green-700 bg-green-100'"
        role="alert"
      >
        <span v-if="isErrorAlert" class="absolute inset-y-0 left-0 flex items-center ml-4">
          <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
            <path
              fill-rule="evenodd"
              d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
              clip-rule="evenodd"
            />
          </svg>
        </span>
        <span v-else class="absolute inset-y-0 left-0 flex items-center ml-4">
          <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
            <path
              fill-rule="evenodd"
              d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
              clip-rule="evenodd"
            />
          </svg>
        </span>
        <p class="ml-6">{{ alertMessage }}</p>
      </div>
    </transition>
  </form>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { useStore, Mutation } from '@/store/index'
import ShortenModel from "@/store/models/ShortenModel";

const BACKEND_URL = (import.meta.env.VITE_BACKEND_URL.length > 0) ? import.meta.env.VITE_BACKEND_URL : window.location.origin;
const UrlRegex = new RegExp(/^((https?|ftp):\/\/(?:www\.|(?!www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|(https?|ftp):\/\/(?:www\.|(?!www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,})$/i);
const UrlHasProtocolRegex = new RegExp(/^(https?|ftp):\/\/.*/i);
const SubmitIdleTime = 2 * 1000; // 2 second
const AlertDisplayTime = 4 * 1000; // 4 second

const store = useStore()

function addNewShorten(id: string, url: string, createdAt: number, visits: number) {
  const newShorten = new ShortenModel(id, url, createdAt, visits);
  store.commit(Mutation.NEW_SHORTEN, newShorten);
}

function shouldSubmit(_url: string): { canSubmit: boolean, url: string, tip: string } {
  let url = _url.trim();
  if (url.length == 0) {
    return { canSubmit: false, url, tip: 'Please input a valid url' };
  }
  if (!UrlHasProtocolRegex.test(url)) {
    url = 'http://' + url;
  }
  if (UrlRegex.test(url)) {
    try {
      const { hostname } = new URL(url);
      if (hostname == BACKEND_URL) {
        return { canSubmit: false, url, tip: 'This is already a shortened link' };
      }
    } catch (error) {
      return { canSubmit: false, url, tip: (error as Error).message };
    }
    return { canSubmit: true, url, tip: '' };
  }
  return { canSubmit: false, url, tip: 'This is not a valid url' };
}

export default defineComponent({
  name: "Shortener",
  data() {
    return {
      isCopyMode: false,
      lastSubmitTime: 0,
      inputUrl: '',
      isErrorAlert: false,
      showingAlert: false,
      alertMessage: '',
      alertTimer: 0,
    }
  },
  methods: {
    showAlert(isErrorAlert: boolean, message: string) {
      this.isErrorAlert = isErrorAlert;
      this.alertMessage = message;
      this.showingAlert = true;
      if (this.alertTimer) clearTimeout(this.alertTimer);
      this.alertTimer = setTimeout(() => {
        this.alertTimer = 0;
        this.showingAlert = false;
      }, AlertDisplayTime);
    },
    async submit() {
      if (this.isCopyMode) {
        this.copy();
        return;
      }
      let { canSubmit, url, tip } = shouldSubmit(this.inputUrl);
      if (canSubmit) {
        if (this.lastSubmitTime + SubmitIdleTime < Date.now()) {
          this.showingAlert = false;
          if (this.alertTimer) clearTimeout(this.alertTimer);
          this.lastSubmitTime = Date.now();
          const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ url })
          };
          await fetch(BACKEND_URL + '/shorten', requestOptions).then(response => {
            if (response.status >= 400 && response.status < 600) {
              throw new Error('Bad response from server');
            }
            return response.json();
          }).then(data => {
            this.inputUrl = BACKEND_URL + '/' + data.id;
            this.isCopyMode = true;
            this.showAlert(false, 'Short link generated, click to copy');
            addNewShorten(data.id as string, data.url as string, data.createdat as number, data.visits as number);
          }).then(() => {
            (this.$refs['inputRef'] as HTMLInputElement).select();
          }).catch(error => {
            console.error(error);
            this.showAlert(true, 'Network error!');
          });
        } else {
          this.showAlert(true, `Submit too soon, please wait ${SubmitIdleTime}ms.`);
        }
      } else {
        this.showAlert(true, tip);
      }
    },
    copy() {
      (this.$refs['inputRef'] as HTMLInputElement).select();
      document.execCommand('copy');
      this.showAlert(false, 'Copied to clipboard');
    },
    onInputChange() {
      if (this.isCopyMode) {
        this.isCopyMode = false;
      }
    },
  },
});
</script>
