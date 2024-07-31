<template>
  <div>
    <GVBTable 
      :columns="data.columns"
      baseURL="/api/menus" 
      ref="gvbTable" 
      :page-size="7" 
      :like-title="'搜索菜单昵称'"
      default-delete>
      <template #add>
        <a-button type="primary" @click="addModal">添加菜单</a-button>
      </template>
      <template #edit="{ record }">
        <a-button type="primary" @click="updateModal(record)">编辑</a-button>
      </template>
      <template #cell="{ column, record }">
        <template v-if="column.key === 'abstract'">
          <span v-for="(item,index) in record.abstract" :key="index">
            {{ item }}
          </span>
        </template>
        <template v-if="column.key === 'banners'">
          <img v-for="item in record.banners" :key="index" :src="item.path" style="height:40px;border-radius: 5px;" />
        </template>
      </template>

    </GVBTable>
  </div>
</template>

<script setup>
import { reactive } from 'vue';
import GVBTable from '@/components/admin/gvb_table.vue'
const data = reactive({
columns: [
  { title: 'id', key: 'id', dataIndex: 'id' },
  { title: '菜单标题', key: 'title', dataIndex: 'title' },
  { title: '路径', key: 'path', dataIndex: 'path' },
  { title: 'slogan', key: 'slogan', dataIndex: 'slogan' },
  { title: '介绍', key: 'abstract', dataIndex: 'abstract' },
  { title: '介绍切换时间', key: 'abstract_time', dataIndex: 'abstract_time' },
  { title: 'banner切换时间', key: 'banner_time', dataIndex: 'banner_time' },
  { title: '排序', key: 'sort', dataIndex: 'sort' },
  { title: 'banner', key: 'banners', dataIndex: 'banners' },
  { title: '发布时间', key: 'created_at', dataIndex: 'created_at' },
  { title: '操作', key: 'action', dataIndex: 'action' },
],

});

function addModal(){

}
function updateModal(record){
  
}
</script>

<style lang="scss"></style>
