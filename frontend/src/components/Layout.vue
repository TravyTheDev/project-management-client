<template>
    <div class="overflow-hidden">
        <Navbar />
        <div class="grid grid-cols-6">
            <div class="col-span-1">
                <SideBar/>
            </div>
            <div class="col-span-5">
                <RouterView :key="route.path" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import Navbar from './Navbar.vue';
import SideBar from './SideBar.vue';

const route = useRoute()

</script>

<style scoped>

</style>