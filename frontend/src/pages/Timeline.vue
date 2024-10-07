<template>
    <div>
        timeline
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { GetAllProjects } from '../../wailsjs/go/projects/ProjectsHandler';
import { projects } from '../../wailsjs/go/models';

const projectsData = ref<projects.Project[]>([])

const getProjects = async () => {
   const res = await GetAllProjects()
   projectsData.value = res
}

onMounted(() => {
    getProjects()
})
</script>

<style scoped>

</style>