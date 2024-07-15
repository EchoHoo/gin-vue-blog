<template>
    <div class="gvb_container">
        <div class="gvb_search">
            <a-input-search placeholder="" style="width: 200px" />
        </div>
        <div class="gvb_actions">
            <a-button type="primary">Primary Button</a-button>
            <a-button>Default Button</a-button>
        </div>
        <div class="gvb_tables">
            <a-table :columns="data.columns" :data-source="data.list">

            </a-table>
        </div>
        <div class="gvb_pages" >
            <a-pagination show-less-items 
            v-model:current="page.page" 
            v-model:page-size="page.limit" :total="85"
            :show-total="total => ` 共 ${total} 页`" />

        </div>

    </div>

</template>
<script setup>
import { reactive } from 'vue'
const page = reactive({
    page:1,
    limit:10
})
const data = reactive({
    columns: [
        {
            title: 'Name',
            key: 'name',
            dataIndex: 'name'
        },
        {
            title: 'age',
            key: 'age',
            dataIndex: 'age'
        },
        {
            title: 'data',
            key: 'data',
            dataIndex: 'data'

        }
    ],
    list: [
        {
            name: 'Tom',
            age: 25,
            data: '2021-01-01'
        },
        {
            name: 'Jerry',
            age: 26,
            data: '2021-01-02'
        }
    ]
})

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