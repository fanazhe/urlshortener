<template>
  <ul class="bg-white rounded-lg my-4">
    <Shorten v-for="shortenItem in shortenList.slice().reverse()"
    :shortenItem="shortenItem"
    ></Shorten>
  </ul>
</template>
<script setup lang="ts">
import { computed } from 'vue'
import { useStore } from '@/store/index'
import Shorten from "./Shorten.vue";

const store = useStore()
const shortenList = computed(() => store.state.shortens)

</script>
