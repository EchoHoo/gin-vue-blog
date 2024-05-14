
<template>
  <div class="gvb_admin">
    <aside>
      <div>

      </div>
    </aside>
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
            <i class="fa fa-moon-o"></i>
            <i class="fa fa-sun-o"></i>
            <i class="fa fa-arrows-alt"></i>   
          </div>
          <div class="avatar">
              <img src="https://pic1.zhimg.com/50/v2-86c2d637bfbc6074a9cced94d8983a75_hd.jpg?source=1940ef5c" alt="">
          </div>
          <div class="drop_menu">
            <a-dropdown>
              <a class="ant-dropdown-link" @click.prevent>
                O.o
                <i class="fa fa-angle-down"></i>
              </a>
              <template #overlay>
                <a-menu>
                  <a-menu-item>
                    <a href="javascript:;">个人中心</a>
                  </a-menu-item>
                  <a-menu-item>
                    <a href="javascript:;">我的消息</a>
                  </a-menu-item>
                  <a-menu-item>
                    <a href="javascript:;">文章列表</a>
                  </a-menu-item>
                  <a-menu-item>
                    <a href="javascript:;">注销退出</a>
                  </a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </div>
        </div>
      </header>
      <div class="tabs"></div>
      <main>
        <div class="gvb_view">
          <router-view/>
        </div>
      </main>
    </div>
  </div>
</template>
<script setup>

</script>

<style lang="scss">
.gvb_admin{
  display: flex;
  width:100%;
  aside{
    width:240px;
    height: 100vh;
    background: #2b3539;
  }
  
  .main{
    width:calc(100% - 240px);
    
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
      background:#ffffff;
      padding: 0 20px;
      display: flex;
      justify-content: space-between
    }
    .tabs{
      height:30px;
      border: 1px solid #f0eeee;
    }
    main{
      background: #f0eeee;
      height: calc(100vh - 90px);
    }
  }
}


</style>