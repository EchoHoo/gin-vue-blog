<template>

    <GVBCCard title="站点信息" class="site_card">
        <div><b>建站时间</b><span>{{ getFormatDateYMD(store.siteInfo.create_at,"-") }}</span></div>
        <div><b>运行时间</b><span>{{ data.date }}</span></div>
        <div><b>文章总数</b><span>{{ data.article_count }}</span></div>
        <div><b>用户总数</b><span>{{ data.user_count }}</span></div>
        <div>
            <div class="my_image">
                <div class="qq_image">
                    <img src="https://gitee.com/guo-cheng-hao/drawing-bed/raw/master/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/qq_image.png"
                        alt="">
                </div>
                <div class="wechat_image">
                    <img src="https://gitee.com/guo-cheng-hao/drawing-bed/raw/master/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/wechat_image.png"
                        alt="">
                </div>
            </div>

        </div>
    </GVBCCard>
</template>

<script setup>
import { getDataSumApi } from "@/api/data_api";
import GVBCCard from "@/components/gvb_card.vue"
import { useStore } from "@/stores/store";
import { getFormatDateYMD } from "@/utils/data";
import { message } from "ant-design-vue";
import { onMounted, onUnmounted, reactive } from "vue";
const store = useStore();

const data = reactive({
    date: "",
    look_count: 2003412,
    article_count: 45,
    user_count:100,
})

onMounted(()=>{
    const intervalId = setInterval(getTimeStamp, 1000);  // 设置定时器，每秒更新一次时间
    onUnmounted(() => {
         clearInterval(intervalId);  // 清除定时器，避免内存泄漏
    });
})
// 动态获取一个事件戳，距离现在的年月日时分秒
function getTimeStamp(){
    let date = store.siteInfo.create_at;
    let oldData = new Date(date).getTime();
    let newData = new Date().getTime();
    let diffTime = newData - oldData;
    let diffDays = Math.floor(diffTime / (24 * 3600 * 1000));
    let diffHours = Math.floor((diffTime % (24 * 3600 * 1000)) / (3600 * 1000));
    let diffMinutes = Math.floor((diffTime % (3600 * 1000)) / (60 * 1000));
    let diffSeconds = Math.floor((diffTime % (60 * 1000)) / 1000);
    data.date = `${diffDays}天${diffHours}时${diffMinutes}分${diffSeconds}秒`;
}

async function getArticleCount(){
    let res = await getDataSumApi();
    if (res.code){
        message.error(res.msg);
        return;
    }
    data.article_count = res.data.article_count;
    data.user_count = res.data.user_count;
 
}
getArticleCount();
</script>

<style lang="scss">
.site_card {
    .body {
        >div {
            font-size: 16px;
            margin-bottom: 8px;

            b {
                margin-right: 5px;

                &::after {
                    content: ":";
                }
            }
        }

        .my_image {
            display: flex;
            margin-top: 10px;
            justify-content: space-around;

            >div {
                width: 50%;
                display: flex;
                flex-direction: column;
                align-items: center;

                img {
                    width: 100px;
                    height: 100px;
                }
            }
        }
    }
}
</style>