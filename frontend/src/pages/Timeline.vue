<template>
    <div>
        timeline
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { GetAllProjects } from '../../wailsjs/go/projects/ProjectsHandler';
import { types } from '../../wailsjs/go/models';

const projectsData = ref<types.Project[]>([])

// const getProjects = async () => {
//    const res = await GetAllProjects()
//    projectsData.value = res
// }

// onMounted(() => {
//     getProjects()
// })
</script>

<style scoped>

</style>