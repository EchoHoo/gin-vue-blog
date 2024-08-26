<template>
    <div :class="{
        gvb_nav: true,
        show: data.is_show,
    }">

        <div class="gvb_nav_container">
            <div class="logo">
                <div>LOGO</div>

            </div>
            <div class="left">
                <span v-for="item in store.navList" :key="item.id">
                    <router-link :to="item.path">{{ item.title }}</router-link>
                </span>
                <span class="search">
                    <GVBTestSearch></GVBTestSearch>
                </span>
            </div>
            <div class="right">
                <span class="login_btn" v-if="store.userInfo.role === 0"><router-link
                        :to="{ name: 'login', query: { redirect_url: $router.path } }">登录</router-link></span>
                <GVBUserInfo v-else class="gvb_user_info"></GVBUserInfo>
            </div>
        </div>
    </div>
</template>
<script setup>

import GVBUserInfo from "@/components/gvb_user_info.vue"
import { useStore } from "@/stores/store";
import { reactive } from "vue";
import GVBTestSearch from "@/components/gvb_text_search.vue"

const store = useStore();
const props = defineProps({
    is_show: {
        type: Boolean,
        default: false,
    }
})

const data = reactive({
    is_show: false,
})

async function getInit() {
    if (props.is_show) {
        data.is_show = true
    } else {
        window.addEventListener("scroll", scroll)
    }
}
getInit()
function scroll() {
    // 获取滚动条距离页面顶部的距离
    let top = document.documentElement.scrollTop
    if (top >= 200) {
        data.is_show = true
    } else {
        data.is_show = false
    }
}

</script>
<style lang="scss">
.gvb_nav {
    position: fixed;
    top: 0;
    width: 100%;
    height: 60px;
    display: flex;
    justify-content: center;
    font-size: 16px;
    z-index: 100;
    color: white;

    .logo {
        width: 10%;
        transform: scale(0.7);
        color: rgb(255, 255, 255);

        div::first-child {
            font-size: 26px;
        }

        div:last-child {
            font-size: 12px;
        }
    }

    .gvb_nav_container {
        width: 1200px;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .left {
        width: 50%;
        display: flex;

        span {
            margin-right: 30px;
        }

        .search i {
            cursor: pointer;

        }


    }

    .right {
        width: 40%;
        display: flex;
        justify-content: right;
        align-items: center;

        .login_btn {
            margin-right: 20px;
        }

    }

    a {
        color: white;

        &:hover {
            color: var(--active)
        }

        &.router-link-exact-active {
            color: var(--active) !important;
        }
    }
}

.gvb_nav.show {
    color: var(--text);
    background-color: var(--card_bg);
    box-shadow: 1px 1px 5px #0003;

    a {
        color: var(--text);

    }
}
</style>