<template>
  <div>
    <div v-for="item in mydatas">
      <div class>
        <a
          href="#"
          class="my-1 p-4 bg-gray-200 hover:bg-gray-300 no-underline !text-black flex justify-between rounded"
          @click.prevent="active = (active == item.id) ? 0 : item.id"
        >
          <strong>{{ item.id }}, {{ item.name }}</strong>
          <span class="text-gray-300" v-show="active != item.id">&#10133;</span>
          <span class="text-gray-300" v-show="active == item.id">&#10134;</span>
        </a>
      </div>
      <div class="overflow-hidden">
        <transition
          enter-active-class="duration-200 ease-in"
          enter-from-class="opacity-0"
          enter-to-class="opacity-100"
          leave-active-class="duration-200 ease-in"
          leave-from-class="opacity-0"
          leave-to-class="opacity-100"
        >
          <div class="transition-all p-5 whitespace-pre-wrap text-black" v-show="active == item.id">
            <slot />
            {{ item.text }}
          </div>
        </transition>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const BACKEND_URL = (import.meta.env.VITE_BACKEND_URL.length > 0) ? import.meta.env.VITE_BACKEND_URL : window.location.origin;
const active = ref(1)
const mydatas = ref([
  {
    id: 1, name: 'What is a URL Shortener?', text: `A URL shortener, also known as a link shortener, seems like a simple tool, but it is a service that can have a dramatic impact on your marketing efforts.

Link shorteners work by transforming any long URL into a shorter, more readable link. When a user clicks the shortened version, they’re automatically forwarded to the destination URL.

Think of a short URL as a more descriptive and memorable nickname for your long webpage address. You can, for example, use a short URL like ${BACKEND_URL}/IAMLUKY so people will have a good idea about where your link will lead before they click it.

If you’re contributing content to the online world, you need a URL shortener.

Make your URLs stand out with our easy to use free link shortener above.`
  },
  {
    id: 2, name: 'Benefits of a Short URL', text: `How many people can even remember a long web address, especially if it has tons of characters and symbols? A short URL can make your link more memorable. Not only does it allow people to easily recall and share your link with others, it can also dramatically improve traffic to your content.

On a more practical side, a short URL is also easier to incorporate into your collateral – whether you’re looking to engage with your customers offline or online.

Bitly is the best URL shortener for everyone, from influencers to small brands to large enterprises, who are looking for a simple way to create, track and manage their links.`
  },
  {
    id: 3, name: 'Does it cost money to use URL Shortener?', text: `No. This site gives you unlimited access to the ${BACKEND_URL} link shortener. You can also upgrade to a paid plan if you need more features.`
  }
])
</script>
