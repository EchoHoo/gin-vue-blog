
<template>
<div>
  <div class="gvb_data_preview">
    <div class="preview_item" v-for="(item,index) in data.sum_data_list" :key="index">
      <div class="icon">
        <i :class="'iconfont '+iconList[index]"></i>
      </div>
      <div class="text">
        <div class="data_title">{{item.label}}</div>
        <div class="data_sum">{{item.value}}</div>
      </div>
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
import {reactive} from "vue";

const data = reactive({
  sum_data_list:[
    {
      label:"用户总数",
      value:11,
    },
      {
      label:"文章总数",
      value:221,
    },
    {
      label:"消息总数",
      value:1,
    },
    {
      label:"今日注册",
      value:2121,
    },
    {
      label:"今日登录",
      value:23321,
    }
  ]
})
</script>

<style lang="scss">
.gvb_data_preview{
  display: grid;
  grid-template-columns: repeat(5,1fr);
  grid-column-gap: 20px;
}
.preview_item{
  display: flex;
  background-color: white;
  border-radius: 5px;
  border: 1px solid #e2e2e2;
  padding: 20px;
  .icon{
    width: 60%;
    display: flex;
    align-items: center;
    justify-content: center;
    i{
      font-size: 45px;
      color: var(--active);
    }
  }
  .text{
    width: 40%;
    font-size: 16px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    .data_sum{
      font-size:18px;
    }
    .data_title{
      color: var(--text)
    }
  }
}
</style>