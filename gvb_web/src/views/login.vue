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
import { reactive, ref } from 'vue';
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

const isLoggingIn = ref(false);

async function emailLogin() {
  // 防止重复提交
  if (isLoggingIn.value) return;
  isLoggingIn.value = true;

  // 检查用户名和密码是否为空
  if (data.user_name === "" || data.password === "") {
    message.error("用户名或密码不能为空");
    isLoggingIn.value = false;
    return;
  }

  try {
    // 发送登录请求
    let res = await emailLoginApi(data);
    if (res.code) {
      message.error(res.msg); // 如果登录失败，显示错误信息
    } else {
      message.success(res.msg); // 显示成功信息
      let userInfo = parseToken(res.data); // 解析token获取用户信息
      userInfo.token = res.data;
      store.setUserInfo(userInfo); // 更新全局状态

      // 根据路由查询参数决定重定向目标
      const redirectUrl = route.query.redirect_url;
      if (redirectUrl) {
        router.push({ path: redirectUrl });
      } else {
        router.push('/admin/home');
      }
    }
  } catch (error) {
    message.error("登录请求失败，请检查网络或联系管理员");
  } finally {
    isLoggingIn.value = false; // 重置登录状态标记
  }
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
