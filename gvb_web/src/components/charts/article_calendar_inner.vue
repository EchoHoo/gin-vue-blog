<template>
    <div id="article_calendar" style="height: 250px;">

    </div>
</template>

<script setup>

import { onMounted, watch } from "vue";
import * as echarts from "echarts";

const props = defineProps({
    dataList: {
        type: Array,
    },
    isTitle: {
        type: Boolean,
        default: true,
    }
})
onMounted(() => {
    articleCalendar(); // 确保等待这个函数执行完成
});

watch(() => props.dataList, () => {
    articleCalendar();
});
async function articleCalendar() {
    let chartDom = document.getElementById('article_calendar');
    if (!chartDom) {
        return;
    }

    console.log(props.dataList);
    let article_data = props.dataList.map(item => [item.date, item.count]);
    let textColor = "#555555"
    let myChart = echarts.init(chartDom, null, { locale: 'ZH' });
    let title = null
    if (props.isTitle) {
        title = {
            text: '文章日历',
            textStyle: {
                color: textColor,
            },
            padding: [15, 20],
        }
    }
    let option = {
        title: title,
        tooltip: {
            position: 'top',
            formatter: function (params) {
                return params.value[0] + ': ' + params.value[1];
            }
        },
        visualMap: {
            min: 0,
            max: 10,  // Adjust if your data range is different
            type: 'piecewise',
            orient: 'horizontal',
            left: 'center',
            top: 35,
            inRange: {
                color: ['#ebedf0', '#c6e48b', '#7bc96f', '#32af4a', '#1a792c', '#0f5e1e']
            }
        },
        calendar: {
            top: 90,  // 增加顶部的间距
            left: 30,  // 增加左边的间距
            right: 30,  // 增加右边的间距
            botton: 40,  // 增加底部的间距，确保日历不会太靠近容器边缘
            cellSize: [16, 16],  // 设置每个格子的大小为 40x40 像素
            range: '2024',  // 确保年份正确
            itemStyle: {
                borderWidth: 0.5,
                borderColor: '#ccc'  // 适当调整边框颜色以提高可视化效果
            },
            yearLabel: { show: false },
            monthLabel: {  // 根据需要调整月份标签的样式
                show: true,
                nameMap: 'cn'
            },
            dayLabel: {  // 根据需要调整日期标签的样式
                show: true,
                nameMap: 'cn',
                firstDay: 1,  // 设置周一为每周的第一天
                position: 'start'
            }
        },
        series: [{
            type: 'heatmap',
            coordinateSystem: 'calendar',
            data: article_data
        }]
    };
    myChart.setOption(option);
    window.addEventListener("resize", () => myChart.resize());
}

</script>