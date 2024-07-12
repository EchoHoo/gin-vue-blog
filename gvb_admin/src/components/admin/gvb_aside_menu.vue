<template>
    <a-menu v-model:selectedKeys="selectedKeys"
            mode="inline"
            :inline-collapsed="false"
            @click="goto"
    >
        <template v-for="menu in data.menuList" :key="menu.name">
            <a-menu-item :key="menu.name" v-if="menu.children.length === 0">
                <template #icon>
                    <i :class="'fa ' + menu.icon"></i>
                </template>
                <span>{{ menu.title }}</span>
            </a-menu-item>
            <a-sub-menu :key="menu.name" v-else>
                <template #icon>
                    <i :class="'fa ' + menu.icon"></i>
                </template>
                <template #title>{{ menu.title }}</template>
                <a-menu-item  v-for="sub_menu in menu.children" :key="sub_menu.name">
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
import {useRouter} from "vue-router";

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
        icon: "fa-users", 
        title: "用户列表", 
        name: "",    // 路由名称
        children: [
            {
                id: 3,
                icon: "fa-list", 
                title: "用户管理", 
                name: "user_list",    // 路由名称
            },
        ]
    },
    {
        id: 5,
        icon: "fa-cogs", 
        title: "系统管理",
        name: "",    // 路由名称
        children: [
            {
                id: 6,
                icon: "fa-cog", 
                title: "系统配置", 
                name: "system_list",    // 路由名称
            },

        ]
    },
    {
        id: 5,
        icon: "fa-cogs", 
        title: "图片管理",
        name: "",    // 路由名称
        children: [
            {
                id: 6,
                icon: "fa-cog", 
                title: "图片列表", 
                name: "system_list",    // 路由名称
            },
        ]
    }
    ]
})
const router = useRouter()
function goto(event){
  router.push({
    name: event.key
  })

}

</script>

<style scoped>
.ant-menu{
  background-color: transparent;
  color:white;
}

.ant-menu-submenu-arrow{
  color: white;
}
.ant-menu-submenu.ant-menu-submenu-inline{
  color: white;
}
</style>
