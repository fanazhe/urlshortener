<template>
  <transition name="fade">
    <div v-show="showModal" class="fixed inset-0 z-30">
      <div
        v-show="showModal"
        class="bg-filter bg-white opacity-75 fixed inset-0 w-full h-full z-20"
        @click="toggleAboutModal"
      ></div>
      <main class="flex flex-col items-center justify-center h-full w-full">
        <transition name="fade-up-down">
          <div v-show="showModal" class="modal-wrapper flex items-center z-30">
            <div
              class="modal max-w-md mx-auto xl:max-w-5xl lg:max-w-5xl md:max-w-2xl bg-white max-h-screen shadow-lg flex-row rounded relative"
            >
              <div class="modal-header p-5 bg-gray-700 text-gray-800 rounded-t">
                <h5 class="text-white text-2xl uppercase">About URL Shortener</h5>
              </div>
              <div class="modal-body p-5 w-full h-full overflow-y-auto leading-6">
                <p class="text-justify py-2">
                  This URL Shortener is created to do a sloppy copy of
                  <a
                    href="https://bitly.com"
                    target="_blank"
                  >bit.ly</a>, which is one of the largest platform in url shortening business.
                </p>
                <p
                  class="text-justify py-2"
                >Bitly is a link management platform that helps people, brands and businesses make their content connect and ignite action with every link.</p>
                <p class="text-justify py-2">
                  This site is a part of
                  <a href="https://bytebase.com" target="_blank">Bytebase</a>'s
                  <a
                    href="https://github.com/bytebase/interview/tree/master/104_url_shortener"
                    target="_blank"
                  >interview</a>, which is a very interesting and challenging journey.
                </p>
              </div>
              <div class="modal-footer py-3 px-5 border0-t text-right">
                <button
                  class="bg-green-500 w-28 px-5 py-2 text-white rounded-md"
                  @click="toggleAboutModal"
                >OK</button>
              </div>
            </div>
          </div>
        </transition>
      </main>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { Mutation, useStore } from '@/store/index'

const store = useStore()
const showModal = computed(() => store.state.showAbout)
function toggleAboutModal() {
  store.commit(Mutation.TOGGLE_ABOUT, undefined);
}
</script>
