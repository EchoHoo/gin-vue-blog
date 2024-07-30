<template>
    <div>
      <a-modal title="添加广告" v-model:open="data.open" @ok="handleOk">
        <a-form :model="state" name="basic" :label-col="{ span: 4 }" :wrapper-col="{ span: 16 }" ref="formRef"
          autocomplete="off">
          <a-form-item label="广告名称" has-feedback name="title"
            :rules="[{ required: true, message: '请输入广告名称', trigger: 'blur' }]">
            <a-input v-model:value="state.title" placeholder="广告名称" />
          </a-form-item>
          <a-form-item label="广告链接" has-feedback name="href"
            :rules="[{ required: true, message: '请输入广告链接', trigger: 'blur' }, { validator: checkUrl, message: '请输入正确的链接', trigger: 'blur' }] ">
            <a-input v-model:value="state.href" placeholder="广告链接" />
          </a-form-item>
          <a-form-item label="图片地址" has-feedback name="images"
            :rules="[{ required: true, message: '请输入图片地址', trigger: 'blur' }]">
            <a-input v-model:value="state.images" placeholder="图片地址" />
          </a-form-item>
          <a-form-item label="是否显示" name="is_show">
            <a-switch v-model:checked="state.is_show" />
          </a-form-item>
          <a-form-item label="图片预览">
            <img :src="state.images" height="120px" width="160px" style="border-radius: 5px;" alt="">
          </a-form-item>
        </a-form>
      </a-modal>
      <GVBTable :columns="data.columns" baseURL="/api/adverts" ref="gvbTable" :page-size="7" :like-title="'搜索广告昵称'"
        default-delete>
        <template #add>
          <a-button type="primary" @click="data.open = true">添加广告</a-button>
        </template>
        <template #edit="{ record }">
          <a-button type="primary">编辑</a-button>
        </template>
        <template #cell="{ column, record }">
          <template v-if="column.key === 'images'">
            <img :src="record.images" style="height: 60px; border-radius: 5px;">
          </template>
          <template v-if="column.key === 'is_show'">
            <a-tag v-if="record.is_show" color="green">显示</a-tag>
            <a-tag v-else color="red">不显示</a-tag>
          </template>
          <template v-if="column.key === 'href'">
            <span>{{ record.href }}</span>
            <a-button type="link" target="_blank" :href="record.href">访问</a-button>
          </template>
        </template>
      </GVBTable>
    </div>
  </template>
  
  <script setup>
  import { reactive, ref, nextTick } from 'vue'
  import GVBTable from '@/components/admin/gvb_table.vue'
  import { useStore } from '@/stores/store';
  import { createAdvertApi } from '@/api/advert_api';
  import { message } from 'ant-design-vue';
  
  const store = useStore();
  const gvbTable = ref(null);
  const formRef = ref(null);
  const _state = reactive({
    href: "",
    title: "",
    images: "",
    is_show: true,
  })
  const state = reactive({
    href: "",
    title: "",
    images: "",
    is_show: true,
  })
  const data = reactive({
    columns: [
      { title: 'id', key: 'id', dataIndex: 'id' },
      { title: '广告标题', key: 'title', dataIndex: 'title' },
      { title: '图片路径', key: 'images', dataIndex: 'images' },
      { title: '跳转地址', key: 'href', dataIndex: 'href' },
      { title: '是否显示', key: 'is_show', dataIndex: 'is_show' },
      { title: '添加时间', key: 'created_at', dataIndex: 'created_at' },
      { title: '操作', key: 'action', dataIndex: 'action' },
    ],
    open: false,
  });
  async function handleOk() {
    console.log(state)
    let res = await createAdvertApi(state)
    if (res.code === 200) {
      message.error(res.message)
      return
    }
    message.success(res.message)
    data.open = false
    Object.assign(state, _state)
    await nextTick()
    // 确保视图已更新
  }
  let checkUrl = async (_rule, value) => {
    try {
      new URL(value)
      return Promise.resolve()
    } catch (err) {
      return Promise.reject('请输入正确的链接')
    }
  }
  </script>
  
  <style lang="scss"></style>
  