<template>
    <div class="overflow-hidden">
        <Navbar v-if="loginUser?.id != 0" />
        <div v-if="loginUser?.id != 0 && !isHideLayout" class="grid grid-cols-6">
            <div class="col-span-1">
                <SideBar v-if="loginUser?.id != 0"/>
            </div>
            <div class="col-span-5">
                <slot />
            </div>
        </div>
        <div v-else >
            <slot />
        </div>
    </div>
</template>

<script setup lang="ts">
import { projects } from '../../wailsjs/go/models';
import Navbar from './Navbar.vue';
import SideBar from './SideBar.vue';
import { inject, ref, Ref } from 'vue';

const loginUser = inject<Ref<projects.User | undefined>>("loginUser")
console.log(loginUser)

const isHideLayout = ref(false)

const hideLayout = () => {
    isHideLayout.value = true
}
</script>

<style scoped>

</style>