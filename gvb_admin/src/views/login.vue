<template>
  <div class="gvb_login_bg">
    <img src="../assets/avatar/login.jpg" alt="">

    <div class="gvb_login_container">
      <div class="gvb_login_main">
        <div class="gvb_login_head">用户登录</div>
        <div class="gvb_login_form">
          <div class="gvb_login_form_item">
            <a-input placeholder="用户名" v-model:value="data.user_name">
              <template #prefix>
                <i class="fa fa-user-o"></i>
              </template>
            </a-input>
          </div>
          <div class="gvb_login_form_item">
            <a-input placeholder="密码" type="password" v-model:value="data.password">
              <template #prefix>
                <i class="fa fa-key"></i>
              </template>
            </a-input>
          </div>
          <div class="gvb_login_form_item" @click="emailLogin">
            <a-button type="primary" block>登录</a-button>
          </div>
        </div>
      </div>
    </div>

  </div>

</template>

<script setup>
import { reactive } from 'vue';
import { message } from 'ant-design-vue';
import { emailLoginApi } from '@/api/user_api';
import { parseToken } from '@/utils/jwt';
import { useStore } from '@/stores/store';
import { useRouter } from 'vue-router';
import { useRoute } from 'vue-router';
const route = useRoute()
const router = useRouter()
const store = useStore()
const data = reactive({
  user_name: "",
  password: ""
})
async function emailLogin() {
  if (data.user_name === "" || data.password === "") {
    message.error("用户名或密码不能为空")
    return
  }
  let res = await emailLoginApi(data)
  if (res.code) {
    message.error(res.msg)
    return
  }
  message.success(res.msg)
  let userInfo = parseToken(res.data)
  userInfo.token = res.data
  store.setUserInfo(userInfo)
  const redirect_url = route.query.redirect_url
  if (redirect_url === undefined) {
    setTimeout(() => {
      router.push('admin/home')
    }, 1000)
  }
  // 跳转页面
  setTimeout(() => {
    router.push({path: redirect_url})
  }, 1000)
}
</script>
<style scoped lang="scss">
.gvb_login_bg {
  width: 100%;
  height: 100vh;
  /* 确保背景图覆盖整个视窗高度 */
  background-size: 100vw 100vh;
  /* 背景图固定大小，覆盖整个视窗 */
  background-position: center;
  /* 背景图居中 */
  background-repeat: no-repeat;
  /* 防止背景图重复 */
  overflow: hidden;
  /* 防止滚动 */
  position: relative;
  /* 确保定位正确 */
  display: flex;
  /* 使内部元素能够正确对齐 */
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
}

.gvb_login_container {
  position: fixed;
  right: 0;
  top: 0;
  width: 400px;
  height: 100vh;
  background-color: rgba(white, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;

  .gvb_login_main {
    width: 75%;

    .gvb_login_head {
      font-size: 24px;
      font-weight: 600;
      color: var(--active);
    }
  }

  .gvb_login_form {
    .gvb_login_form_item {
      margin-bottom: 10px;
    }
  }
}
</style>
