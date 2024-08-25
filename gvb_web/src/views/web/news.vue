<template>
    <div class="base_view index_view">
        <GVBNav></GVBNav>
        <GVBBanner url="https://c-ssl.duitang.com/uploads/blog/202007/22/20200722113343_4893c.jpeg" slogan="xxx"
            abstract="abstract" :is-article=false>

        </GVBBanner>
        <div class="gvb_base_container">
            <div class="gvb_inner_container">

            </div>
        </div>
        <div class="gvb_footer">

        </div>
    </div>
</template>

<script setup>
import GVBNav from "@/components/gvb_nav.vue"
import GVBBanner from "@/components/gvb_banner.vue"
</script>


<style lang="scss">
.base_view {
    background-color: var(--card_bg);


    .gvb_base_container {
        display: flex;
        justify-content: center;

        .gvb_inner_container {
            background-color: white;
            min-height: 1000px;
            width: 1200px;
            margin-top: 20px;
        }
    }
}
</style>