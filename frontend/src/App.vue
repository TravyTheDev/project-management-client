<script lang="ts" setup>
import { RouterView } from 'vue-router';
import Layout from './components/Layout.vue'
import Main from './pages/Main.vue';
</script>

<template>
  <Layout class="bg-backgroundColor h-[100vh] text-fontColor">
    <RouterView />
  </Layout>
</template>

<style>
</style>
