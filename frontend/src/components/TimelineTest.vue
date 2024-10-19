<template>
    <div id="chart" class="overflow-y-auto h-[87vh]">
        <VueApexCharts type="rangeBar" :options="chartOptions" :series="data.series"></VueApexCharts>
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
const chartOptions = computed(() => {
    const height = props.chartData ? (props.chartData?.length * 75) + 'px' : 0 
    return {
        chart: {
            height: height,
            type: 'rangeBar',
            zoom: {
                enabled: false
            }
        },
        plotOptions: {
            bar: {
                horizontal: true,
                // isDumbbell: true,
                barHeight: "50%",
                columnWidth: '50%',
            }
        },
        xaxis: {
            type: 'datetime',
            position: 'top',
        },
        grid: {
            xaxis: {
                lines: {
                    show: true
                },
            },
            yaxis: {
                lines: {
                    show: true
                }
            }
        },
        annotations: {
            xaxis: [{
                x: new Date().getTime(),
                strokeDashArray: 0,
                borderColor: '#775DD0',
                label: {
                  borderColor: '#775DD0',
                  style: {
                    color: '#fff',
                    background: '#775DD0',
                  },
                  text: 'Today',
                }
              }]
        },
    }
})

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