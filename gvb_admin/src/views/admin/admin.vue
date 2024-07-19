
<template>
  <div class="gvb_admin">
    <GVBAside>

    </GVBAside>
    <div class="main">
      <header>
        <div class="left">
          <a-breadcrumb>
            <a-breadcrumb-item>首页</a-breadcrumb-item>
            <a-breadcrumb-item><a href="#">个人中心</a></a-breadcrumb-item>
            <a-breadcrumb-item>用户列表</a-breadcrumb-item>
          </a-breadcrumb>
        </div>
        <div class="right">
          <div class="icon_actions">   
            <i class="fa fa-home"></i>
            <GVBTheme></GVBTheme>

            <i class="fa fa-arrows-alt"></i>   
          </div>
          <GVBUserInfo :is-avatar="true">
          </GVBUserInfo>
        </div>
      </header>
      <div class="tabs"></div>
      <main>
        <div class="gvb_view">
          <div class="gvb_view">
            <router-view/>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>
<script setup>
import GVBTheme from "../../components/gvb_theme.vue"
import GVBAside from "../../components/admin/gvb_aside.vue"
import GVBUserInfo from "../../components/gvb_user_info.vue"
import {ref} from "vue";


</script>

<style lang="scss">
.gvb_admin{
  display: flex;
  width:100%;
  aside{
    width:240px;
    height: 100vh;
    background: var(--gradient_bg);
    //background-color: transparent;

  }
  
  .main{
    width:calc(100% - 240px);
    height: 100vh;
    overflow-y: auto;

    .actions{
      align-items: center;
    }
    .right{

      align-items: center;
      display: flex;
      .icon_actions{
        margin-right:20px;
        i{
          margin-left:10px;
          cursor: pointer;
          color:var(--text)
        }
        i:hover{
          color:var(--active)
        }
      }
      .avatar{
        img{
          width:40px;
          height:40px;
          border-radius:50%;
        }
      }
      .drop_menu{
        margin-left: 10px;
        color: var(--text);
      }
    }
   
    header{
      height: 60px;
      padding: 0 20px;
      display: flex;
      color: var(--text);
      background-color: var(--card_bg);
      justify-content: space-between;
      align-items: center;
    }
    .tabs{
      height:30px;
      border-color:var(--order);
      border: 1px solid var(--order);
      background-color:var(--card_bg)
    }
    main{
      background: var(--bg);
      height: calc(100vh - 90px);
      padding: 20px;
      .gvb_view{

      }
    }
  }
}


</style>