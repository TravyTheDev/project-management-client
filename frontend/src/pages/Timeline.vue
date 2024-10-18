<template>
  <div class="h-[87vh]">
    <TimelineContainer v-if="dataCharts" :chart-data="dataCharts" />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { GetAllProjects } from '../../wailsjs/go/projects/ProjectsHandler';
import { types } from '../../wailsjs/go/models';
import TimelineContainer from '../components/TimelineContainer.vue';
import { todoStatus, urgencyStatus } from '../consts';

const projectsData = ref<types.Project[]>([])
const dataCharts = ref()

const getProjects = async () => {
  const res = await GetAllProjects()
  projectsData.value = res
  setChartData(res)
}

const setChartData = (projects: types.Project[]) => {
  const chartData = projects.map((project) => {
    return {
      label: project.title,
      backgroundColor: [
        todoStatus[project.status].color,
      ],
      borderColor: [
        urgencyStatus[project.urgency].color,
      ],
      borderWidth: 1,
      barPercentage: 1,
      categoryPercentage: 1,
      borderSkipped: false,
      borderRadius: 5,
      data: [
        { x: [project.startDate.toString(), project.endDate.toString()], y: project.title, z: project.description }
      ],
    }
  })
  dataCharts.value = chartData
}

onMounted(() => {
  getProjects()
})

</script>

<style scoped></style>