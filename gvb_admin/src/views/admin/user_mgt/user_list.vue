<template>
    <div class="gvb_container">
        <a-modal v-model:visible="data.modalVidible" title="添加用户" @ok="handleOk">
            <a-form :model="formState" name="basic" :label-col="{ span: 4 }" :wrapper-col="{ span: 16 }" :ref="formRef"
                autocomplete="off">
                <a-form-item label="用户名" name="user_name" :rules="[{ required: true, message: '请输入用户名' }]">
                    <a-input v-model:value="formState.user_name" :placeholder="'用户名'" />
                </a-form-item>

                <a-form-item label="昵称" name="nick_name" :rules="[{ required: true, message: '请输入昵称' }]">
                    <a-input-password v-model:value="formState.nick_name" :placeholder="'昵称'" />
                </a-form-item>

                <a-form-item label="密码" name="password" :rules="[{ required: true, message: '请输入密码' }]">
                    <a-input-password v-model:value="formState.password" :placeholder="'密码'" />
                </a-form-item>
                <a-form-item label="确认密码" name="re_password" :rules="[{ required: true, message: '请确认密码' }]">
                    <a-input-password v-model:value="formState.re_password" :placeholder="'确认密码'" />
                </a-form-item>
                <a-form-item label="权限" name="role" :rules="[{ required: true, message: '请选择权限' }]">
                    <a-select v-model:value="formState.role" style="width: 200px" :options="options"
                        :placeholder="'请选择权限'">

                    </a-select>
                </a-form-item>


            </a-form>
        </a-modal>

        <div class="gvb_search">
            <a-input-search placeholder="搜索用户昵称" style="width: 200px" />
        </div>
        <div class="gvb_actions">
            <a-button type="primary" @click="data.modalVidible = true">添加</a-button>
            <a-button type="primary" danger @click="removeBatch" v-if="data.selectedRowKeys.length">批量删除</a-button>
        </div>
        <div class="gvb_tables">
            <a-table :columns="data.columns" :data-source="data.list" :row-key="id"
                :row-selection="{ selectedRowKeys: data.selectedRowKeys, onChange: onSelectChange }">
                <template #bodyCell="{ column, record }">
                    <template v-if="column.key === 'avatar'">
                        <img class="gvb_table_avatar" :src="record.avatar" alt="">
                    </template>
                    <template v-if="column.key === 'created_at'">
                        {{ getFormatDate(record.created_at) }}
                    </template>
                    <template v-if="column.key === 'action'">
                        <a-button class="gvb_table_action update" type="primary">编辑</a-button>
                        <a-button class="gvb_table_action delete" type="primary" danger>删除</a-button>
                    </template>
                </template>
            </a-table>
        </div>
        <div class="gvb_pages">
            <a-pagination show-less-items v-model:current="page.page" v-model:page-size="page.limit" :total="data.count"
                :show-total="total => ` 共 ${total} 人`" />

        </div>

    </div>

</template>
<script setup>

import { reactive,ref } from 'vue'
import { getFormatDate } from '@/utils/data';
import { userListApi } from '@/api/user_api';
const page = reactive({
    page: 1,
    limit: 10
})
const formRef = ref(null)
const data = reactive({
    columns: [
        { title: 'id', key: 'id', dataIndex: 'id' },
        { title: '昵称', key: 'nick_name', dataIndex: 'nick_name' },
        { title: '头像', key: 'avatar', dataIndex: 'avatar' },
        { title: '邮箱', key: 'email', dataIndex: 'email' },
        { title: '地址', key: 'addr', dataIndex: 'addr' },
        { title: 'ip', key: 'ip', dataIndex: 'ip' },
        { title: '角色', key: 'role', dataIndex: 'role' },
        { title: '注册来源', key: 'sign_status', dataIndex: 'sign_status' },
        { title: '创建时间', key: 'created_at', dataIndex: 'created_at' },
        { title: '操作', key: 'action', dataIndex: 'action' },
    ],
    list: [
        {
            "id": 1,
            "created_at": "2024-04-06T10:48:43.354+08:00",
            "nick_name": "xiaohei",
            "user_name": "xiaohei",
            "avatar": "/uploads/avatar/default.jpg",
            "email": "x****@qq.com",
            "tel": "",
            "addr": "本地",
            "token": "",
            "ip": "127.0.0.1",
            "role": "用户",
            "sign_status": "邮箱"
        },
    ],
    selectedRowKeys: [],
    count: 0,
    modalVidible: false,
})
const options = [{
    label: '管理员',
    value: 1
},
{
    label: '普通用户',
    value: 2
},
{
    label: '游客',
    value: 3
}]
const formState = reactive({
    nick_name: "",
    user_name: "",
    password: "",
    re_password: "",
    role: 1
})
function onSelectChange(selectedRowkeys) {
    data.selectedRowKeys = selectedRowkeys
}
function removeBatch() {
    // data.selectedRowKeys
}
async function getData() {
    let res = await userListApi({})
    data.list = res.data.list
    data.count = res.data.count
    console.log(res.data)
}
async function handleOk() {
    try{
        await formRef.value.validate()
    }catch(e){}
}
getData()
</script>

<style lang="scss">
.gvb_container {
    background-color: white;

    .gvb_actions {
        padding: 10px;

        .ant-btn {
            // padding: 10px;
            // border-bottom: 1px solid var(--bg);
            margin-right: 10px;
        }
    }

    .gvb_tables {
        padding: 10px;
    }

    .gvb_pages {
        display: flex;
        justify-content: center;
        padding: 10px;
    }

    .gvb_table_avatar {
        width: 40px;
        height: 40px;
        border-radius: 1cap;
    }

    .gvb_table_action.update {
        margin-right: 10px;
    }
}
</style>