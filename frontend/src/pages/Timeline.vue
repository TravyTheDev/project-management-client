<template>
  <div class="h-[87vh]">
    <TimelineContainer :chart-data="projectsData" />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { GetAllProjects } from '../../wailsjs/go/projects/ProjectsHandler';
import { types } from '../../wailsjs/go/models';
import TimelineContainer from '../components/TimelineContainer.vue';

const projectsData = ref<types.Project[]>([])
const dataCharts = ref()

const getProjects = async () => {
  const res = await GetAllProjects()
  projectsData.value = res
}

onMounted(() => {
  getProjects()
})

</script>

<style scoped></style>