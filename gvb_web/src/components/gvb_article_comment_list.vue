<template>
    <div class="gvb_article_comment_list">
        <div class="gvb_comment_item" v-for="item in data.list" :key="item.model.id">
            <div class="avatar">
                <img :src="item.user.avatar" alt="">
            </div>
            <div class="comment">
                <div class="name">{{ item.user.nick_name }}</div>
                <div class="comment_content">
                    {{ item.content }}
                </div>
                <div class="info">
                    <span class="date">{{ getFormatDate(item.model.created_at) }}</span>
                    <span class="apply" @click="showApply(item.model.id, item.user.nick_name)">回复 ({{
                        item.comment_count }})</span>
                    <a-popconfirm title="确定删除这条评论吗" ok-text="删除" cancel-text="取消"
                        @confirm="removeComment(item.model.id)">
                        <span class="delete" v-if="role === 1 || user_id === item.user_id">删除</span>
                    </a-popconfirm>

                </div>

                <div class="sub_comment_list">
                    <div class="sub_comment_item" v-for="sub in item.sub_comments" v-if="item.sub_comments.length"
                        :key="sub.model.id">
                        <div class="sub_avatar">
                            <img :src="sub.user.avatar" alt="">
                        </div>
                        <div class="sub_comment">
                            <div class="line">
                                <span class="name">{{ sub.user.nick_name }}</span><span class="content">{{ sub.content
                                    }}</span>
                            </div>
                            <div class="info">
                                <span class="date">{{ getFormatDate(sub.model.created_at) }}</span>

                                <a-popconfirm title="确定删除这条评论吗" ok-text="删除" cancel-text="取消"
                                    @confirm="removeComment(sub.model.id)">
                                    <span class="delete" @click="removeComment(sub.model.id)"
                                        v-if="role === 1 || user_id === sub.user_id">删除</span>
                                </a-popconfirm>

                            </div>

                        </div>
                    </div>
                    <div class="sub_comment_apply" v-if="state.parent_comment_id === item.model.id">
                        <img :src="store.userInfo.avatar" alt="" class="user_avatar">
                        <a-input class="comment_ipt" @keydown.ctrl.enter="apply" v-model:value="state.content"
                            :placeholder="'回复 @ ' + state.nick_name"></a-input>
                        <a-button type="primary" @click="apply">回复</a-button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { useRoute } from "vue-router";
import { commentCreateApi, getArticleCommentListApi, commentRemoveApi } from "@/api/comment_api";
import { reactive } from "vue";
import { getFormatDate } from "@/utils/data";
import { useStore } from "@/stores/store";
import { message } from "ant-design-vue";

const store = useStore()
const route = useRoute()
const articleID = route.params.id
const data = reactive({
    list: [
        {
            "article_id": "",
            "comment_count": 0,
            "comment_model": "",
            "content": "",
            "digg_count": 0,
            "parent_comment_id": 0,
            "sub_comments": [],
            "model": {
                "id": "",
                "created_at": "",
            },
            "user": {
                "addr": "",
                "avatar": "",
                "created_at": "",
                "email": "",
                "id": 0,
                "integral": 0,
                "ip": "",
                "link": "",
                "nick_name": "",
                "role": 1,
                "sign": "",
                "sign_status": 1,
                "tel": "",
                "token": "",
                "user_name": ""
            },
            "user_id": 0
        }
    ]
})
const state = reactive({
    article_id: articleID,
    content: "",
    parent_comment_id: null,
    nick_name: ""
})
const role = store.userInfo.role
const user_id = store.userInfo.user_id

async function getData() {
    let res = await getArticleCommentListApi(articleID);
    data.list = formatComments(res.data);
    console.log(data.list);
}
function formatComments(comments) {
    return comments.map(comment => {
        // 格式化每条评论的创建时间
        comment.created_at = getFormatDate(comment.created_at);

        // 如果有子评论，则递归地格式化子评论
        if (comment.sub_comments && comment.sub_comments.length > 0) {
            comment.sub_comments = formatComments(comment.sub_comments);
        }

        return comment;
    });
}
// async function getData() {
//     let res = await getArticleCommentListApi(articleID)
//     data.list = res.data
//     for (var i = 0; i < data.list.length; i++) {    
//         var createAt= getFormatDate(data.list[i].model.created_at)
//         data.list[i].created_at=createAt
//         for (var j = 0; j< data.list[i].sub_comments.length ;j++){
//             getData()
//         }
//     }
//     console.log(data.list)
// }

function showApply(parentID, nick_name) {
    if (store.userInfo.role === 0){
        message.warn("请先登录")
        return
    }
    console.log("Called showApply with parentID:", parentID, "and nickName:", nick_name);
    if (!parentID) {
        console.error("parentID is undefined");
    }
    state.content = ""
    state.parent_comment_id = parentID
    state.nick_name = nick_name
}

async function apply() {
    if (state.content.trim() === "") {
        message.warn("评论内容不可为空")
        return
    }
    let res = await commentCreateApi(state)
    if (res.code) {
        message.error(res.msg)
        return
    }
    message.success(res.msg)
    state.content = ""
    getData()
}

async function removeComment(id) {
    let res = await commentRemoveApi(id)
    if (res.code) {
        message.error(res.msg)
        return
    }
    message.success(res.msg)
    getData()

}

defineExpose({
    getData
})

getData()
console.log(data.list);  // 查看所有评论数据以确认id是否正确设置

</script>

<style lang="scss">
.gvb_article_comment_list {
    .gvb_comment_item {

        display: flex;
        justify-content: space-between;
        margin-bottom: 20px;

        .avatar {
            width: 40px;

            img {
                width: 40px;
                height: 40px;
                display: block;
                border-radius: 50%;
            }

        }

        .comment {
            width: calc(100% - 60px);
            background-color: var(--card_bg);
            border-radius: 5px;
            padding: 10px;

            >.name {
                color: var(--active);
                font-weight: 600;
                font-size: 16px;
            }

            .comment_content {
                color: var(--text);
            }

            .info {
                color: var(--text);

                .date {}

                .apply,
                .delete {
                    cursor: pointer;
                }

                >span {
                    margin-right: 10px;
                }
            }

            .sub_comment_list {
                margin-top: 10px;

                .sub_comment_item {
                    margin-bottom: 10px;
                    display: flex;
                    justify-content: space-between;

                    .sub_avatar {
                        width: 30px;

                        img {
                            width: 30px;
                            height: 30px;
                            border-radius: 50%;
                        }
                    }

                    .sub_comment {
                        width: calc(100% - 40px);

                        .line {
                            .name {
                                color: var(--h2);
                                font-weight: 600;
                            }

                            .content {
                                margin-left: 10px;
                                color: var(--text);
                            }
                        }
                    }
                }

                .sub_comment_apply {
                    display: flex;
                    justify-content: space-between;
                    align-items: center;

                    .user_avatar {
                        width: 30px;
                        height: 30px;
                        border-radius: 50%;
                    }

                    .comment_ipt {
                        margin: 0 10px;
                    }

                }
            }
        }
    }
}
</style>