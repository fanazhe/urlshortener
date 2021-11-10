<template>
  <li
    class="w-full md:h-16 p-5 md:flex border-b-2 last:border-none box-border justify-between items-center"
  >
    <span
      class="text-gray-400 max-w-full block md:max-w-1/2 md:inline overflow-hidden overflow-ellipsis whitespace-nowrap"
      :title="'Generated at ' + new Date(props.shortenItem.CreatedAt * 1000).toString()"
    >{{ props.shortenItem.Url }}</span>
    <span class="align-middle max-w-full block md:max-w-1/2 md:inline">
      <span
        class="overflow-hidden overflow-ellipsis whitespace-nowrap pr-2 max-w-full block md:max-w-1/4 md:inline"
      >
        <a
          :href="BACKEND_URL + '/' + props.shortenItem.Id"
          target="_blank"
          :title="'Shorten URL for ' + props.shortenItem.Url"
          class="!text-blue-600"
        >{{ BACKEND_URL + '/' + props.shortenItem.Id }}</a>
      </span>
      <span class="whitespace-nowrap py-2 block w-full md:inline">
        <transition>
          <button
            class="transition duration-100 w-full md:w-20 h-10 text-sm font-bold items-center rounded-lg hover:drop-shadow-sm"
            :class="copying.isCopying ? 'bg-lime-600 hover:bg-lime-600 text-white hover:text-white' : 'text-blue-400 bg-blue-100 hover:bg-blue-200 hover:text-blue-500'"
            @click="copy"
          >{{ copying.isCopying ? 'Copied!' : 'Copy' }}</button>
        </transition>
      </span>
    </span>
  </li>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import ShortenModel from '@/store/models/ShortenModel'

const BACKEND_URL = (import.meta.env.VITE_BACKEND_URL.length > 0) ? import.meta.env.VITE_BACKEND_URL : window.location.origin;
const props = defineProps<{
  shortenItem: ShortenModel,
}>()
const copying = ref({
  isCopying: false,
  timer: 0,
})
function copy() {
  copying.value.isCopying = true;
  const tempInput = document.createElement('input')
  tempInput.value = BACKEND_URL + '/' + props.shortenItem.Id;
  tempInput.setAttribute('width', '0');
  tempInput.setAttribute('height', '0');
  tempInput.setAttribute('position', 'fixed');
  tempInput.setAttribute('top', '0');
  tempInput.setAttribute('left', '0');
  document.body.append(tempInput);
  tempInput.select();
  document.execCommand('copy');
  document.body.removeChild(tempInput);
  if (copying.value.timer) {
    clearTimeout(copying.value.timer);
  }
  copying.value.timer = setTimeout(() => {
    copying.value.isCopying = false;
  }, 2000);
}
</script>
