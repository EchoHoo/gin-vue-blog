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
                    ref="formSiteRef"
                    :label-col="{ span: 3 }"
                    :wrapper-col="{ span: 21 }"
                    label-align="left"
                    autocomplete="off"
                >
                    <a-form-item label="网站标题" name="title" has-feedback 
                    :rules="[{required:true, message:'请输入网站标题',trigger:'blur'}]">
                        <a-input v-model:value="state.title" placeholder="请输入网站标题" />
                    </a-form-item>
                    <a-form-item label="slogan" name="slogan" has-feedback 
                    :rules="[{required:true, message:'请输入slogan',trigger:'blur'}]">
                        <a-input v-model:value="state.slogan" placeholder="请输入网站口号" />
                    </a-form-item>
                    <a-form-item label="英文slogan" name="slogan_en" has-feedback 
                    :rules="[{required:true, message:'请输入英文slogan',trigger:'blur'}]">
                        <a-input v-model:value="state.slogan_en" placeholder="请输入英文网站口号" />
                    </a-form-item>
                    <a-form-item label="备案" name="bei_an" has-feedback 
                    :rules="[{required:true, message:'请输入备案',trigger:'blur'}]">
                        <a-input v-model:value="state.bei_an" placeholder="请输入备案信息" />
                    </a-form-item>
                    <a-form-item label="网站版本" name="version" has-feedback 
                    :rules="[{required:true, message:'请输入网站版本',trigger:'blur'}]">
                        <a-input v-model:value="state.version" placeholder="请输入网站版本" />
                    </a-form-item>
                    <a-form-item label="建站日期" name="created_at" has-feedback :rules="[{}]">
                        <!-- <a-data-picker v-model:value="state.create_at" placeholder="请选择建站日期" /> -->
                        <a-date-picker v-model:value="createData" placeholder="请选择建站日期" />
                        <!-- <a-input v-model:value="state.create_at" placeholder="请输入建站日期" /> -->
                    </a-form-item>
                    <a-form-item label="QQ" name="qq_image" has-feedback 
                    :rules="[{required:true, message:'请输入QQ',trigger:'blur'}]">
                        <a-input v-model:value="state.qq_image" placeholder="请输入QQ号" />
                        <a-upload
                            name="images"
                            list-type="picture-card"
                            class="gvb_image_upload"
                            :show-upload-list="false"
                            action="/api/images"
                            :headers="{token:store.userInfo.token}"
                            @change="uploadChange($event,'qq')"
                        >
                        <img :src="state.qq_image" width="120" height="120" style="margin-top: 10px;" alt="">   
                        </a-upload>
                    </a-form-item>
                    <a-form-item label="Wechat" name="wechat_image" has-feedback 
                    :rules="[{required:true, message:'请输入Wechat',trigger:'blur'}]">
                        <a-input v-model:value="state.wechat_image" placeholder="请输入微信号" />
                        <a-upload
                            name="images"
                            list-type="picture-card"
                            class="gvb_image_upload"
                            :show-upload-list="false"
                            action="/api/images"
                             :headers="{token:store.userInfo.token}"
                            @change="uploadChange($event,'wechat')"
                        >
                        <img :src="state.wechat_image" width="120" height="120" style="margin-top: 10px;" alt="">
                        </a-upload>
                        
                    </a-form-item>
                </a-form>
            </div>
            <div class="gvb_settings_head">
                <h3>个人信息</h3>
            </div>
            <div class="gvb_setting_body">
                <a-form
                    :model="state"
                    name="basic"
                    ref="formMyInfoRef"
                    :label-col="{ span: 3 }"
                    :wrapper-col="{ span: 21 }"
                    label-align="left"
                    autocomplete="off"
                >
                    <a-form-item label="昵称" name="name" has-feedback 
                    :rules="[{required:true, message:'请输入昵称',trigger:'blur'}]">
                        <a-input v-model:value="state.name" placeholder="请输入昵称" />
                    </a-form-item>
                    <a-form-item label="工作" name="job" has-feedback 
                    :rules="[{required:true, message:'请输入工作',trigger:'blur'}]">
                        <a-input v-model:value="state.job" placeholder="请输入工作" />
                    </a-form-item>
                    <a-form-item label="地址" name="addr" has-feedback 
                    :rules="[{required:true, message:'请输入地址',trigger:'blur'}]">
                        <a-input v-model:value="state.addr" placeholder="请输入地址" />
                    </a-form-item>
                    <a-form-item label="邮箱" name="email" has-feedback 
                    :rules="[{required:true, message:'请输入邮箱',trigger:'blur'}]">
                        <a-input v-model:value="state.email" placeholder="请输入邮箱" />
                    </a-form-item>
                </a-form>
            </div>

            <div class="gvb_settings_head">
                <h3>链接地址</h3>
            </div>
            <div class="gvb_setting_body">
                <a-form
                    :model="state"
                    name="basic"
                    ref="formLinkRef"
                    :label-col="{ span: 3 }"
                    :wrapper-col="{ span: 21 }"
                    label-align="left"
                    autocomplete="off"
                >
                    <a-form-item label="网站地址" name="web" has-feedback 
                            :rules="[{required:true, message:'请输入网站地址',trigger:'blur'}]">
                        <a-input v-model:value="state.web" placeholder="请输入网站地址" />
                    </a-form-item>
                    <a-form-item label="哔哩哔哩" name="bilibili_url" has-feedback 
                            :rules="[{required:true, message:'请输入哔哩哔哩地址',trigger:'blur'}]">
                        <a-input v-model:value="state.bilibili_url" placeholder="请输入哔哩哔哩地址" />
                    </a-form-item>
                    <a-form-item label="Gitee" name="gitee_url" has-feedback 
                            :rules="[{required:true, message:'请输入Gitee地址',trigger:'blur'}]">
                        <a-input v-model:value="state.gitee_url" placeholder="请输入Gitee地址" />
                    </a-form-item>
                    <a-form-item label="GitHub" name="github_url" has-feedback 
                            :rules="[{required:true, message:'请输入GitHub地址',trigger:'blur'}]">
                        <a-input v-model:value="state.github_url" placeholder="请输入GitHub地址" />
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
import { message } from 'ant-design-vue';
import { useStore } from '@/stores/store';
import { getSiteInfoApi, updateSiteInfoApi } from '@/api/system_api';
const formRef = ref(null);

const state = reactive({
    create_at: "",
    bei_an: "",
    title: "",
    qq_image: "",
    version: "",
    email: "",
    wechat_image: "",
    name: "",
    job: "",
    addr: "",
    slogan: "",
    slogan_en: "",
    web: "",
    bilibili_url: "",
    gitee_url: "",
    github_url: ""
});
const store = useStore();
const createData = ref(dayjs('2023-08-19','YYYY-MM-DD'))
async function update(){
    // 对时间进行转换
    state.create_at = getFormatDateYMD(createData.value.toDate())
    let res = await updateSiteInfoApi(state)
    if(res.code){
        message.error(res.msg)
        return
    }
    message.success(res.msg)
    getData()
    return
}
function uploadChange(e,type) {
    if (e.file.status === 'done'){
        let res = e.file.response
        if(res.code){
            message.error(res.msg)
            return
        }
        let path = res.data[0].file_name
        if(!path.startsWith('/')){
            path = '/' + path
        }
        if(type === 'qq'){
            state.qq_image = path
        }
        if(type === 'wechat'){
            state.wechat_image = path
        }
    }
}
const formSiteRef = ref(null)
const formMyInfoRef = ref(null)
const formLinkRef = ref(null)

async function getData(){
    let res = await getSiteInfoApi()
    Object.assign(state,res.data)
    createData.value = dayjs(res.data.create_at)
    try {
        await formSiteRef.value.validate()
        await formMyInfoRef.value.validate()
        await formLinkRef.value.validate()
    }catch(e){
        return
    }
}
getData()
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
    .gvb_image_upload{
        margin-top: 10px;
    }
    .gvb_setting_btn{
        margin-bottom: 20px;
    }

}
</style>
