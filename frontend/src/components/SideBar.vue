<template>
    <div class="text-fontColor ml-2">
        <div class="hover:cursor-pointer">
            <h3 @click="setActive('/')">Projects</h3>
            <div class="ml-2" v-if="active == '/'">
                <h4>todo</h4>
                <h4>in progress</h4>
                <h4>done</h4>
            </div>
        </div>
        <h3 class="hover:cursor-pointer" @click="setActive('/timeline')">Timeline</h3>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()


const active = ref("/")

const setActive = (val:string) => {
    active.value = val
    router.push(`${val}`)
}
</script>

<style scoped>

</style>