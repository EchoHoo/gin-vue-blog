<template>
    <a-menu v-model:selectedKeys="selectedKeys" mode="inline" :inline-collapsed="false" @click="goto"
        @openChange="onOpenChange" :open-keys="data.openKeys">
        <template v-for="menu in data.menuList" :key="menu.name">
            <a-menu-item :key="menu" v-if="menu.children.length === 0">
                <template #icon>
                    <i :class="'fa ' + menu.icon"></i>
                </template>
                <span>{{ menu.title }}</span>
            </a-menu-item>
            <a-sub-menu :key="menu.id" v-else>
                <template #icon>
                    <i :class="'fa ' + menu.icon"></i>
                </template>
                <template #title>{{ menu.title }}</template>
                <a-menu-item v-for="sub_menu in menu.children" :key="sub_menu">
                    <template #icon>
                        <i :class="'fa ' + sub_menu.icon"></i>
                    </template>
                    <span> {{ sub_menu.title }}</span>
                </a-menu-item>
            </a-sub-menu>
        </template>
    </a-menu>
</template>


<script setup>
import { useStore } from '@/stores/store';
import { reactive, ref } from 'vue';
import { onBeforeRouteUpdate, useRoute, useRouter } from "vue-router";
const levelColor = {
    "debug": "green",
    "info": "blue",
    "warn": "orange",
    "error": "red"
}
const levelOptions = [
    { label: "debug", value: 1, }, { label: "info", value: 2, }, { label: "warn", value: 3, }, { label: "error", value: 4, },
]
const selectedKeys = ref([]);
const store = useStore();
const data = reactive({
    menuList: [{
        id: 1,
        icon: "fa-home",
        title: "首页",
        name: "home",
        children: []
    },
    {
        id: 2,
        icon: "fa-users",
        title: "个人中心",
        name: "user_center",
        children: [
            {
                id: 21,
                icon: "fa-users",
                title: "我的信息",
                name: "user_info",
            },
            {
                id: 22,
                icon: "fa-users",
                title: "我的发布",
                name: "user_article_list",
            },
            {
                id: 23,
                icon: "fa-users",
                title: "我的收藏",
                name: "user_collects",
            },
            {
                id: 24,
                icon: "fa-users",
                title: "我的消息",
                name: "user_messages",
            },
        ]
    },
    {
        id: 3,
        icon: "fa-users",
        title: "用户管理",
        name: "",
        children: [
            {
                id: 31,
                icon: "fa-list",
                title: "用户列表",
                name: "user_list",
            },
            {
                id: 32,
                icon: "fa-cogs",
                title: "群聊消息",
                name: "chat_list",
            },
            {
                id: 33,
                icon: "fa-cogs",
                title: "消息列表",
                name: "message_list",
                children: []
            },
        ]
    },
    {
        id: 4,
        icon: "iconfont icon-tupianguanli1",
        title: "图文管理",
        name: "",
        children: [
            {
                id: 41,
                icon: "iconfont icon-tupianguanli1",
                title: "图片管理",
                name: "image_list",
            },
            {
                id: 42,
                icon: "fa-cogs",
                title: "文章添加",
                name: "add_article",
            },
            {
                id: 43,
                icon: "fa-cogs",
                title: "文章管理",
                name: "article_list",
            },
            {
                id: 14,
                icon: "fa-cogs",
                title: "标签管理",
                name: "tag_list",
            },
        ]
    },
    {
        id: 5,
        icon: "fa-cogs",
        title: "系统管理",
        name: "",
        children: [
            {
                id: 51,
                icon: "fa-cogs",
                title: "网站设置",
                name: "system_site",
            },
            {
                id: 52,
                icon: "fa-cog",
                title: "系统日志",
                name: "log_list",
            },
            {
                id: 53,
                icon: "iconfont icon-tupianguanli1",
                title: "菜单管理",
                name: "menu_list",
            },
            {
                id: 54,
                icon: "iconfont icon-tupianguanli1",
                title: "广告管理",
                name: "advert_list",
            }
        ]
    },

    ],
    openKeys: []
});
const router = useRouter();
const route = useRoute();

function goto({
    item,
    key,
    keyPath
}) {
    store.addTab({
        name: key.name,
        title: key.title
    });
    router.push({
        name: key.name
    });
}

function onOpenChange(openKeys) {
    const latestOpenKey = openKeys.find(key => data.openKeys.indexOf(key) === -1);
    data.openKeys = latestOpenKey ? [latestOpenKey] : [];
}

function loadRoute(name) {
    if (name === undefined) {
        name = route.name;
    }
    for (const menu of data.menuList) {
        if (menu.name === name) {
            selectedKeys.value = [menu];
            return;
        }
        for (const sub_menu of menu.children) {
            if (sub_menu.name === name) {
                selectedKeys.value = [sub_menu];
                data.openKeys = [menu.id];
                return;
            }
        }
    }
}

onBeforeRouteUpdate((to, from, next) => {
    loadRoute(to.name);
    next();
});

loadRoute();
</script>

<style>
.ant-menu {
    background-color: transparent;
    color: white;
}

.ant-menu-submenu-arrow {
    color: white;
}

.ant-menu-submenu.ant-menu-submenu-inline {
    color: white;
}

.ant-menu .ant-menu-submenu.ant-menu-submenu-selected {
    color: inherit !important;
}
</style>