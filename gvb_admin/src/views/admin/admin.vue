
<template>
  <div class="gvb_admin">
    <aside>
      <div>
        <i class = "iconfont icon-shangchuantupian"></i>
      </div>
    </aside>
    <div class="main">
      <header>
        <i class="iconfont icon-shangchuantupian"></i>
        <i class="fa fa-battery-three-quarters"></i>
      </header>
      <div class="tabs"></div>
      <main></main>
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

    header{
      height: 60px;
      background:white;
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