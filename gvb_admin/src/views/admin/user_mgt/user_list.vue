<template>
    <div class="gvb_container">
        <div class="gvb_search">
            <a-input-search placeholder="" style="width: 200px" />
        </div>
        <div class="gvb_actions">
            <a-button type="primary">添加</a-button>
            <a-button type="primary" danger @click="removeBatch" v-if="data.selectedRowKeys.length">批量删除</a-button>
        </div>
        <div class="gvb_tables">
            <a-table 
                :columns="data.columns" 
                :data-source="data.list"
                :row-key="id"
                :row-selection="{ selectedRowKeys: data.selectedRowKeys, onChange: onSelectChange }"
                >

            </a-table>
        </div>
        <div class="gvb_pages">
            <a-pagination show-less-items v-model:current="page.page" v-model:page-size="page.limit" :total="85"
                :show-total="total => ` 共 ${total} 页`" />

        </div>

    </div>

</template>
<script setup>
import { reactive } from 'vue'
const page = reactive({
    page: 1,
    limit: 10
})
const data = reactive({
    columns: [
        { title: 'id', key: 'id', dataIndex: 'id' },
        { title: '昵称', key: 'nick_name', dataIndex: 'nick_name' },
        { title: '头像', key: 'avatar_id', dataIndex: 'avatar_id' },
        { title: '邮箱', key: 'email', dataIndex: 'email' },
        { title: '地址', key: 'addr', dataIndex: 'addr' },
        { title: 'ip', key: 'ip', dataIndex: 'ip' },
        { title: '角色', key: 'role', dataIndex: 'role' },
        { title: '注册来源', key: 'sign_status', dataIndex: 'sign_status'},
        { title: '创建时间', key: 'created_at', dataIndex: 'created_at' },
    ],
    list: [
        {
            "id": 1,
            "created_at": "2024-04-06T10:48:43.354+08:00",
            "nick_name": "xiaohei",
            "user_name": "xiaohei",
            "avatar_id": "/uploads/avatar/default.jpg",
            "email": "x****@qq.com",
            "tel": "",
            "addr": "本地",
            "token": "",
            "ip": "127.0.0.1",
            "role": "用户",
            "sign_status": 3
        },
    ],
    selectedRowKeys:[
        
    ]
})
function onSelectChange(selectedRowkeys){
    data.selectedRowKeys = selectedRowkeys
}
function removeBatch(){
    // data.selectedRowKeys
}
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
}
</style>