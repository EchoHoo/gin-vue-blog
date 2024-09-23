<template>
    <div class="base_view chat_group_view">
        <GVBNav is_show></GVBNav>
        <div class="gvb_base_container" :style="{ height: data.allHeight }">
            <div class="gvb_inner_container">
                <div class="gvb_chat_head">
                    <div class="title">在线聊天室</div>
                    <div class="people_num">在线人数： 1</div>
                </div>
                <div class="gvb_chat_body" ref="chatBody" :style="{ height: data.setHeight }">

                    <div class="message" v-for="(item, index) in data.message_list" :key="index">
                        <template v-if="item.msg_type === inRoomMsg">
                            <div class="in_room_msg">
                                <a-tooltip placement="top">
                                    <template #title>
                                        <span>{{ getFormatDate(item.created_at) }}</span>
                                    </template>
                                    <span>{{ item.content }}</span>
                                </a-tooltip>

                            </div>
                        </template>
                        <template v-if="item.msg_type === outRoomMsg">
                            <div class="out_room_msg">
                                <a-tooltip placement="top">
                                    <template #title>
                                        <span>{{ getFormatDate(item.created_at) }}</span>
                                    </template>
                                    <span>{{ item.content }}</span>
                                </a-tooltip>
                            </div>
                        </template>
                        <template v-if="item.msg_type === textMsg">
                            <div :class="{ me: item.is_me }">
                                <div class="avatar">
                                    <img :src="item.avatar" alt="">
                                </div>
                                <div class="message-main">
                                    <div class="message-user">{{ item.nick_name }}</div>
                                    <div class="message-content">
                                        <div class="content">
                                            <div class="txt-message">{{ item.content }}</div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </template>


                    </div>

                </div>
                <div class="gvb_chat_menu">
                    <span><i class="fa fa-cloud-upload"></i></span>
                    <span><i class="fa fa-file-image-o"></i></span>
                    <span><i class="fa fa-file-audio-o"></i></span>
                </div>
                <div class="gvb_chat_footer">
                    <a-textarea v-model:value="data.content" placeholder="发送你的消息，聊起来吧" @keydown.ctrl.enter="sendMessage"
                        :auto-size="{ minRows: 4, maxRows: 4 }"></a-textarea>
                    <a-button @click="sendMessage" class="chat_btn">发送</a-button>
                </div>
            </div>
        </div>
        <!--    <GVBFooter></GVBFooter>-->
    </div>
</template>

<script setup>
import GVBNav from "@/components/gvb_nav.vue"

import { reactive, ref } from "vue";
import { chatGroupApi } from "@/api/chat_group_api";
import { getFormatDate } from "@/utils/data";

let socket = null
let index = 0

// allHeight - 60 - 20 -60 - 30 - 148  - 2 - 20
/*
InRoomMsg  ctype.MsgType = 1 // 进入聊天室
TextMsg    ctype.MsgType = 2 // 发文本消息
ImageMsg   ctype.MsgType = 3 // 图片消息
VoiceMsg   ctype.MsgType = 4 // 语音消息
VideoMsg   ctype.MsgType = 5 // 视频消息
SystemMsg  ctype.MsgType = 6 // 系统消息
OutRoomMsg ctype.MsgType = 7 // 退出聊天室
 */
const inRoomMsg = 1;
const textMsg = 2;


const systemMsg = 6;
const outRoomMsg = 7;
const chatBody = ref(null)

const data = reactive({
    allHeight: "100vh",
    setHeight: "597px",
    message_list: [],
    user_info: {
        nick_name: "",
        avatar: "",
    },
    content: ""
})


async function getData() {
    // console.log(document.documentElement.offsetHeight)
    // console.log(document.documentElement.clientHeight)
    // console.log(document.documentElement.scrollHeight)
    let allHeight = document.documentElement.scrollHeight
    data.allHeight = allHeight + "px"
    let setHeight = allHeight - 340
    data.setHeight = setHeight + "px"

    // 加载聊天记录
    let res = await chatGroupApi({ limit: 50 })
    data.message_list = res.data.list
    data.message_list.reverse()
    // 建立websocket连接

    let websocketURL = import.meta.env.VITE_WEBSOCKET
    console.log("WebSocket URL:", websocketURL); // 添加这一行

    socket = new WebSocket("ws://127.0.0.1:10800/ws" + "/chat_groups")
    // 接收消息
    socket.onmessage = messageApply

    // 连接成功之后的回调
    socket.onopen = function (ev) {
        var _data = {
            content: "",
            msg_type: inRoomMsg,
        }
        socket.send(JSON.stringify(_data));
        console.log("onopen: ", ev)
    }
    socket.onclose = function (ev) {
        var _data = {
            content: "",
            msg_type: outRoomMsg,
        }
        socket.send(JSON.stringify(_data));
        console.log("onopen: ", ev)
    }
    // 错误
    socket.onerror = function (ev) {
        console.log("onerror: ", ev)
    }
    // 服务端关闭
    socket.onclose = function (ev) {
        console.log("onclose: ", ev)
    }
}

function messageApply(event) {
    let _data = event.data
    let jsonData = JSON.parse(_data)
    if (index === 0) {
        data.user_info.nick_name = jsonData.nick_name
        index++
        return
    }
    pushMessage(jsonData)
    index++
}

function pushMessage(jsonData) {
    console.log(jsonData)
    if (jsonData.nick_name === data.user_info.nick_name) {
        jsonData.is_me = true
    }
    data.message_list.push(jsonData)
    data.message_list = [...data.message_list]; // 通过数组解构重新赋值
    // data.message_list = [...data.message_list, jsonData];
}

function sendMessage() {
    let _data = {
        content: data.content,
        msg_type: textMsg,
    }

    if (socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify(_data));
        data.content = "";
    } else {
        console.error("WebSocket is not open. Current state: " + socket.readyState);
    }

    setTimeout(() => {
        let top = chatBody.value.scrollHeight  // 10000
        let h = chatBody.value.clientHeight
        let timer = null
        timer = setInterval(() => {
            let newTop = chatBody.value.scrollTop
            if (top <= newTop + h) {
                clearInterval(timer)
                return
            }
            chatBody.value.scrollTop += 20
        }, 5)

        // chatBody.value.scrollTop = chatBody.value.scrollHeight
    }, 100)


}


getData()

</script>

<style lang="scss">
.chat_group_view {
    .gvb_base_container {
        height: 100vh;

        .gvb_inner_container {
            margin-top: 80px;
        }
    }

    .gvb_chat_head {
        height: 60px;
        width: 100%;
        background-color: var(--card_bg);
        border-radius: 5px 5px 0 0;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;

        .title {
            font-size: 16px;
        }

        .people_num {
            font-size: 12px;
        }
    }

    .gvb_chat_body {
        margin-top: 1px;
        height: 597px;
        background-color: var(--card_bg);
        overflow-y: auto;
        padding: 10px 20px;

        &::-webkit-scrollbar {
            width: 8px
        }

        &::-webkit-scrollbar-thumb {
            border-radius: 10px;
            background-color: #e0dddd
        }

        &::-webkit-scrollbar-button {
            background-color: transparent
        }


        .message {
            margin-bottom: 20px;

            .in_room_msg,
            .out_room_msg {
                display: flex;
                justify-content: center;

                span {
                    background-color: #e2e2e2;
                    border-radius: 5px;
                    color: #494949;
                    font-size: 12px;
                    padding: 2px 10px;
                    cursor: pointer;
                }
            }

            .out_room_msg {
                span {
                    background-color: #f7e5e5;
                }
            }

            >div {
                display: flex;

                .avatar {
                    img {
                        width: 40px;
                        height: 40px;
                        border-radius: 50%;
                    }
                }

                .message-main {
                    margin-left: 10px;

                    .message-content {
                        margin-top: 5px;

                        .content {


                            display: flex;
                            margin-left: 10px;
                            justify-content: left;
                        }

                        .txt-message {
                            background-color: #d2d2d2;
                            padding: 10px;
                            position: relative;
                            width: fit-content;
                            min-height: 41px;
                            white-space: pre;
                            border-radius: 5px;
                        }

                        .txt-message::before {
                            content: "";
                            display: block;
                            position: absolute;
                            left: -32px;
                            top: 6px;
                            border-width: 10px 20px;
                            border-style: solid;
                            border-color: transparent #d2d2d2 transparent transparent;
                        }

                    }
                }
            }

            >.me {
                justify-content: right;
                flex-direction: row-reverse;

                .message-main {
                    .message-content {
                        .content {
                            display: flex;
                            margin-right: 10px;
                            justify-content: right;
                        }

                        .txt-message::before {
                            right: -31px;
                            left: inherit;
                            top: 6px;
                            border-width: 10px 20px;
                            border-style: solid;
                            border-color: transparent transparent transparent #d2d2d2;
                        }
                    }
                }


            }
        }

    }

    .gvb_chat_menu {
        margin-top: 1px;
        height: 30px;
        background-color: var(--card_bg);
        display: flex;
        align-items: center;

        span {
            padding: 5px;
            width: 30px;
            height: 30px;
            display: flex;
            justify-content: center;
            align-items: center;

            &:hover {
                background-color: #f0f0f0;
            }
        }
    }

    .gvb_chat_footer {
        position: relative;
        background-color: var(--card_bg);
        border-radius: 0 0 5px 5px;
        height: 148px;

        .ant-input {
            border: none;
            outline: none;

            &:focus {
                border: none;
                outline: none;
                box-shadow: none;
            }
        }

        .chat_btn {
            position: absolute;
            bottom: 10px;
            left: 10px;
        }
    }
}
</style>