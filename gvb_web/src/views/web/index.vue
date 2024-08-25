<template>
    <div class="base_view index_view">
        <GVBNav></GVBNav>
        <GVBBanner></GVBBanner>

        <div class="gvb_base_container">
            <div class="gvb_inner_container gvb_index_main">
                <div class="left">
                    <GVBNews></GVBNews>
                </div>
                <div class="right">

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

.gvb_index_main {
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