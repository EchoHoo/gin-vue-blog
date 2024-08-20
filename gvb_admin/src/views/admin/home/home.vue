<template>
  <div>
    <div class="gvb_data_preview">
      <div class="preview_item" v-for="(item, index) in data.sum_data_list" :key="index">
        <div class="icon">
          <i :class="'iconfont ' + iconList[index]"></i>
        </div>
        <div class="text">
          <div class="data_title">{{ item.label }}</div>
          <div class="data_sum">{{ item.value }}</div>
        </div>
      </div>
    </div>
    <div class="gvb_data_charts">
      <div class="left">
        <div class="article_calendar">
          <div class="title" style="text-align: center;font-size: 20px;">
            文章日历
          </div>
          <div id="article_calendar" style="height: 400px; width:100%">

          </div>
        </div>
      </div>
      <div class="right">

      </div>
    </div>
  </div>

</template>
<script setup>
const iconList = [
  'icon-yonghutongji',
  'icon-icon_xytjbb',
  'icon-xiaoxitongji',
  'icon-yonghuzhucetongji',
  'icon-denglutongji'
]
import { getArticleCalendarApi } from "@/api/article_api";
import { onMounted, reactive } from "vue";
import * as echarts from "echarts";
onMounted(() => {
  articleCalendar(); // 确保等待这个函数执行完成
});
// async function articleCalendar(){
//   let color = "#000012"
//   let inRangeColor = [
//     "#ebedf0", "#c6e48b", "#7bc96f", "#32af4a",
//     "#1a792c", "#0f5e1e", "#0f491a", "#02340c"
//   ]
//   let borderColor = "#fff"
//   let chart = document.getElementById("article_calendar")
//   if (!chart) {
//     return
//   }

//   let res = await getArticleCalendarApi()
//   console.log("res:",res)

//   let article_data = []
//   let article_change = []
//   let len = res.data.length
//   article_change.push(res.data[0].date)
//   article_change.push(res.data[len - 1].date)

//   for (const item of res.data) {
//     article_data.push([item.date, item.count])
//   }
//   console.log("chart:",chart)
//   console.log("article_data:",article_data)
//   console.log("article_change:",article_change)
//   let myChart = echarts.init(chart, null,  {locale: 'ZH'}), option = {
//     tooltip: {
//       padding: 10,
//       backgroundColor: "#555",
//       borderColor: "#777",
//       borderWidth: 1,
//       formatter: function (e) {
//         e = e.value;
//         return '<div style="font-size: 14px;">' + e[0] + ":" + e[1] + "</div>"
//       }
//     },
//     visualMap: {
//       show: !0,
//       showLabel: !0,
//       categories: [0, 1, 2, 3, 4, 5, 6, 7],
//       // min: 0,
//       // max: 10,
//       calculable: !0,
//       inRange: {
//         symbol: "rect",
//         color: inRangeColor
//       },  // 对应上面的四个值
//       itemWidth: 120,
//       itemHeight: 120,
//       // type:'piecewise',
//       orient: "horizontal",
//       left: "center",
//       bottom: 0,
//       textStyle: {
//         color: color
//       }
//     },
//     calendar: [{
//       top: 25,
//       left: "center",
//       range: article_change,  // 时间范围
//       cellSize: [10, 10],
//       splitLine: {show: !1},
//       itemStyle: {color: "#196127", borderColor: borderColor, borderWidth: 2},
//       yearLabel: {show: !1},
//       monthLabel: {
//         nameMap: "cn",
//         fontSize: 11,
//         color: color
//       },
//       dayLabel: {
//         formatter: "{start}  1st",
//         nameMap: "cn",
//         fontSize: 11,
//         color: color
//       }
//     }],
//     series: [
//       {
//         type: "heatmap",
//         coordinateSystem: "calendar",
//         calendarIndex: 0,
//         data: article_data,
//       }
//     ],
//   };
//   myChart.setOption(option)

//   window.onresize = () => {
//     myChart.resize();
//   }
// }
async function articleCalendar() {
  let chartDom = document.getElementById('article_calendar');
  if (!chartDom) {
    return;
  }
  let res = await getArticleCalendarApi();
  let article_data = res.data.map(item => [item.date, item.count]);

  let myChart = echarts.init(chartDom, null, { locale: 'ZH' });
  let option = {
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
      top: 100,  // 增加顶部的间距
      left: 50,  // 增加左边的间距
      right: 50,  // 增加右边的间距
      bottom: 50,  // 增加底部的间距，确保日历不会太靠近容器边缘
      cellSize: [40, 40],  // 设置每个格子的大小为 40x40 像素
      range: '2024',  // 确保年份正确
      itemStyle: {
        borderWidth: 1,
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

const data = reactive({
  sum_data_list: [
    {
      label: "用户总数",
      value: 11,
    },
    {
      label: "文章总数",
      value: 221,
    },
    {
      label: "消息总数",
      value: 1,
    },
    {
      label: "今日注册",
      value: 2121,
    },
    {
      label: "今日登录",
      value: 23321,
    }
  ]
})
</script>

<style lang="scss">
.gvb_data_preview {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  grid-column-gap: 20px;
}

.preview_item {
  display: flex;

  background-color: var(--card_bg);
  border-radius: 5px;
  border: 1px solid var(--card_border);
  padding: 20px;

  .icon {
    width: 60%;
    display: flex;
    align-items: center;
    justify-content: center;

    i {
      font-size: 45px;
      color: var(--active);
    }
  }

  .text {
    width: 40%;
    font-size: 16px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    .data_sum {
      font-size: 18px;
    }

    .data_title {
      color: var(--text)
    }
  }
}

.gvb_data_charts {
  display: flex;
  margin-top: 20px;

  .left {
    width: 1780px;
    margin-right: 20px;
  }

  .article_calendar {
    background-color: var(--card_bg);
    padding: 10px 20px;
    border: 1px solid var(--card_border);
    border-radius: 5px;
  }

  .right {
    flex: 1;
    // 其他样式
  }
}
</style>