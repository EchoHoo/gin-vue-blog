<template>
    <div>
        <a-modal title="完善文章信息" @cancel="dialogShow" v-model:open="props.visible" @ok="okHandler">
            <a-form :model="data" name="basic" ref="formRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }"
                autocomplete="off">
                <a-form-item label="文章标题" name="title" has-feedback
                    :rules="[{ required: true, message: '请输入文章标题', trigger: 'blur' }]">
                    <a-input v-model:value="data.title" placeholder="文章标题" />
                </a-form-item>
                <a-form-item label="文章简介">
                    <a-textarea v-model:value="data.abstract" :auto-size="{ minRows: 2, maxRows: 5 }"
                        placeholder="文章简介" />
                </a-form-item>
                <a-form-item label="文章分类">
                    <a-auto-complete v-model:value="data.category" :options="initData.categoryList"
                        placeholder="文章分类" />
                </a-form-item>
                <a-form-item label="文章标签">
                    <a-select class="gvb_select" v-model:value="data.tags" allowClear mode="tags"
                        :options="initData.tagList" placeholder="文章标签"></a-select>
                </a-form-item>
                <a-form-item label="文章封面">
                    <a-select ref="select" v-model:value="data.banner_id" placeholder="选择banner">
                        <a-select-option :value="item.id" v-for="item in initData.bannerList" :key="item.id">
                            <img :src="item.path" alt="" height="30" style="border-radius: 5px; margin-right: 10px">
                            <span>{{ item.name }}</span>
                        </a-select-option>
                        <template #tagRender="{ value: val, label, closable, onClose, option }">
                            <img :src="getLabel(label)" height="30" style="border-radius: 5px; margin-right: 5px" />
                        </template>
                    </a-select>
                </a-form-item>
                <a-form-item label="原文地址">
                    <a-input v-model:value="data.link" placeholder="原文地址" />
                </a-form-item>
                <a-form-item label="文章来源">
                    <a-input v-model:value="data.source" placeholder="文章来源" />
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>
<script setup>

import { reactive, ref } from 'vue';
import 'md-editor-v3/lib/style.css';
import { imageNameListApi, } from '@/api/image_api';
import { getCategoryListApi } from '@/api/article_api';
import { getTagNameListApi } from '@/api/tag_api';
const formRef = ref(null)
const data = reactive({
    title: "",
    abstract: "",
    banner_id: null,
    category: "",
    tags: [],
    link: "",
    source: "",

})
const initData = reactive({
    tagList: [],
    categoryList: [],
    bannerList: [],
})
const props = defineProps({
    visible: {
        type: Boolean,
        default: true
    },
    state: {
        type: Object,
        default: {
            title: "",
            abstract: "",
            banner_id: null,
            category: "",
            tags: [],
            link: "",
            source: "",
        }
    },
    isEdit: {
        type: Boolean,
        default: false
    },
    initDataState: {
        type: Object,
        default: {
            tagOptions: [],
            categoryOptions: [],
            bannerOptions: [],
        }
    }
})
const emit = defineEmits(['update:visible', "ok"])
const dialogShow = () => {
    emit('update:visible', false)
}
function okHandler() {
    try {
        formRef.value.validate()
    } catch (e) {
        return
    }
    emit('ok', data)
}
function getLabel(label) {
    return label[0].props.src
}
async function getData() {
    if (!props.isEdit) {
        // 添加文章
        let t1 = await getTagNameListApi()
        initData.tagList = t1.data
        let t2 = await getCategoryListApi()
        initData.categoryList = t2.data
        let t3 = await imageNameListApi()
        const list = t3.data
        initData.bannerList = list
        // 随机选择一张封
      
        const banner = list[Math.floor(Math.random() *  list.length)]
        data.banner_id = banner.id
    }
    // 编辑文章
    if (props.isEdit) {
        initData.tagList = props.initDataState.tagOptions
        initData.categoryList = props.initDataState.categoryOptions
        initData.bannerList = props.initDataState.bannerOptions
        Object.assign(data, props.state)
    }
}
getData()
</script>
<style scoped></style>