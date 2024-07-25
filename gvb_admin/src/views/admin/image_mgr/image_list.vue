<template>
    <div>
        <a-modal 
            title="编辑图片" 
            v-model:visible="data.modalVisible"
            @ok="update"
        >
            <a-form :model="updateState" name="basic" :label-col="{ span: 4 }" :wrapper-col="{ span: 16 }" ref="formRef"
                autocomplete="off">
                <a-form-item label="文件名称" has-feedback name="name"
                    :rules="[{ required: true, message: '请输入文件名称', trigger: 'blur' }]">
                    <a-input v-model:value="updateState.name" placeholder="文件名称" />
                </a-form-item>
                <a-form-item label="图片预览">
                    <img :src="updateState.path" style="height: 100px; border-radius: 5px;">
                </a-form-item>
            </a-form>
        </a-modal>
        <GVBTable 
            :columns="data.columns" 
            baseURL="/api/images" 
            ref="gvbTable"
            :page-size="7"
            :like-title="'搜索图片昵称'"
            default-delete
        >
            <template #add>
                <!-- Add button or any content here -->
            </template>
            <template #edit="{record}">
                <a-button type="primary" @click="updateModal(record)">编辑</a-button>
            </template>
            <template #cell="{column, record}">
                <template v-if="column.key === 'path'">
                    <img :src="record.path" style="height: 60px; border-radius: 5px;">
                </template>
            </template>
        </GVBTable>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import GVBTable from '@/components/admin/gvb_table.vue'
import { imageRenameApi, imageDeleteApi } from '@/api/image_api';  // 确保你有正确的删除 API
import { message } from 'ant-design-vue';

const gvbTable = ref(null);
const formRef = ref(null);

const data = reactive({
    columns: [
        { title: 'id', key: 'id', dataIndex: 'id' },
        { title: '图片名称', key: 'name', dataIndex: 'name' },
        { title: '图片', key: 'path', dataIndex: 'path' },
        { title: '类型', key: 'image_type', dataIndex: 'image_type' },
        { title: '上传时间', key: 'created_at', dataIndex: 'created_at' },
        { title: '操作', key: 'action', dataIndex: 'action' },
    ],
    modalVisible: false,
});

const updateState = reactive({
    id: 0,
    name: "",
    path:"",
});

function updateModal(record) {
    data.modalVisible = true;
    updateState.id = record.id;
    updateState.name = record.name;
    updateState.path = record.path;
}

async function update(){
    try{
        let res = await imageRenameApi(updateState) 
        if (res.code){
            message.error(res.msg)
            return
        }
        message.success(res.msg)
        data.modalVisible = false;
        gvbTable.value.ExportList()
    }catch(e){
        message.error('更新失败');
    }
}

async function Delete(id){
    try{
        let res = await imageDeleteApi([id]);  // 确保正确的 API 被调用
        if (res.code){
            message.error(res.msg);
            return;
        }
        message.success(res.msg);
        gvbTable.value.ExportList();
    }catch(e){
        message.error('删除失败');
    }
}
</script>

<style lang="scss"></style>
