<template>
  <BaseTemplate>
    <template v-slot:shorten>
      <Shortener />
      <ShortenList />
    </template>
    <template v-slot:default>
      <FAQ />
    </template>
  </BaseTemplate>
  <About />
  <Statistics />
  <Login />
</template>
<script setup lang="ts">
import BaseTemplate from '@/components/BaseTemplate.vue'
import Shortener from '@/components/Shortener.vue';
import ShortenList from '@/components/ShortenList.vue';
import FAQ from '@/components/FAQ.vue';
import About from '@/components/About.vue';
import Statistics from '@/components/Statistics.vue';
import Login from '@/components/Login.vue';
</script>
