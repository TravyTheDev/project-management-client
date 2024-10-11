<template>
    <RouterLink :to="'/project/' + message.id">
        <p>{{ message.message }}</p>
    </RouterLink>
</template>

<script setup lang="ts">
import { main } from '../../wailsjs/go/models';


interface props {
    message: main.Notification;
}

const props = defineProps<props>()

</script>

<style scoped>

</style>