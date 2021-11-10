<template>
  <header class="select-none">
    <div class="bg-white w-full shadow">
      <nav class="bg-white" role="navigation">
        <div class="container mx-auto p-4 flex flex-wrap items-center md:flex-no-wrap">
          <div class="ml-1 flex-1 flex items-center">
            <a href="#">
              <img src="@/assets/logo.png" alt="logo" class="float-left w-20 h-20" />
            </a>
            <a href="#">
              <div class="pl-2 text-2xl no-underline font-sans font-bold">URL Shortener</div>
            </a>
          </div>
          <div class="ml-auto hidden">
            <button class="flex items-center px-3 py-2 border rounded" type="button">
              <svg class="h-3 w-3" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <title>Menu</title>
                <path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z" />
              </svg>
            </button>
          </div>
          <div class="w-full md:w-auto md:flex-grow md:flex md:items-center">
            <ul
              class="flex flex-col mt-4 -mx-4 pt-4 border-t md:flex-row md:items-center md:mx-0 md:ml-auto md:mt-0 md:pt-0 md:border-0"
            >
              <li>
                <a
                  class="block px-4 py-1 md:p-2 lg:px-4 hover:text-gray-500"
                  href="#"
                  @click="toggleAboutModal"
                >About</a>
              </li>
              <li>
                <a
                  class="block px-4 py-1 md:p-2 lg:px-4 hover:text-gray-500"
                  href="#"
                  @click="toggleStatsModal"
                >Statistics</a>
              </li>
              <li>
                <a
                  class="block px-4 py-1 md:p-2 lg:px-4 text-purple-200"
                  href="#"
                  @click="toggleLoginModal"
                >Sign-in</a>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </div>
    <div class="md:container px-4 md:mx-auto mt-6 xl:prose md:px-6 md:py-12 grid grid-cols-5">
      <div class="col-start-1 col-end-4 mx-4">
        <h1 class="font-sans text-3xl md:text-5xl font-black">Short links, big results</h1>
        <p
          class="py-8 text-md md:text-xl leading-relaxed tracking-wide text-gray-700"
        >A URL shortener built with powerful tools to help you grow and protect your brand.</p>
      </div>
      <div class="col-start-4 col-end-6 mx-auto">
        <img src="@/assets/brand.png" alt="brand" class="object-scale-down" />
      </div>
    </div>
    <div
      class="github-ribbon"
      style="
        position: absolute;
        right: 0px;
        top: 0px;
        width: 150px;
        height: 150px;
        overflow: hidden;
        z-index: 99999;
      "
    >
      <a
        style="
          display: inline-block;
          width: 200px;
          overflow: hidden;
          padding: 6px 0px;
          text-align: center;
          transform: rotate(45deg);
          text-decoration: none;
          color: rgb(255, 255, 255);
          position: inherit;
          top: 45px;
          right: -40px;
          border-width: 1px 0px;
          border-style: dotted;
          border-color: rgba(255, 255, 255, 0.7);
          font: 700 13px 'Helvetica Neue', Helvetica, Arial, sans-serif;
          box-shadow: rgba(0, 0, 0, 0.5) 0px 2px 3px 0px;
          background-color: rgb(170, 0, 0);
        "
        href="https://github.com/fanazhe/urlshortener"
        target="blank"
      >Fork me on GitHub</a>
    </div>
  </header>
  <main class="flex-grow">
    <div class="bg-gray-700">
      <div class="md:container md:mx-auto p-4 md:px-6 prose-blue">
        <slot name="shorten"></slot>
      </div>
    </div>
    <div class="bg-gray-100 bg-opacity-60">
      <div class="md:container md:mx-auto p-4 md:px-6 prose-blue">
        <slot></slot>
      </div>
    </div>
  </main>
  <footer class="container py-6 mx-auto text-sm text-center text-gray-700">
    <p>
      URL Shortener by
      <a href="https://github.com/fanazhe" class="underline" target="_blank">Fanazhe</a>
      &copy; {{ thisYear }}.
      <template
        v-if="BUILD_DATE"
      >Site built {{ BUILD_DATE.toLocaleDateString() }}.</template>
      <template v-else>Development - VERSION {{ VERSION }}</template>
    </p>
  </footer>
</template>
<script setup lang="ts">
import { useHead } from '@vueuse/head'
import { useStore } from '@/store/index'
import { Mutation, mutations } from '@/store/mutations'

const store = useStore()

const props = defineProps<{
  title?: string
}>()

// Note: these are static. If you need these to be reactive, use ref() or reactive()
useHead({
  bodyAttrs: {
    title: props.title,
  },
})

const VERSION = import.meta.env.VITE_APP_VERSION
const BUILD_DATE = import.meta.env.VITE_APP_BUILD_EPOCH
  ? new Date(Number(import.meta.env.VITE_APP_BUILD_EPOCH))
  : undefined
const thisYear = new Date().getFullYear()
function toggleAboutModal() {
  store.commit(Mutation.TOGGLE_ABOUT, undefined);
}
function toggleStatsModal() {
  store.commit(Mutation.TOGGLE_STATS, undefined);
}
function toggleLoginModal() {
  store.commit(Mutation.TOGGLE_LOGIN, undefined);
}
</script>
