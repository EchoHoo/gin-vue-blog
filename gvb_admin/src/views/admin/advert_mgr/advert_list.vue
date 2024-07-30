<template>
    <div>
        <GVBTable 
            :columns="data.columns" 
            baseURL="/api/adverts" 
            ref="gvbTable"
            :page-size="7"
            :like-title="'搜索广告昵称'"
            default-delete
        >
            <template #add>
                <a-button type="primary" >上传图片</a-button>
            </template>
            <template #edit="{record}">
                <a-button type="primary">编辑</a-button>
            </template>
            <template #cell="{column, record}">
                <template v-if="column.key === 'images'">
                    <img :src="record.images" style="height: 60px; border-radius: 5px;">
                </template>
                <template v-if="column.key === 'is_show'">
                    <a-tag v-if="record.is_show" color="green">显示</a-tag>
                    <a-tag v-else color="red">不显示</a-tag>
                </template>
                <template v-if="column.key === 'href'">
                    <span>{{ record.href}}</span>
                    <a-button type="link" target="_blank" :href="record.href" >访问</a-button>
                </template>
            </template>
        </GVBTable>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import GVBTable from '@/components/admin/gvb_table.vue'
import { imageRenameApi } from '@/api/image_api';
import { message } from 'ant-design-vue';
import { useStore } from '@/stores/store';

const store = useStore();
const gvbTable = ref(null);
const formRef = ref(null);

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
    modalVisible: false,
    visible: false,
});

</script>

<style lang="scss"></style>