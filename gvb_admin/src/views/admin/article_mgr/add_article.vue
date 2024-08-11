<template>
    <md-editor v-model="data.content" @onUploadImg="onUploadImg" @on-save="onSave" />
</template>

<script setup>
import { onUnmounted, reactive, ref } from 'vue';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { uploadImageApi } from '@/api/image_api';

const onUploadImg = async (files, callback) => {
    const res = await Promise.all(
        files.map((file) => {
            return uploadImageApi(file);
        })
    );
    // 处理返回结果，提取文件路径
    const fileUrls = res.map((item) => {
        const data = item.data;
        if (Array.isArray(data) && data.length > 0) {
            return data[0].file_name; // 假设每次只上传一个文件
        }
        return '';
    }).filter(url => url !== ''); // 过滤掉空字符串
    callback(fileUrls);
};
const data = reactive({
    content: "hello"
})


function ctrlSave(e) {
    if (e.ctrlKey && e.key === 's') {
        showDrawer();
        e.preventDefault();
        onSave(text.value, '');
    }
}
function showDrawer() {
    visible.value = true
}
window.addEventListener('keydown', ctrlSave);
onUnmounted(() => {
    window.removeEventListener('keydown', ctrlSave);
});
// ctrl + s 保存 md是md原文，h是转换后的
const onSave = (md, h) => {
    showDrawer();
}
</script>