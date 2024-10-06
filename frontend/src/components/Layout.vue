<template>
    <div>
        <Navbar />
        <div class="grid grid-cols-6">
            <div class="col-span-1">
                <SideBar />
            </div>
            <div class="col-span-5">
                <slot />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import Navbar from './Navbar.vue';
import SideBar from './SideBar.vue';


</script>

<style scoped>

</style>