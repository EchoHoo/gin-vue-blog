<template>
    <div class="gvb_tabs">
        <div :class="isActive(item)" 
        v-for="(item,index) in store.tabList" 
        :key="index"
        @click="checkTable(item)"
        >
            {{item.title}}
            <span @click.stop ="removeTab(item)" v-if="item.name!== 'home'" class="gvb_tab_remove_icon">x</span>
        </div>
      
        <div class="gvb_tab_item remove_all" @click="removeTabAll">
            全部关闭
        </div>
    </div>
</template>
<script setup>
import { useStore } from '@/stores/store';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const store = useStore();
const router = useRouter();
function isActive(item){
    if (route.name === item.name){
        return "gvb_tab_item active"
    }
    return "gvb_tab_item"
}
function checkTable(item){
    router.push({
        name:item.name
    })
}
function removeTab(item){
    let index = store.removeTab(item);
    // 判断是不是当前的tab，是的话向前走一步路由
    if (route.name === item.name){
        router.push({
            name:store.tabList[index-1].name
        })
    }
}
function removeTabAll(){
    store.removeTabAll()
    router.push({
            name:"home"
    })
}
store.loadTabs();
window.onbeforeunload = function(e) {
    // 保存tabs
    store.saveTabs();
    return undefined;
}
</script>
<style lang="scss">
.gvb_tabs{
    height:30px;
    border-color:var(--order);
    border-style: solid;
    border-width: 1px 0 1px 0;
    border: 1px solid var(--order);
    background-color:var(--card_bg);
    display: flex;
    justify-content: left;
    align-items: center;
    position: relative;
    .gvb_tab_item{
        background-color: var(--card_bg);
        border: 1px solid var(--card_border);
        padding: 0 10px;
        margin-right: 10px;
        cursor: pointer;
        color:var(--text);
        display: flex;
        align-items: center;

        .gvb_tab_remove_icon{
            cursor: pointer;
            border-radius: 5px;
            width: 12px;
            height: 12px;
            display: flex;
            justify-content: center;
            align-items: center;
            margin: 5px;
            &:hover{
                background-color: #f0eeee;

            }
        }
        &:first-child{
            margin-left: 20px;
        }
        &.active{
            background-color: var(--active);
            border-color: var(--active);
            color: white;
            &::before{
                display: inline-block;
                content:"";
                width: 12px;
                height:12px;
                background-color: white;
                border-radius: 50%;
                margin-right: 5px;
            }
            .gvb_tab_remove_icon:hover{
                background-color: #96d2f8;
            }
        }
        &.remove_all{
            position: absolute;
            right:20px;
        }
    }
}
</style>