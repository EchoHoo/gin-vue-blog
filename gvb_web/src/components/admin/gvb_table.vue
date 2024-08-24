<template>
    <div class="gvb_container">
        <div class="gvb_search">
            <a-input-search 
                :placeholder="props.likeTitle" 
                v-model:value="page.key" 
                style="width: 200px" 
                @search="onSearch" 
            />
            <slot name="filters"></slot>
            <div class="gvb_refresh">
                <a-button title="刷新" @click="refresh"><i class="fa fa-refresh"></i></a-button>
            </div>
        </div>
        <div class="gvb_divider"></div>
        <div class="gvb_actions">
            <slot name="add">
                <!-- <a-button type="primary" @click="addModal" v-if="props.isAdd">添加</a-button> -->
            </slot>
            <slot name="batchRemove">
                <a-button type="primary" danger @click="removeBatch"
                    v-if="props.isDelete && data.selectedRowKeys.length">批量删除</a-button>
            </slot>
        </div>
        <div class="gvb_tables">
            <a-spin :spinning="data.spinning" tip="Loading..." :delay="500">
                <a-table :columns="props.columns" :data-source="data.list" :row-key="record => record.id"
                    :pagination="false"
                    :row-selection="props.isSelect? { selectedRowKeys: data.selectedRowKeys, onChange: onSelectChange }:undefined">
                    <template #bodyCell="{ column, record }">
                        <slot name="cell" v-bind="{ column, record }">
                            <template v-if="column.key === 'created_at'">
                                {{ getFormatDate(record.created_at) }}
                            </template>
                            <template v-if="column.key === 'action'">
                                <slot name="edit" v-bind="{ column, record }">
                                    <a-button type="primary" v-if="props.isEdit">编辑</a-button>
                                </slot>
                                <slot name="delete" v-bind="{ column, record }">
                                    <a-popconfirm title="是否确定删除?" 
                                        ok-text="确认" 
                                        cancel-text="取消"
                                        @confirm="userRemove(record.id)" 
                                        v-if="props.isDelete">
                                        <a-button class="gvb_table_action delete" type="primary" danger>删除</a-button>
                                    </a-popconfirm>
                                </slot>
                            </template>
                        </slot>
                    </template>
                </a-table>
            </a-spin>
        </div>
        <div class="gvb_pages">
            <a-pagination show-less-items v-model:current="page.page" @change="pageChange"
                v-model:page-size="page.limit" :total="data.count" :show-total="total => `  ${total} `" />
        </div>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { getFormatDate } from '@/utils/data';
import { baseDeleteApi, baseListApi } from '@/api/base_api';
import { message } from 'ant-design-vue';

const emits = defineEmits(["delete", "edit"])
const props = defineProps({
    columns: {
        type: Array,
    },
    baseURL: {
        type: String,
    },
    isAdd: {
        type: Boolean,
        default: true
    },
    isEdit: {
        type: Boolean,
        default: true
    },
    isDelete: {
        type: Boolean,
        default: true
    },
    isSelect: {
        type: Boolean,
        default: true
    },
    filters:{
        type: Array,
    },
    pageSize:{
        type: Number,
        default: 7
    },
    likeTitle:{
        type: String,
        default: "模糊搜索"
    },
    defaultDelete:{
        type: Boolean, 
        default: false
    }
})

const page = reactive({
    page: 1,
    limit: props.pageSize,
    key: "",
});
const data = reactive({
    list: [],
    selectedRowKeys: [],
    count: 0,
    spinning: true,
});

function onSelectChange(selectedRowkeys) {
    data.selectedRowKeys = selectedRowkeys;
}

async function removeBatch() {
    if (props.defaultDelete){
        let res = await baseDeleteApi(props.baseURL, data.selectedRowKeys);
        if (res.code){
            message.error(res.msg);
            return
        }
        message.success(res.msg);
        getData(page);
        return;
    }
    emits("delete", data.selectedRowKeys);
    data.selectedRowKeys = [];
}

function pageChange(_page, limit) {
    page.page = _page;
    page.limit = limit;
    getData(page);

}

async function userRemove(id) {
    if (props.defaultDelete) {
    let res = await baseDeleteApi(props.baseURL, [id])
    if (res.code) {
      message.error(res.msg)
      return
    }
    message.success(res.msg)
    getData(page)
    return
  }
    emits("delete", [id]);
}

async function getData(params) {
    if(props.baseURL === ""){
        data.spinning = false;
        data.list = props.list;
        return;
    }
    data.spinning = true;
    let res = await baseListApi(props.baseURL, params);
    data.spinning = false;
    if (res.data.list === undefined){
        data.list = res.data;
        data.count = res.data.length;
        return;
    }
    data.list = res.data.list;
    data.count = res.data.count;
    data.spinning = false;
}

function onSearch(){
    page.key = page.key.trim();
    page.page = 1;
    getData(page);
}

function refresh() {
    getData(page);
    message.success('刷新成功');
}

function addModal() {
    data.modalVisible = true;
}

function ExportList(params) {
    if (!params) {
        params = {};
    }
    page.page = 1;
    Object.assign(page, params);
    getData(page);
}

defineExpose({
    ExportList,
})

getData(page);
</script>

<style lang="scss">
.gvb_container {
    background-color: white;

    .gvb_search {
        padding: 10px;
        border-bottom: var(--card_bg);

        position: relative;

        .gvb_refresh {
            position: absolute;
            right: 10px;
            top: 10px;

            i {
                color: var(--text);
            }
        }
    }

    .gvb_divider {
        width: 100%;
        height: 1px;
        background-color: var(--divider-color, #dcdcdc);
        margin: 10px 0;
    }

    .gvb_actions {
        padding: 10px;

        .ant-btn {
            margin-right: 10px;
        }
    }

    .gvb_tables {
        padding: 10px;
        .ant-btn{
            margin-left:10px;
        }
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
