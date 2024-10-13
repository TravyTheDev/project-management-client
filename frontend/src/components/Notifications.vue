<template>
    <RouterLink :to="'/project/' + message.id">
        <p>{{ message.message }}</p>
    </RouterLink>
</template>

<script setup lang="ts">
import { types } from '../../wailsjs/go/models';


interface props {
    message: types.Notification;
}

const props = defineProps<props>()

</script>

<style scoped>

</style>