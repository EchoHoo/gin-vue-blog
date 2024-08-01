<template>
    <div>
        <a-modal v-model:visible="data.modalVidible" title="添加用户" @ok="handleOk">
            <a-form :model="formState" name="basic" :label-col="{ span: 4 }" :wrapper-col="{ span: 16 }" ref="formRef"
                autocomplete="off">
                <a-form-item label="用户名" has-feedback name="user_name"
                    :rules="[{ required: true, message: '请输入用户名', trigger: 'blur' }]">
                    <a-input v-model:value="formState.user_name" placeholder="用户名" />
                </a-form-item>

                <a-form-item label="昵称" has-feedback name="nick_name"
                    :rules="[{ required: true, message: '请输入昵称', trigger: 'blur' }]">
                    <a-input v-model:value="formState.nick_name" placeholder="昵称" />
                </a-form-item>

                <a-form-item label="密码" has-feedback name="password"
                    :rules="[{ required: true, message: '请输入密码', trigger: 'blur' }]">
                    <a-input-password v-model:value="formState.password" placeholder="密码" />
                </a-form-item>
                <a-form-item label="确认密码" has-feedback name="re_password"
                    :rules="[{ validator: validatePassword, trigger: 'blur' }]">
                    <a-input-password v-model:value="formState.re_password" placeholder="确认密码" />
                </a-form-item>
                <a-form-item label="权限" has-feedback name="role" :rules="[{ required: true, message: '请选择权限' }]">
                    <a-select v-model:value="formState.role" style="width: 200px" :options="options"
                        placeholder="请选择权限">
                    </a-select>
                </a-form-item>
            </a-form>
        </a-modal>
        <a-modal v-model:visible="data.modalUpdateVidible" title="编辑用户" @ok="update">
            <a-form :model="formUpdateState" name="basic" :label-col="{ span: 4 }" :wrapper-col="{ span: 16 }"
                ref="formRef" autocomplete="off">
                <a-form-item label="昵称" has-feedback name="nick_name"
                    :rules="[{ required: true, message: '请输入昵称', trigger: 'blur' }]">
                    <a-input v-model:value="formUpdateState.nick_name" placeholder="昵称" />
                </a-form-item>

                <a-form-item label="权限" has-feedback name="role" :rules="[{ required: true, message: '请选择权限' }]">
                    <a-select v-model:value="formUpdateState.role" style="width: 200px" :options="options"
                        placeholder="请选择权限">
                    </a-select>
                </a-form-item>
            </a-form>
        </a-modal>

        <GVBTable 
            @delete="userDelete" 
            @edit="updateModal" 
            :columns="data.columns" 
            baseURL="/api/users" 
            ref="gvbTable"
            :page-size="7"
            :like-title="'搜索用户昵称'"
        >
            <template #add>
                <a-button type="primary" @click="data.modalVidible = true">添加</a-button>
            </template>
            <template #edit="{ record }">
                <a-button class="gvb_table_action update" type="primary" @click="updateModal(record)">编辑</a-button>
            </template>
            <template #cell="{ column, record }">
                <template v-if="column.key === 'avatar'">
                    <img class="gvb_table_avatar" :src="record.avatar" alt="" />
                </template>
            </template>
            <template #filters>
                <a-select class="gvb_select" allowClear v-model:value="filter" @change="onFilters"
                    style="width: 200px" :options="options" placeholder="选择权限">
                </a-select>
                <!-- <div class="gvb_refresh">
                    <a-button title="刷新" @click="refresh"><i class="fa fa-refresh"></i></a-button>
                </div> -->
            </template>
        </GVBTable>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { userCreateApi, updateUserNickNameApi, userListApi } from '@/api/user_api';
import { message } from 'ant-design-vue';
import GVBTable from '@/components/admin/gvb_table.vue'
import { baseDeleteApi } from '@/api/base_api';

const formRef = ref(null);
const filter = ref(undefined)
const gvbTable = ref(null);
const options = [
    { label: '管理员', value: 1 },
    { label: '用户', value: 2 },
    { label: '游客', value: 3 }
];

const data = reactive({
    modalVidible: false, // 添加用户modal
    modalUpdateVidible: false, // 编辑用户modal
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
});

const formUpdateState = reactive({
    nick_name: "",
    role: undefined,
    user_id: 0
});

const formState = reactive({
    nick_name: "",
    user_name: "",
    password: "",
    re_password: "",
    role: 1
});

function onFilters(){
    console.log(filter.value);
    gvbTable.value.ExportList({role:filter.value})
}
async function userDelete(userIdList) {
    let res = await baseDeleteApi('/api/users', userIdList);
    if (res.code) {
        message.error(res.msg);
        return;
    }
    message.success(res.msg);
    gvbTable.value.ExportList(); // Refresh table data
}

function updateModal(record) {
    data.modalUpdateVidible = true;
    formUpdateState.user_id = record.id;
    formUpdateState.nick_name = record.nick_name;
    formUpdateState.role = record.role;
    getData();
}

async function handleOk() {
    try {
        await formRef.value.validate();
        console.log(formState);
        let res = await userCreateApi(formState);
        if (res.code) {
            message.error(res.msg);
            return;
        }
        message.success(res.msg);
        data.modalVidible = false;
        formRef.value.resetFields();
        gvbTable.value.ExportList();
        getData();
    } catch (e) {
        console.error('Validation or API call failed:', e);
    }
}

let validatePassword = async (_rule, value) => {
    if (value === '') {
        return Promise.reject('密码不能为空');
    } else if (value !== formState.password) {
        return Promise.reject('两次密码输入不一致');
    } else {
        return Promise.resolve();
    }
};

async function update() {
    console.log(formUpdateState);

    let res = await updateUserNickNameApi(formUpdateState);
    if (res.code) {
        message.error(res.msg);
        return;
    }
    message.success(res.msg);
    data.modalUpdateVidible = false;
    gvbTable.value.ExportList();
    getData();
}
async function getData() {
    // 调用 GVBTable 组件中的 getData 方法
    const tableComponent = ref(null);
    tableComponent.value.getData();
    gvbTable.value.ExportList();
}
</script>

<style lang="scss"></style>
