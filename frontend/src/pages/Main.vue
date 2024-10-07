<template>
    <div>
        <div v-for="project in projectsData">
            <div @click="viewProject(project.id)">
                <p>{{ project.title }}</p>
                <p>{{ project.description }}</p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { GetAllProjects } from '../../wailsjs/go/projects/ProjectsHandler';
import { projects } from '../../wailsjs/go/models';
import router from '../router';

const projectsData = ref<projects.Project[]>([])

const getProjects = async () => {
   const res = await GetAllProjects()
   projectsData.value = res
}

onMounted(() => {
    getProjects()
})

const viewProject = (id: number) => {
    router.push(`/project/${id}`)
}

</script>

<style scoped>

</style>