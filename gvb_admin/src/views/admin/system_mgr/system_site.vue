<template>
    <div class="gvb_settings_bg">
        <div class="gvb_settings_box">
            <div class="gvb_settings_head">
                <h3>网站设置</h3>
            </div>
            <div class="gvb_setting_body">
                <a-form
                    :model="state"
                    name="basic"
                    ref="formRef"
                    :label-col="{ span: 3 }"
                    :wrapper-col="{ span: 21 }"
                    label-align="left"
                    autocomplete="off"
                >
                    <a-form-item label="网站标题" name="title" has-feedback :rules="[{}]">
                        <a-input v-model:value="state.title" placeholder="请输入网站标题" />
                    </a-form-item>
                    <a-form-item label="slogan" name="slogan" has-feedback :rules="[{}]">
                        <a-input v-model:value="state.slogan" placeholder="请输入网站口号" />
                    </a-form-item>
                    <a-form-item label="英文slogan" name="slogan_en" has-feedback :rules="[{}]">
                        <a-input v-model:value="state.slogan_en" placeholder="请输入英文网站口号" />
                    </a-form-item>
                    <a-form-item label="备案" name="bei_an" has-feedback :rules="[{}]">
                        <a-input v-model:value="state.bei_an" placeholder="请输入备案信息" />
                    </a-form-item>
                    <a-form-item label="网站版本" name="version" has-feedback :rules="[{}]">
                        <a-input v-model:value="state.version" placeholder="请输入网站版本" />
                    </a-form-item>
                    <a-form-item label="建站日期" name="created_at" has-feedback :rules="[{}]">
                        <!-- <a-data-picker v-model:value="state.create_at" placeholder="请选择建站日期" /> -->
                        <a-date-picker v-model:value="createData" placeholder="请选择建站日期" />
                        <!-- <a-input v-model:value="state.create_at" placeholder="请输入建站日期" /> -->
                    </a-form-item>
                    <a-form-item label="QQ" name="qq_image" has-feedback :rules="[{}]">
                        <a-input v-model:value="state.qq_image" placeholder="请输入QQ号" />
                        <img :src="state.qq_image" width="180" height="180" style="margin-top: 10px;" alt="">
                    </a-form-item>
                    <a-form-item label="Wechat" name="wechat_image" has-feedback :rules="[{}]">
                        <a-input v-model:value="state.wechat_image" placeholder="请输入微信号" />
                        <img :src="state.wechat_image" width="180" height="180" style="margin-top: 10px;" alt="">
                    </a-form-item>
                </a-form>
            </div>
            <div class="gvb_setting_btn">
                <a-button type="primary" @click="update()">更新</a-button>
            </div>
        </div>
    </div>
</template>


<script setup>
import { reactive, ref } from 'vue';
import dayjs from 'dayjs';
import { getFormatDateYMD } from '@/utils/data';
const formRef = ref(null);

const state = reactive({
    create_at: "",
    bei_an: "闽ICP备2024031099号-1",
    title: "Obto的个人博客",
    qq_image: "https://gitee.com/guo-cheng-hao/drawing-bed/raw/master/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/qq_image.png",
    version: "6.0.1",
    email: "Gch030819@163.com",
    wechat_image: "https://gitee.com/guo-cheng-hao/drawing-bed/raw/master/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/wechat_image.png",
    name: "123123",
    job: "全栈",
    ddr: "福建福清",
    slogan: "Obto",
    slogan_en: "ObtoBlog",
    web: "http://new.qq.com/omn/20220418/20220418A0C42W00.html",
    bilibili_url: "http://new.qq.com/omn/20220418/20220418A0C42W00.html",
    gitee_url: "http://new.qq.com/omn/20220418/20220418A0C42W00.html",
    github_url: "http://new.qq.com/omn/20220418/20220418A0C42W00.html"
});
const createData = ref(dayjs('2023-08-19','YYYY-MM-DD'))
function update(){
    // 对时间进行转换
    state.create_at = getFormatDateYMD(createData.value.toDate())
}
</script>

<style lang="scss">
.gvb_settings_bg {
    background-color: var(--card_bg);
    min-height: calc(100vh - 130px);
    border-radius: 5px;
    display: flex;
    justify-content: center;

    .gvb_settings_box {
        margin: 20px;
        width: 1000px;
    }

    .gvb_settings_head {
        font-size: 18px;
        display: flex;
        align-items: center;

        &::before {
            width: 2px;
            height: 1.5rem;
            display: inline-block;
            content: "";
            margin-right: 4px;
            background-color: var(--active);
        }
    }
}
</style>
