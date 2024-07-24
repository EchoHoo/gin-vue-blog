<template>
    <div>
        <GVBTable 
            @delete="Delete" 
            @edit="updateModal" 
            :columns="data.columns" 
            baseURL="/api/images" 
            ref="gvbTable"
            :page-size="7"
            :like-title="'搜索图片昵称'"
        >
            <template #add>
               
            </template>
            <template #edit="{record}">
              
            </template>
            <template #cell="{column, record}">
                <template v-if="column.key === 'path'">
                    <img :src="record.path" style="width: 100px; height: 100px;">
                </template>
            </template>
        </GVBTable>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import GVBTable from '@/components/admin/gvb_table.vue'

const gvbTable = ref(null);

const data = reactive({
    columns: [
        { title: 'id', key: 'id', dataIndex: 'id' },
        { title: '图片名称', key: 'name', dataIndex: 'name' },
        { title: '路径', key: 'path', dataIndex: 'path' },
        { title: '上传时间', key: 'created_at', dataIndex: 'created_at' },
        { title: '操作', key: 'action', dataIndex: 'action' },
    ],
});


function Delete(idList) {
    console.log(userIdList);
}
</script>

<style lang="scss"></style>
