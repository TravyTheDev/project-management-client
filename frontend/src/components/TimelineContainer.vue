<template>
    <div id="chart" class="overflow-y-auto h-[87vh]">
        <VueApexCharts type="rangeBar" :options="chartOptions" :series="data.series"></VueApexCharts>
    </div>
</template>

<script setup lang="ts">
import VueApexCharts from 'vue3-apexcharts'
import { computed, ref } from 'vue';
import { todoStatus } from '../consts';
import { types } from '../../wailsjs/go/models';

interface props {
    chartData: types.Project[]
}

const props = defineProps<props>()
const chartRef = ref()

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
            },
            toolbar: {
                show: false,
            },
        },
        plotOptions: {
            bar: {
                horizontal: true,
                barHeight: "50%",
                columnWidth: '50%',
            }
        },
        xaxis: {
            type: 'datetime',
            position: 'top',
            labels: {
                style: {
                    fontSize: '0.75rem',
                }
            }
        },
        yaxis: {
            labels: {
                style: {
                    fontSize: '1rem'
                }
            }
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
                borderColor: '#ef4444',
                label: {
                    borderColor: '#ef4444',
                    style: {
                        color: '#fff',
                        background: '#ef4444',
                    },
                    text: 'Today',
                }
            }]
        },
    }
})

const makeChartData = () => {
    return props.chartData.map((project) => {
        return {
            x: project.title,
            y: [
                project.startDate ? new Date(project.startDate).getTime() : new Date().getTime(),
                project.endDate ? new Date(project.endDate).getTime() : new Date().getTime() + 1
            ],
            fillColor: todoStatus[project.status].color,
        }
    })
}

</script>

<style scoped>
.sticky {
    position: fixed;
}
</style>