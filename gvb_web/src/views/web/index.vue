<template>
    <div class="base_view index_view">
        <GVBNav></GVBNav>
        <GVBBanner></GVBBanner>

        <div class="gvb_base_container">
            <div class="gvb_inner_container gvb_index_main">
                <div class="left">
                    <GVBNews></GVBNews>
                    <GVBArticleCalendarCard style="margin-top: 20px"></GVBArticleCalendarCard>
                </div>
                <div class="right">
                    <GVBAdvertCard></GVBAdvertCard>
                    <GVBTagCloud style="margin-top: 20px"></GVBTagCloud>
                    <GVBSiteCard style="margin-top: 20px"></GVBSiteCard>
                    <GVBFadeBackCard style="margin-top: 20px"></GVBFadeBackCard>
                </div>
            </div>
        </div>
        <div class="gvb_footer">

        </div>
    </div>
</template>

<script setup>
import GVBNav from "@/components/gvb_nav.vue"
import GVBBanner from "@/components/gvb_banner.vue"
import GVBNews from "@/components/index/gvb_news.vue"
import GVBArticleCalendarCard from "@/components/index/gvb_article_calendar.vue"
import GVBTagCloud from "@/components/index/gvb_tag_card.vue"
import GVBSiteCard from "@/components/index/gvb_site_card.vue"
import GVBAdvertCard from "@/components/index/gvb_advertis_card.vue"
import GVBFadeBackCard from "@/components/index/gvb_fade_back_card.vue"
</script>


<style lang="scss">
.base_view {
    background-color: var(--card_bg);


    .gvb_base_container {
        display: flex;
        justify-content: center;

        .gvb_inner_container {
            background-color: var(--card_bg);
            min-height: 1000px;
            width: 1200px;
            margin-top: 20px;
        }
    }
}

.index_view {
    .gvb_inner_container {
        display: flex;
        justify-content: space-between;
    }

    .left {
        width: calc(100% - 416px);
    }

    .right {
        width: 396px;

    }

}
</style>