<template>
  <div>
    <a-modal :title="editID?'编辑菜单':'添加菜单'" v-model:visible="data.visible" @ok="handleOk">
      <a-form
      :model="state"
      name="basic"
      ref="formRef"
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
          <a-textarea 
          v-model:value="state_abstract" 
          name="abstract" 
          placeholder="菜单简介，支持多行" 
          :auto-size="{ minRows: 2, maxRows: 6 }"
        >
          </a-textarea>
          <!-- <a-input v-model:value="state_abstract" placeholder="菜单简介"></a-input> -->
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
        <a-form-item label="banner选择">
             <a-select v-model:value="bannerIDList" placeholder="请选择banner" mode="multiple" @change="handleBannerChange">
               <a-select-option v-for="item in data.imageNameList" :key="item.id" :value="item.id">
                <img :src="item.path" alt="" height="30" style="border-radius: 5px; margin-right: 10px;">
                <span>{{ item.name }}</span>
              </a-select-option>
              <template #tagRender="{value:val,label,closable,onClose,option}">
                  <img :src="getLabel(label)" height="30" style="border-radius: 5px; margin-right: 10px;">
              </template>
             </a-select>
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
          <div style="display: grid; grid-template-columns: repeat(3,1fr);grid-column-gap: 5px;grid-column-gap: 5px;">
            <div v-for="(item, index) in record.abstract" :key="index">{{ item }}</div>
          </div>
        </template>
        <template v-if="column.key === 'banners'">
          <div style="display: grid; grid-template-columns: repeat(3,1fr);grid-row-gap: 5px;grid-column-gap: 5px;">
            <img v-for="item in record.banners" :key="index" :src="item.path" style="height:40px;border-radius: 5px;" />
          </div>

   
        </template>
      </template>

    </GVBTable>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue';
import GVBTable from '@/components/admin/gvb_table.vue'
import { imageNameListApi } from '@/api/image_api';
import { createMenuApi, updateMenuApi } from '@/api/menu_api';
import { message } from 'ant-design-vue';


const state_abstract = ref("")
const formRef = ref(null)
const bannerIDList = ref([])
const gvbTable = ref(null);
const _state = {  
  abstract_time: 7,
  banner_time: 7,
  path: "",
  slogan: "",
  sort: 1,
  title: "",
  abstract:[],
  image_sort_list:[],
}
const state = reactive({
  abstract_time: 7,
  banner_time: 7,
  path: "",
  slogan: "",
  sort: 1,
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
  imageNameList: [],
});
function getLabel(label){
  return label[0].props.src
}
function handleBannerChange(idList){
  state.image_sort_list = []
  for(let i=0;i<idList.length;i++){
    state.image_sort_list.push({
      image_id: idList[i],
      sort: i
    })
  }
  
}
async function getData(){
  let res = await imageNameListApi()
  data.imageNameList = res.data
}
getData()
function addMenu() {
  Object.assign(state, _state)
  editID.value = 0
  data.visible = true

}
async function handleOk() {
  try{
    formRef.value.validate();
  }catch(err){
    return 
  }
  state.abstract = state_abstract.value.split("\n")
  let res 
  if(editID.value){
    res = await updateMenuApi(editID.value, state)
  }else{
    res = await createMenuApi(state)
  }

  if (res.code){
    message.error(res.msg)
    return
  }
  message.success(res.msg)
  data.visible = false
  state_abstract.value = ""
  bannerIDList.value = []
  gvbTable.value.ExportList()
  
  console.log(state)
}
const editID = ref(0)
function updateModal(record) {
  editID.value = record.id
  state.title = record.title
  state.path = record.path
  state.slogan = record.slogan
  state.sort = record.sort
  state.abstract_time = record.abstract_time
  state.banner_time = record.banner_time
  state_abstract.value = record.abstract.join("\n")
  
  bannerIDList.value = []
  state.image_sort_list=[]
  for(let i=0;i<record.banners.length;i++){
    bannerIDList.value.push(record.banners[i].id) 
    state.image_sort_list.push({
      image_id: record.banners[i].id,
      sort: i
    })
  }
  // state.abstract = record.abstract
  // state.image_sort_list = []
  data.visible = true
}
</script>

<style lang="scss"></style>
