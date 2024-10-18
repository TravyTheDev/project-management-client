<template>
    <div id="chart">
        <VueApexCharts type="rangeBar" height="350" :options="chartOptions" :series="data.series"></VueApexCharts>
    </div>
</template>

<script setup lang="ts">
import VueApexCharts from 'vue3-apexcharts'
import { computed } from 'vue';
import { todoStatus, urgencyStatus } from '../consts';
import { types } from '../../wailsjs/go/models';

const props = defineProps({
    chartData: Array
})

const data = computed(() => {
    return {
          series: [
    {
        data: 
        makeChartData()
    }
]
    }
  
})

const chartOptions = {
    chart: {
        height: 350,
        type: 'rangeBar'
    },
    plotOptions: {
        bar: {
            horizontal: true
        }
    },
    xaxis: {
        type: 'datetime'
    }
}

const makeChartData = () => {
    //@ts-ignore
  return props.chartData.map((project) => {
    return {
            //@ts-ignore
      x: project.title,
      y: [
            //@ts-ignore
      project.startDate ? new Date(project.startDate).getTime() : new Date().getTime(),
          //@ts-ignore
      project.endDate ? new Date(project.endDate).getTime() : new Date().getTime() + 1
      ],
          //@ts-ignore
      fillColor: todoStatus[project.status].color
    }
  })
}

</script>

<style scoped></style>