<template>
  <div>
    <a-modal title="添加菜单" v-model:visible="data.visible">
      <a-form
      :model="state"
      name="basic"
      ref="formReef"
      :label-col="{ span: 8 }"
      :wrapper-col="{ span: 20 }"
      autocomplete="off"
      
      >
        <a-form-item label="菜单名称" name="title" has-feedback :rules="[{required: true, message: '请输入菜单名称',trigger:'blur'}]">
          <a-input v-model:value="state.title" placeholder="菜单名称"></a-input>
        </a-form-item>
        <a-form-item label="菜单路径" name="path" has-feedback :rules="[{required: true, message: '请输入菜单路径',trigger:'blur'}]">
          <a-input v-model:value="state.path" placeholder="菜单路径"></a-input>
        </a-form-item>
        <a-form-item label="slogan" name="slogan" has-feedback>
          <a-input v-model:value="state.slogan" placeholder="slogan"></a-input>
        </a-form-item>
        <a-form-item label="菜单简介" name="abstract" has-feedback>
          <a-input v-model:value="state.abstract" placeholder="菜单简介"></a-input>
        </a-form-item>
        <a-form-item label="banner时间" name="banner_time" has-feedback >
          <a-input-number v-model:value="state.banner_time" placeholder="banner时间"></a-input-number>
        </a-form-item>
        <a-form-item label="简介时间" name="abstract_time" has-feedback>
          <a-input-number v-model:value="state.abstract_time" placeholder="abstract_time"></a-input-number>
        </a-form-item>
        <a-form-item label="序号" name="sort" has-feedback>
          <a-input-number v-model:value="state.sort" placeholder="请输入序号"></a-input-number>
        </a-form-item>

      </a-form>
    </a-modal>
    <GVBTable :columns="data.columns" baseURL="/api/menus" ref="gvbTable" :page-size="7" :like-title="'搜索菜单昵称'"
      default-delete>
      <template #add>
        <a-button type="primary" @click="addMenu">添加菜单</a-button>
      </template>
      <template #edit="{ record }">
        <a-button type="primary" @click="updateModal(record)">编辑</a-button>
      </template>
      <template #cell="{ column, record }">
        <template v-if="column.key === 'abstract'">
          <span v-for="(item, index) in record.abstract" :key="index">
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
import { reactive, ref } from 'vue';
import GVBTable from '@/components/admin/gvb_table.vue'

const state_abstract = ref("")
const _state = {  
  abstract_time: 7,
  banner_time: 7,
  path: "",
  slogan: "",
  sort: 0,
  title: "",
  abstract:[],
  image_sort_list:[],
}
const state = reactive({
  abstract_time: 7,
  banner_time: 7,
  path: "",
  slogan: "",
  sort: 0,
  title: "",
  abstract:[],
  image_sort_list:[],
})
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
  visible:false,
});

function addMenu() {
  Object.assign(state, _state)

  data.visible = true

}
function updateModal(record) {

}
</script>

<style lang="scss"></style>
