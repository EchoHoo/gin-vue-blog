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
          <ArticleCalendar :is-title="true"></ArticleCalendar>
        </div>
        <div class="week_data">
          <WeekChart></WeekChart>
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

import { reactive } from "vue";
import WeekChart from "@/components/charts/week_chart.vue"
import ArticleCalendar from "@/components/charts/article_calendar.vue";
import { getDataSumApi } from "@/api/data_api";
import { message } from "ant-design-vue";
async function getData() {
  let res = await getDataSumApi()
  if (res.code) {
    message.error(res.msg)
    return
  }
  data.sum_data_list[0].value = res.data.user_count;
  data.sum_data_list[1].value = res.data.article_count;
  data.sum_data_list[2].value = res.data.message_count;
  data.sum_data_list[3].value = res.data.now_sign_count
  data.sum_data_list[4].value = res.data.now_login_count
}

getData()
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
  ],
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
  justify-content: center; /* 水平居中对齐 */
  
  .left {
    justify-content: center;
    width: 1780px;
    margin-right: 20px;

  }
  // .right {
  //   min-height: 400px; /* 设置一个合适的最小高度 */
  //   width: calc(100% - 780px);
  // }
  .week_data{
    margin-top: 30px;
    background-color: var(--card_bg);
    padding: 10px 20px;
    border: 1px solid var(--card_border);
    border-radius: 5px;
  }
  .article_calendar {

    background-color: var(--card_bg);
    padding: 10px 20px;
    border: 1px solid var(--card_border);
    border-radius: 5px;
  }
}

</style>