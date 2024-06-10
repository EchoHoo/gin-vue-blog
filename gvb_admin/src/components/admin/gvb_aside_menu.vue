<template>
    <a-menu v-model:selectedKeys="selectedKeys" mode="inline" theme="dark" :inline-collapsed="false">
        <template v-for="menu in data.menuList" :key="menu.id">
            <a-menu-item :key="menu.id" v-if="menu.children.length === 0">
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
                <a-menu-item  v-for="sub_menu in menu.children" :key="sub_menu.id">
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
import { reactive, ref } from 'vue';
const selectedKeys = ref(["1"]);

const data = reactive({
    menuList: [{
        id: 1,
        icon: "fa-home", // 菜单图标
        title: "首页", // 菜单标题
        name: "home",    // 路由名称
        children: []
    },
    {
        id: 2,
        icon: "fa-users", // 菜单图标
        title: "用户列表", // 菜单标题
        name: "",    // 路由名称
        children: [
            {
                id: 3,
                icon: "fa-list", // 菜单图标
                title: "用户管理", // 菜单标题
                name: "user_list",    // 路由名称
            },
            {
                id: 4,
                icon: "fa-user-plus", // 菜单图标
                title: "添加用户", // 菜单标题
                name: "user_create",    // 路由名称
            },
        ]
    },
    {
        id: 5,
        icon: "fa-cogs", // 菜单图标
        title: "系统管理", // 菜单标题
        name: "",    // 路由名称
        children: [
            {
                id: 6,
                icon: "fa-cog", // 菜单图标
                title: "系统配置", // 菜单标题
                name: "user_list",    // 路由名称
            },

        ]
    },
    ]
})
</script>

<style scoped></style>
