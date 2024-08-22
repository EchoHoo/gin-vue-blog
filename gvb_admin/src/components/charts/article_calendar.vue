<template>
    <ArticleCalendarInner :is-title="props.isTitle" :data-list="data.data_list"></ArticleCalendarInner>
</template>

<script setup>

const data = reactive({
    data_list:[]
})
import ArticleCalendarInner from "@/components/charts/article_calendar_inner.vue"
import { getArticleCalendarApi } from "@/api/article_api";
import { reactive } from "vue";

const props = defineProps({
    isTitle:{
        type:Boolean,
        default:true,
    }
})
async function getData(){
  let res = await getArticleCalendarApi();

  data.data_list = res.data;
}
getData()
</script>