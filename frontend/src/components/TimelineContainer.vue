<template>
    <div class="ganttChartArea h-full overflow-y-auto pb-10">
        <div class="flex justify-start gap-1 w-[450px]">
            <VueDatePicker class="datePicker" placeholder="日付を選択" locale="ja" v-model="month" :format="format"
                :enable-time-picker="false" :day-class="getDayClass" week-start="0" :month-change-on-scroll="false"
                hide-offset-dates no-today :disabled="disabledFlag" @update:model-value="chartFilter" monthPicker />
            <div class="flex gap-4">
                <label @change="changeScale" class="changeScaleBtn">
                    <input type="radio" name="changeScaleBtn" value="year" :disabled="disabledFlag">
                    <p>Year</p>
                </label>
                <label @change="changeScale" class="changeScaleBtn">
                    <input type="radio" name="changeScaleBtn" value="month" checked :disabled="disabledFlag">
                    <p>Month</p>
                </label>
                <label @change="changeScale" class="changeScaleBtn">
                    <input type="radio" name="changeScaleBtn" value="week" :disabled="disabledFlag">
                    <p>Week</p>
                </label>
            </div>
        </div>
        <div class="gantt">
            <div class="chart" :style="{
                height: ganttHeight + 'px',
            }">
                <canvas id="myChart"></canvas>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'
import {
    Chart,
    registerables,
    type ChartConfigurationCustomTypesPerDataset,
    type ChartConfiguration,
    type TooltipItem,
} from 'chart.js';
import 'chartjs-adapter-date-fns';
import { addDays, isEqual, set } from 'date-fns';

const props = defineProps({
    chartData: Array
})

type TodayLine = {
    id: string,
    afterDatasetDraw: (chart: Chart) => void;
};

Chart.register(...registerables)

const myChartDataChangeFlag = ref<boolean>(false)
let myChart = ref()
const thisYear = ref<number>(new Date().getFullYear())
const thisMonth = ref<number>(new Date().getMonth() + 1)
const thisMonthFormatted = ref<string>(('0' + thisMonth.value).slice(-2))
const disabledFlag = ref<boolean>(false)

const getDayClass = (date: any, _internalDate: any) => {
    if (isEqual(date, addDays(set(new Date(), { hours: 0, minutes: 0, seconds: 0, milliseconds: 0 }), 1)))
        return 'marked-cell';
    return '';
};

const todayLine: TodayLine = {
    id: 'todayLine',
    afterDatasetDraw(chart) {

        const { ctx, data, scales: { x, y }, chartArea } = chart;
        let canvas = document.getElementById("myChart") as HTMLCanvasElement;
        let ctxSec = canvas.getContext('2d')

        const now = Date.now(); // UNIXタイムスタンプを取得する
        const xPixel = x.getPixelForValue(now); // UNIXタイムスタンプを元にx座標を計算する

        ctx.save();

        if (xPixel >= chartArea.left && xPixel <= chartArea.right) {
            //ここで縦線をひいてます。
            ctx.beginPath();
            ctx.lineWidth = 2;
            ctx.strokeStyle = 'rgba(255, 26, 104, 1)';
            ctx.setLineDash([4, 4]);
            ctx.moveTo(xPixel, chartArea.top);
            ctx.lineTo(xPixel, chartArea.bottom);
            ctx.stroke();
            ctx.restore();
            ctx.setLineDash([]);

            //ここで下向きの三角形を作ってます。
            ctx.beginPath();
            ctx.lineWidth = 1;
            ctx.strokeStyle = 'rgba(255, 26, 104, 1)';
            ctx.fillStyle = 'rgba(255, 26, 104, 1)';
            ctx.moveTo(xPixel, chartArea.top + 3)
            ctx.lineTo(xPixel - 6, chartArea.top - 6);
            ctx.lineTo(xPixel + 6, chartArea.top - 6);
            ctx.closePath();
            ctx.stroke();
            ctx.fill();
            ctx.restore();

            // Todayという文字を描画します。
            ctx.font = 'bold 12px sans-serif';
            ctx.fillStyle = 'rgba(255, 26, 104, 1)';
            ctx.textAlign = 'center';
            if (ctxSec) ctxSec.fillText("Today", xPixel, chartArea.bottom + 15);
        }
    }
}

const myChartData: ChartConfiguration<"bar", { x: [string, string]; y: string; z: string; }[], unknown> | ChartConfigurationCustomTypesPerDataset<"bar", { x: [string, string]; y: string; z: string; }[], unknown> = {
    type: 'bar',
    data: {
        //@ts-ignore
        datasets: props.chartData,
    },
    options: {
        responsive: true,
        maintainAspectRatio: false,
        indexAxis: "y",
        layout: {
            padding: {
                left: 10,
                bottom: 20
            }
        },
        scales: {
            x: {
                offset: false,
                position: "top",
                min: `${thisYear.value}-${thisMonthFormatted.value}-01`,
                type: 'time',
                time: {
                    unit: 'month',
                    displayFormats: {
                        year: 'yyyy',
                        month: 'yyyy/MM',
                        week: 'yyyy/MM/dd',
                    }
                },
                ticks: {
                    align: 'center',
                },
                grid: {
                    drawTicks: true,
                    drawOnChartArea: true
                }
            },
            y: {
                beginAtZero: true,
                ticks: {
                    font: {
                        weight: 500
                    }
                },
            }
        },
        plugins: {
            legend: {
                display: false,
                position: 'bottom',
                labels: {
                    color: '#000'
                }
            },
            tooltip: {
                displayColors: false,
                yAlign: 'bottom',
                enabled: true,
                callbacks: {
                    label: () => {
                        return '';
                    },
                    title: (ctx: TooltipItem<'bar'>[]) => {
                        const parsed = JSON.parse(ctx[0].formattedValue)
                        const startDate = new Date(parsed[0]);
                        const endDate = new Date(parsed[1]);
                        const formattedStartDate = startDate.toLocaleString([], {
                            year: 'numeric',
                            month: 'short',
                            day: 'numeric',
                        })
                        const formattedEndDate = endDate.toLocaleString([], {
                            year: 'numeric',
                            month: 'short',
                            day: 'numeric',
                        })
                        return `Duration : ${formattedStartDate} - ${formattedEndDate}`
                    }
                }
            },
        }
    },
    plugins: [todayLine]
}

const chartFilter = (date: { year: number; month: number; }) => {
    if (date) {
        const year = date.year
        const month = date.month + 1
        let startDate;

        if (month.toString().length < 2) {
            startDate = `${year}-0${month}`
        } else {
            startDate = `${year}-${month}`
        }

        if (myChartData.options && myChartData.options.scales && myChartData.options.scales.x) myChartData.options.scales.x.min = startDate;
        myChartDataChangeFlag.value = !myChartDataChangeFlag.value
    } else {
        return
    }
}


const changeScale = (e: Event) => {
    if (e.target instanceof HTMLInputElement) {
        const scale = e.target.value;
        // @ts-ignore
        if (myChartData.options && myChartData.options.scales && myChartData.options.scales.x) myChartData.options.scales.x.time.unit = scale;
    }
    myChartDataChangeFlag.value = !myChartDataChangeFlag.value
}

const month = ref({
    month: new Date().getMonth(),
    year: new Date().getFullYear(),
    day: new Date().getDate()
});

const format = (date: Date) => {
    const month = date.getMonth() + 1;
    const year = date.getFullYear();

    if (month.toString().length < 2) {
        return `${year}/0${month}`;
    } else {
        return `${year}/${month}`;
    }

}

const ganttHeight = ref<number>();

const ganttHeightAdjustment = () => {
    ganttHeight.value = myChartData.data.datasets.length * 65
}

onMounted(() => {
    ganttHeightAdjustment()

    disabledFlag.value = !disabledFlag.value

    let ctx: HTMLCanvasElement | null = document.getElementById("myChart") as HTMLCanvasElement;
    if (ctx) myChart.value = new Chart(ctx, myChartData);

    setTimeout(() => {
        disabledFlag.value = !disabledFlag.value
    }, 1000);

})

watch((myChartDataChangeFlag), () => {

    disabledFlag.value = !disabledFlag.value

    myChart.value.destroy();

    let ctx: HTMLCanvasElement | null = document.getElementById("myChart") as HTMLCanvasElement;
    if (ctx) myChart.value = new Chart(ctx, myChartData);

    setTimeout(() => {
        disabledFlag.value = !disabledFlag.value
    }, 1000);

})
</script>

<style scoped></style>