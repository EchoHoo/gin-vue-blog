<template>
    <div>
        <GVBTable :columns="data.columns" baseURL="/api/articles" like-title="搜索文章标题" ref="gvbTable" default-delete>
            <template #add>
                <a-button type="primary">添加</a-button>
            </template>
            <template #edit="{ record }">
                <a-button type="primary">编辑</a-button>
            </template>
            <template #cell="{ column, record, index }">
                <template v-if="column.key === 'index'">
                    <span>{{ index + 1 }}</span>
                </template>
                <template v-if="column.key === 'title'">
                    <span class="gvb_article_title" v-html="record.title"></span>
                </template>
                <template v-if="column.key === 'banner_url'">
                    <img :src="'/' + record.banner_url" alt="" height="60" style="border-radius: 5px">
                </template>
                <template v-if="column.key === 'data'">
                    <div class="gvb_article_data">
                        <span><i class="fa fa-eye"></i>{{ record.look_count }}</span>
                        <span><i class="fa fa-thumbs-o-up"></i>{{ record.digg_count }}</span>
                        <span><i class="fa fa-comment-o"></i>{{ record.comment_count }}</span>
                        <span><i class="fa fa-star-o"></i>{{ record.collects_count }}</span>
                    </div>
                </template>
                <template v-if="column.key === 'abstract'">
                    <span class="gvb_article_abstract">{{ record.abstract }}</span>
                </template>
                <template v-if="column.key === 'tags'">
                    <div class="gvb_article_tags">
                        <a-tag :color="getColor(i)" v-for="(item, i) in record.tags" :key="i">{{ item }}</a-tag>
                    </div>
                </template>
            </template>

            <template #filters>
                <a-select class="gvb_select" v-model:value="tag" style="width: 200px" allowClear @change="onFilter"
                    :options="data.tagOptions" placeholder="筛选文章标签"></a-select>
                <a-select class="gvb_select" v-model:value="category" style="width: 200px" allowClear @change="onFilter"
                    :options="data.categoryOptions" placeholder="筛选文章分类"></a-select>
            </template>
        </GVBTable>
    </div>
</template>

<script setup>

// import { getCategoryListApi } from "@/api/article_api";s
import { reactive, ref } from "vue";
import GVBTable from "@/components/admin/gvb_table.vue"
import { getTagNameListApi } from "@/api/tag_api";

const tag = ref(null)
const category = ref(null)
const gvbTable = ref(null)
const data = reactive({
    list: [
        {
            "abstract": "需要一个全局变量，用于保存配置文件，存放在global目录下",
            "banner_id": 3,
            "banner_url": "uploads/file/29.jpg",
            "category": "p1",
            "collects_count": 0,
            "comment_count": 0,
            "created_at": "2023-02-24 21:21:29",
            "digg_count": 0,
            "id": "P46Vg4YB6uoytGZA-bE7",
            "look_count": 0,
            "tags": [
                "golang",
                "后端",
                "教程",
            ],
            "title": "后端项目搭建",
            "updated_at": "2023-02-24 21:31:13",
            "user_avatar": "/uploads/avatar/default.png",
            "user_id": 1,
            "user_nick_name": "admin"
        },
    ],
    columns: [
        { title: 'id', dataIndex: 'id', key: 'id' },
        // { title: 'Index', dataIndex: 'index', key: 'index' },
        { title: '文章标题', dataIndex: 'title', key: 'title' },
        { title: '文章分类', dataIndex: 'category', key: 'category' },
        { title: '文章简介', dataIndex: 'abstract', key: 'abstract' },
        { title: '作者', dataIndex: 'user_nick_name', key: 'user_nick_name' },
        { title: '封面', dataIndex: 'banner_url', key: 'banner_url' },
        { title: '文章数据', dataIndex: 'data', key: 'data' },
        // {title: '浏览量', dataIndex: 'look_count', key: 'look_count'},
        // {title: '点赞', dataIndex: 'digg_count', key: 'digg_count'},
        // {title: '评论', dataIndex: 'comment_count', key: 'comment_count'},
        // {title: '收藏', dataIndex: 'collects_count', key: 'collects_count'},
        { title: '标签', dataIndex: 'tags', key: 'tags' },
        { title: '发布时间', dataIndex: 'created_at', key: 'created_at' },
        { title: '操作', dataIndex: 'action', key: 'action' },
    ],
    tagOptions: [],
    categoryOptions: []
})

const colorList = ["red", "blue", "green", "purple", "cyan", "orange", "pink"]

function getColor(index) {
    return colorList[index]
}

function onFilter() {
    gvbTable.value.ExportList({ tag: tag.value, key: category.value })
}

async function getData() {
    let res = await getTagNameListApi()
    data.tagOptions = res.data
    // let c = await getCategoryListApi()
    // data.categoryOptions = c.data
}

getData()

</script>

<style lang="scss">
.gvb_article_data {
    span {
        margin-right: 7px;

        i {
            margin-right: 3px;
        }

        &:last-child {
            margin-right: 0;
        }
    }
}

.gvb_article_abstract {
    max-width: 12rem;
    overflow-x: hidden;
    display: inline-block;
    white-space: nowrap;
    text-overflow: ellipsis;
}

.gvb_article_tags {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-column-gap: 10px;
    grid-row-gap: 10px;
    justify-items: self-start;
}

.gvb_article_title {
    em {
        color: red;
        margin-right: 1px;
    }
}
</style>