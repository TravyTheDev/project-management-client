<template>
    <div>
        <div v-for="project in projects">
            <p>{{ project.title }}</p>
            <p>{{ project.description }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { GetAllProjects } from '../../wailsjs/go/main/ProjectsHandler';
import { main } from '../../wailsjs/go/models';

const projects = ref<main.Project[]>([])

const getProjects = async () => {
   const res = await GetAllProjects()
   projects.value = res
}

onMounted(() => {
    getProjects()
})

</script>

<style scoped>

</style>