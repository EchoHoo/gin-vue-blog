<template>
    <div>
        
        <a-modal :title="title" v-model:open="visible" class="chat-modal" width="600px">
            <div class="chat-body">
                <div v-for="message in data.data" :key="message.id"
                     :class="{'message-right': IsMe(message), 'message-left': !IsMe(message)}">
                    <div v-if="IsMe(message)" class="message-wrap message-right">
                        <div class="message-content">
                            <div class="message-text">{{ message.content }}</div>
                            <div class="message-time">{{ new Date(message.created_at).toLocaleTimeString() }}</div>
                        </div>
                        <div class="avatar">
                            <img :src="message.send_user_avatar" alt="user avatar"/>
                        </div>
                    </div>
                    <div v-else class="message-wrap message-left">
                        <div class="avatar">
                            <img :src="message.rev_user_avatar" alt="user avatar"/>
                        </div>
                        <div class="message-content">
                            <div class="message-text">{{ message.content }}</div>
                            <div class="message-time">{{ new Date(message.created_at).toLocaleTimeString() }}</div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="message-input">
                <a-input v-model="messageContent" @input="handleInput"  placeholder="输入消息..."></a-input>
                <a-button type="primary" @click="sendMessage">发送</a-button>
                <!-- <span>{{ messageContent }}</span> -->
            </div>
        </a-modal>

        <GVBTable :columns="data.columns" baseURL="/api/messages_all" ref="gvbTable" :page-size="7"
            :like-title="'搜索消息列表'" :is-add="false" :is-delete="false" :is-select="false" default-delete>
            <template #cell="{ column, record }">
                <template v-if="column.key === 'send_user'">
                    <img :src="record.send_user_avatar" height="40" style="border-radius: 50%;" alt="">
                    <span style="margin-left: 15px;">{{ record.send_user_nick_name }}</span>
                </template>
                <template v-if="column.key === 'rev_user'">
                    <img :src="record.rev_user_avatar" height="40" style="border-radius: 50%;" alt="">
                    <span style="margin-left: 15px;">{{ record.rev_user_nick_name }}</span>
                </template>
            </template>
            <template #edit="{ record }">
                <a-button type="primary" @click="applyModal(record)">回复</a-button>
              
            </template>

        </GVBTable>
    </div>
</template>

<script setup>
import { getMessageRecordApi, sendMessageApi } from '@/api/user_center_api';
import GVBTable from '@/components/admin/gvb_table.vue';
import { useStore } from '@/stores/store';
import { message } from 'ant-design-vue';
import { nextTick, reactive, ref } from 'vue';

const gvbTable = ref(null);
const sendMessageData = reactive({
    send_user_id: 0,
    rev_user_id: 0,
    content: "",
})
const messageContent = ref("");
const store = useStore();
const data = reactive({
    list: [
        {
            "id": 2,
            "created_at": "2024-08-16T10:07:26.533+08:00",
            "send_user_id": 1,
            "send_user_nick_name": "xiaohei",
            "send_user_avater": "/uploads/avatar/default.jpg",
            "rev_user_id": 2,
            "rev_user_nick_name": "xiaobai",
            "rev_user_avatar": "/uploads/avatar/default.jpg",
            "is_read": false,
            "content": "你好1",
            "is_me":false,
        },
    ],
    columns: [
        // { title: 'id', dataIndex: 'id', key: 'id' },

        // { title: '发送人昵称', dataIndex: 'send_user_nick_name', key: 'send_user_nick_name' },
        { title: '发送人', dataIndex: 'send_user', key: 'send_user' },
        // { title: '发送人头像', dataIndex: 'send_user_avater', key: 'send_user_avater' },
        // { title: '接收人昵称', dataIndex: 'rev_user_nick_name', key: 'rev_user_nick_name' },
        { title: '接收人', dataIndex: 'rev_user', key: 'rev_user' },
        // { title: '接收人头像', dataIndex: 'rev_user_avatar', key: 'rev_user_avatar' },
        { title: '消息内容', dataIndex: 'content', key: 'content' },
        // { title: '消息总数', dataIndex: 'message_count', key: 'message_count' },
        { title: '消息时间', dataIndex: 'created_at', key: 'created_at' },
        { title: '操作', dataIndex: 'action', key: 'action' },
    ],

})
const visible = ref(false);
const title = ref('');
async function applyModal(record) {
    visible.value = true

    // 与某某聊天
    // 判断发送方 和 接收方哪一个不是自己
    let user_id = null
    let flag = false
    if (record.send_user_id === store.userInfo.user_id) {
        // 我是发送方
        title.value = '与' + record.rev_user_nick_name + '聊天';
        user_id = record.rev_user_id
        flag = true
    }
    if (record.rev_user_id === store.userInfo.user_id) {
        // 我是接收方
        title.value = '与' + record.send_user_nick_name + '聊天';
        user_id = record.send_user_id
        flag=false
    }
    sendMessageData.send_user_id = store.userInfo.user_id
    sendMessageData.rev_user_id = user_id
    sendMessageData.content = messageContent.value;
    let res = await getMessageRecordApi(user_id)
    Object.assign(data,res)
    console.log(data)
  
}
async function getData(id){
    let res = await getMessageRecordApi(id)
    Object.assign(data,res)
}


function handleInput(event) {
    sendMessageData.content = event.target.value;
}
function IsMe(msg) {
    return store.userInfo.user_id === msg.send_user_id
}
async function sendMessage(){
    let res = await sendMessageApi(sendMessageData)
    if (res.code){
        message.error(res.msg)
        return
    }
    message.success(res.msg)
    messageContent.value = ''
    getData(sendMessageData.send_user_id)
    gvbTable.value.ExportList()
}
</script>

<style scoped>
.chat-modal .chat-body {
    height: 500px;
    overflow-y: auto;
    background-color: #f2f0f0;
}

.message-right, .message-left {
    display: flex;
    align-items: flex-start;
    margin: 10px;
}

.message-right {
    justify-content: flex-end; /* 发送方消息向右对齐 */
}

.message-left {
    justify-content: flex-start; /* 接收方消息向左对齐 */
}

.avatar {
    width: 40px;
    height: 40px;
    border-radius: 20px;
    margin: 0 10px;
}

.message-content {
    max-width: 70%;
    padding: 8px 12px;
    border-radius: 5px;
    position: relative;
    background-color: #9fdf9f; /* 发送方消息颜色 */
}

.message-left .message-content {
    background-color: #ffffff; /* 接收方消息颜色 */
}

.message-text {
    word-wrap: break-word;
}

.message-time {
    font-size: 12px;
    color: #888;
    position: absolute;
    bottom: -18px;
    right: 0;
}

.message-input {
    display: flex;
    margin-top: 10px;
    padding: 10px;
    background-color: #fff;
}

.message-input a-input {
    flex: 1;
    margin-right: 10px;
}
img{
    width: 40px;
    height: 40px;
    border-radius: 50%;
}

</style>