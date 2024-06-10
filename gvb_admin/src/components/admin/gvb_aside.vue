<template>
  <aside class="gvb_aside">
    <div class="gvb_aside_header flex">
      <div class="gvb_aside_logo flex">
        <img src="https://pic1.zhimg.com/50/v2-86c2d637bfbc6074a9cced94d8983a75_hd.jpg?source=1940ef5c" alt="">
      </div>
      <div class="gvb_aside_title">后台管理系统</div>
    </div>
    <div class="gvb_aside_body flex">
      <GVBAsideMenu></GVBAsideMenu>
      

    </div>
  </aside>
</template>

<script setup>
  import GVBAsideMenu from "./gvb_aside_menu.vue"
</script>

<style lang="scss">
.gvb_aside {
  width: 240px;
  height: 100vh;
  background-color: var(--slide);

  .gvb_aside_header {
    width: 100%;
    height: 180px;
    flex-direction: column;
  }

  .gvb_aside_logo {
    padding: 10px;

    img {
      width: 80px;
      height: 80px;
      border-radius: 50%;
    }
  }

  .gvb_aside_title {
    margin-top: 20px;
    font-size: 20px;
    color: #aaa7a7;
  }

}
</style>