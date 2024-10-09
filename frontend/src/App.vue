<script lang="ts" setup>
import { RouterView, useRoute } from 'vue-router';
import Layout from './components/Layout.vue'

const route = useRoute()
</script>

<template>
  <Layout class="bg-backgroundColor h-[100vh] text-fontColor">
    <RouterView :key="route.path" />
  </Layout>
</template>

<style>
</style>
