<template>
    <md-editor ref="editorRef" v-model="content" @on-upload-img="onUploadImg" @on-save="onSave" />
</template>

<script setup>
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { ref, watch, onMounted } from "vue";
import { uploadImageApi } from "@/api/image_api";
import { message } from 'ant-design-vue';

const props = defineProps({
    content: {
        type: String,
    }
})
const content = ref("")
const editorRef = ref(null)

function getData() {
    content.value = props.content
}

getData()

const emit = defineEmits(['update:content', "onSave"])
watch(content, () => {
    emit('update:content', content.value)
})
watch(() => props.content, () => {
    content.value = props.content
}, { immediate: true })


const onUploadImg = async (files, callback) => {
    const res = await Promise.all(
        files.map((file) => {
            return uploadImageApi(file)
        })
    );
    callback(res.map((item) => {
        if (item.code){
            message.error(item.msg)
            return
        }
        message.success("上传成功")
        return item.data
    }));
};

const onSave = (md, h) => {
    emit("onSave", md)
}

onMounted(() => {
    editorRef.value?.focus();
})

</script>

<style></style>