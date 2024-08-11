<template>
    <md-editor v-model="text" @onUploadImg="onUploadImg" />
</template>

<script setup>
import { ref } from 'vue';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { uploadImageApi } from '@/api/image_api';

const text = ref(''); // 初始化空字符串
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
</script>