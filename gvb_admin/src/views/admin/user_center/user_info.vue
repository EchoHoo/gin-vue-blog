<template>
    <div class="gvb_user_info_bg">
      <div class="gvb_user_info_view">
        <div class="user_head">
          个人信息
        </div>
        <div class="body user_info_view">
          <a-form :label-col="{ span: 2 }"
                  :wrapper-col="{ span: 22 }">
            <a-form-item label="昵称">
              <a-input placeholder="昵称" class="user_ipt"></a-input>
            </a-form-item>
            <a-form-item label="我的签名">
              <a-textarea placeholder="我的签名" class="user_ipt" :auto-size="{ minRows: 2, maxRows: 5 }"></a-textarea>
            </a-form-item>
            <a-form-item label="我的博客">
              <a-input placeholder="我的博客" class="user_ipt"></a-input>
            </a-form-item>
            <a-form-item label="邮箱">
              <span>556454545@qq.com</span>
            </a-form-item>
            <a-form-item label="现地址">
              <span>四川成都</span>
            </a-form-item>
            <a-form-item label="当前角色">
              <span>管理员</span>
            </a-form-item>
            <a-form-item label="注册来源">
              <span>QQ</span>
            </a-form-item>
            <a-form-item label="我的积分">
              <span>658</span>
            </a-form-item>
          </a-form>
        </div>
        <div class="user_head">
          每日奖励
        </div>
        <div class="body daily_reward">
          <div>
            <img src="../../../assets/icon/ok.svg" alt="">
            每日登陆 +2
          </div>
          <div>
            <img src="../../../assets/icon/ok.svg" alt="">
            发文 +3
          </div>
          <div>
            <img src="../../../assets/icon/ok.svg" alt="">
            发评论 +1
          </div>
          <div>
            <img src="../../../assets/icon/error.svg" alt="">
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
          <a-button type="primary">绑定邮箱</a-button>
          <a-button type="primary">修改密码</a-button>
          <a-button type="danger">注销退出</a-button>
        </div>
  
  
      </div>
    </div>
  </template>
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
  
      > div {
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