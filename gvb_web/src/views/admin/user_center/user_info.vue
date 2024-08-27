<template>
  <div class="gvb_user_info_bg">

    <a-modal title="绑定邮箱" v-model:open="bindEmailVisible">
      <div>
        <a-steps :current="step">
          <a-step v-for="item in steps" :key="item.title" :title="item.title" :description="item.description" />
        </a-steps>
        <a-form :model="formState" name="basic" ref="formRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }"
          autocomplete="off">
          <templete v-if="step === 0">
            <a-form-item label="邮箱地址" name="email" has-feedback
              :rules="[{ required: true, message: '请输入邮箱地址' }, { validator: validateEmail, message: '非法邮箱', trigger: 'blur' }]">
              <a-input v-model:value="formState.email" placeholder="请输入邮箱"></a-input>
            </a-form-item>
          </templete>
          <templete v-if="step === 1">
            <a-form-item label="密码">
              <a-input v-model:value="formState.password" placeholder="请输入密码"></a-input>
            </a-form-item>
            <a-form-item label="验证码">
              <a-input v-model:value="formState.code" placeholder="请输入验证码"></a-input>
            </a-form-item>
          </templete>
          <templete v-if="step === 2">
            <a-form-item label="绑定成功">
              <span>绑定成功</span>
            </a-form-item>
          </templete>
        </a-form>
      </div>
      <template #footer>
        <a-button v-if="step === 1" @click="bindEmailCache">取消</a-button>
        <a-button type="primary" v-if="step === 0" @click="sendEmailCode">下一步</a-button>
        <a-button type="primary" v-if="step === 1" @click="step--">上一步</a-button>
        <a-button v-if="step === 1" @click="bindEmail">确认</a-button>
      </template>
    </a-modal>

    <a-modal title="修改密码" v-model:open="updatePasswordVisible" @ok="updatePassword">
      <a-form :model="pwdState" name="basic" ref="pwdformRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }"
        autocomplete="off">
        <a-form-item label="旧密码" name="old_pwd" has-feedback>
          <a-input-password v-model:value="pwdState.old_pwd" placeholder="请输入旧密码"></a-input-password>
        </a-form-item>

        <a-form-item label="新密码" name="new_pwd" has-feedback>
          <a-input-password v-model:value="pwdState.pwd" placeholder="请输入新密码"></a-input-password>
        </a-form-item>
        <a-form-item label="确认密码" name="re_pwd" has-feedback
          :rules="[{ validator: validatePassword, message: '两次密码不一致', trigger: 'blur' }]">
          <a-input-password v-model:value="pwdState.re_pwd" placeholder="确认密码"></a-input-password>
        </a-form-item>
      </a-form>
    </a-modal>

    <div class="gvb_user_info_view">
      <div class="user_head">
        个人信息
      </div>
      <div class="body user_info_view">
        <a-form :label-col="{ span: 2 }" :wrapper-col="{ span: 22 }">
          <a-form-item label="昵称">
            <a-input @blur="updateUserInfo('nick_name')" placeholder="昵称" v-model:value="userInfo.nick_name"
              class="user_ipt"></a-input>
          </a-form-item>
          <a-form-item label="我的签名">
            <a-textarea @blur="updateUserInfo('sign')" placeholder="我的签名" v-model:value="userInfo.sign" class="user_ipt"
              :auto-size="{ minRows: 2, maxRows: 5 }"></a-textarea>
          </a-form-item>
          <a-form-item label="我的博客">
            <a-input @blur="updateUserInfo('link')" placeholder="我的博客" v-model:value="userInfo.link"
              class="user_ipt"></a-input>
          </a-form-item>
          <a-form-item label="邮箱">
            <span>{{ userInfo.email }}</span>
          </a-form-item>
          <a-form-item label="现地址">
            <span>{{ userInfo.addr }}</span>
          </a-form-item>
          <a-form-item label="当前角色">
            <span>{{ userInfo.role }}</span>
          </a-form-item>
          <a-form-item label="sign_status">
            <span>{{ userInfo.sign_status }}</span>
          </a-form-item>
          <a-form-item label="我的积分">
            <span>{{ userInfo.integral }}</span>
          </a-form-item>
        </a-form>
      </div>
      <div class="user_head">
        每日奖励
      </div>
      <div class="body daily_reward">
        <div>
          <img src="../../../assets/avatar/day_login.png" alt="">
          每日登陆 +2
        </div>
        <div>
          <img src="../../../assets/avatar/day_blog.png" alt="">
          发文 +3
        </div>
        <div>
          <img src="../../../assets/avatar/day_comment.png" alt="">
          发评论 +1
        </div>
        <div>
          <img src="../../../assets/avatar/day_chat_group.png" alt="">
          群聊 +0.5
        </div>
      </div>

      <div class="user_head">
        成就勋章
      </div>
      <ul class="body medal">
        <li>
          <img src="https://i0.hdslb.com/bfs/face/21fdd1bcd9ad1b52d5725026d71c185a68681644.png" alt="">
          <div class="name">初来乍到</div>
        </li>
        <li>
          <img src="https://i0.hdslb.com/bfs/face/51ca16136e570938450bca360f28761ceb609f33.png" alt="">
          <div class="name">有爱萌新</div>
        </li>
        <li>
          <img src="https://i0.hdslb.com/bfs/face/89491e6cee3b1c88292b618a58703ab614665f66.png" alt="">
          <div class="name">老司机</div>
        </li>
      </ul>


      <div class="user_head">
        操作
      </div>
      <div class="body actions">
        <a-button type="primary" @click="bindEmailVisible = true">绑定邮箱</a-button>
        <a-button type="primary" @click="updatePasswordVisible = true">修改密码</a-button>
        <a-button type="primary" danger @click="logout">注销退出</a-button>
      </div>
    </div>
  </div>
</template>
<script setup>
import { updateUserPassWordApi } from '@/api/user_api';
import { bindEmailApi, getUserInfoApi, sendEmailCodeApi, updateUserInfoApi } from '@/api/user_center_api';
import { logout } from '@/utils/logout';
import { message } from 'ant-design-vue';
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const formRef = ref(null)
const pwdformRef = ref(null)
const userInfo = reactive({
  addr: "",
  email: "",
  integral: 0,
  nick_name: "",
  link: "",
  role: "",
  sign: "",
  sign_status: "",
})
const formState = reactive({
  email: "",
  password: "",
  code: "",
})
const state = reactive({
  nick_name: "",
  link: "",
  sign: "",
})
const steps = [{
  title: '绑定邮箱',
  content: 'First-content',
},
{
  title: '发送验证码',
  content: 'Second-content',
}]
const step = ref(0)
const bindEmailVisible = ref(false)
const updatePasswordVisible = ref(false)
const pwdState = reactive({
  old_pwd: "",
  pwd: "",
  re_pwd: "",
})
async function getData() {
  let res = await getUserInfoApi()
  Object.assign(userInfo, res.data)
  state.sign = userInfo.sign
  state.link = userInfo.link
  state.nick_name = userInfo.nick_name
}

async function updateUserInfo(cloumn) {
  const val = userInfo[cloumn]
  const oldVal = state[cloumn]
  if (oldVal === val) {
    return
  }
  const data = {}
  data[cloumn] = val
  let res = await updateUserInfoApi(data)
  if (res.code) {
    message.error(res.msg)
    return
  }
  message.success(res.msg)
}
async function sendEmailCode() {
  let res = await sendEmailCodeApi(formState.email)
  if (res.code) {
    message.error(res.msg)
    return
  }
  message.success(res.msg)
  step.value = 1
}
async function bindEmail() {
  console.log(formState)

  let res = await bindEmailApi(formState)
  if (res.code) {
    message.error(res.msg)
    return
  }
  message.success(res.msg)
  bindEmailCache()
  getData()
}
function bindEmailCache() {
  step.value = 0
  bindEmailVisible.value = false
  formState.email = ""
  formState.code = ""
  formState.password = ""
}


async function updatePassword() {
  try {
    await pwdformRef.value.validate()
  } catch (e) {
    return
  }
  let res = await updateUserPassWordApi(pwdState)
  if (res.code) {
    message.error(res.msg)
    return
  }
  message.success(res.msg)
  updatePasswordVisible.value = false
  setTimeout(() => {
    router.push({ name: "login" })
  }, 500)
}
let validateEmail = async (_rule, value) => {
  const reg = /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/;
  if (reg.test(value)) {
    return Promise.resolve();
  }
  return Promise.reject("请输入正确的邮箱地址");
}
let validatePassword = async (_rule, value) => {

  if (value !== pwdState.pwd) {
    return Promise.reject("两次输入的密码不一致");
  }
  return Promise.resolve();
}

getData()

</script>
<style lang="scss">
.gvb_user_info_bg {
  background-color: var(--card_bg);
  display: flex;
  justify-content: center;
  padding: 20px;

  .gvb_user_info_view {
    width: 1000px;
  }

  .user_head {
    font-size: 18px;
    display: flex;
    align-items: center;
    font-weight: 600;
    margin-bottom: 10px;

    &::before {
      width: 3px;
      height: 1.5rem;
      display: inline-block;
      content: "";
      margin-right: 5px;
      background-color: var(--active);
    }
  }

  .body {
    margin-bottom: 20px;
  }

  .user_info_view {
    .ant-form-item {
      margin: 0 0 5px;
    }

    .user_ipt {
      width: 400px;
    }
  }

  .actions {

    .ant-btn {
      margin-right: 10px;
    }
  }

  .medal {
    display: flex;
    flex-wrap: wrap;

    li {
      margin-right: 20px;

      .name {
        text-align: center;
      }

      &:last-child {
        margin-right: 0;
      }

      img {
        width: 60px;
        height: 60px;
      }
    }
  }

  .daily_reward {
    display: flex;

    >div {
      display: flex;
      flex-direction: column;
      align-items: center;
      margin-right: 50px;
    }

    img {
      width: 60px;
      height: 60px;
    }
  }


}
</style>