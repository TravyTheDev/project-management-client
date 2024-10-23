<template>
    <div class="text-fontColor ml-2">
        <div>
            <RouterLink to="/main">Projects</RouterLink>
            <div class="ml-2 flex flex-col w-[70%]">
                <RouterLink to="/main" class="hover:bg-slate-300">Status</RouterLink>
                <RouterLink to="/urgency" class="hover:bg-slate-300">Urgency</RouterLink>
            </div>
        </div>
        <RouterLink to="/timeline">Timeline</RouterLink>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const active = ref("main")

const setActive = (val:string) => {
    active.value = val
}

</script>

<style scoped>
.active {
    border: 1px solid rgb(203 213 225);
}
</style>