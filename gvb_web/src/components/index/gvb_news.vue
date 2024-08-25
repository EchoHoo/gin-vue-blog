<template>
    <GVBCard class="news_card" title="今日热搜" link_name="查看更多" link="/news">
        <span class="index">123</span>
        <a href="">新闻巴拉巴拉</a>
        <span class="num">68.7万</span>
    </GVBCard>
</template>


<script setup>
import GVBCard from "@/components/gvb_card.vue"
</script>


<style lang="scss">
.news_card {
    .body {
        display: flex;
        justify-content: space-between;
        align-items: center;

        .index {
            width: 5%;
        }

        a {
            width: 85%;
        }

        .num {
            width: 10%;
            text-align: right;
        }
    }

}
</style>
