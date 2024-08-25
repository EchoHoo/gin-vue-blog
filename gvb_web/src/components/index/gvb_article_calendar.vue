<template>
    <GVBCard class="article_calendar_card" title="文章日历">
        <GVBArticleCalendar :is-title="false"></GVBArticleCalendar>
    </GVBCard>
</template>


<script setup>
import GVBCard from "@/components/gvb_card.vue"
import GVBArticleCalendar from "@/components/charts/article_calendar.vue"
</script>


<style lang="scss"></style>
