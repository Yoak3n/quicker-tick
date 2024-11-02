<script setup lang="ts">

import Layout from './layout/index.vue'
import NaiveProvider from "@/components/provider/index.vue"

</script>

<template>
  <NaiveProvider>
  <div class="container" style="height: 100%;">
    <Layout/>
  </div>
</NaiveProvider>
</template>

<style lang="less" scoped> 
.content {
  height: auto;
}

.footer {
  position: absolute;
  width: 100%;
  bottom: 0px;
  z-index: 999;
  padding: 10px 0;

  .navigator {
      display: flex;
      justify-content: space-around;
      align-items: center;

      button {
          width: 100px;
          height: 50px;
          border: none;
          border-radius: 5px;
      }
  }
}
</style>