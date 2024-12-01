<template>
  <div class="layout">
    <div class="header">
      <app-bar />
    </div>
    <div class="body">
      <router-view v-slot="{ Component }">
        <component :is="Component" />
      </router-view>
    </div>
    <div class="footer">
      <bottom-nav/>
    </div>
  </div>
<!-- 


  <n-float-button :bottom="15" :right="5" menu-trigger="hover">
    <n-icon>
      <MenuOutline />
    </n-icon>
    <template #menu>
      <n-float-button @click="$router.push('/calendar')" shape="circle" type="primary">
        <n-icon>
          <CalendarOutline />
        </n-icon>
      </n-float-button>
      <n-float-button type="primary" @click="() => {
        showActionEditor()
      }">
        <n-icon>
          <Accessibility />
        </n-icon>
      </n-float-button>
      <n-float-button type="primary" @click="() => {
        if ($route.params.date != '' && $route.params.date != undefined) {
          $router.push({
            name: 'Modify',
            params: {
              date: $route.params.date
            }
          });
        } else {
          
            const date = new Date()
            const d = `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`
        
          $router.push(
            {
              name: 'Modify',
              params: {
                date: d
              }
            });
        }
      }">
        <n-icon>
          <PencilOutline />
        </n-icon>
      </n-float-button>
    </template>

  </n-float-button> -->
</template>

<script lang="ts" setup>
import { ref, h } from "vue";
import { NFloatButton, NIcon, useModal } from "naive-ui";
import AppBar from '@/components/common/AppBar/index.vue'
import BottomNav from '@/components/common/BottomNav/index.vue'
import { useRouter, useRoute } from 'vue-router';
import {
  MenuOutline,
  CalendarOutline,
  PencilOutline,
  Accessibility
} from "@vicons/ionicons5"
import AcitonEditor from "@/components/editor/action/index.vue"


const $router = useRouter();
const $route = useRoute();
window.$modal = useModal()

const showActionEditor = () => {
  const m = window.$modal.create({
    title: 'Action Editor',
    preset: 'card',
    style: {
      width: '400px'
    },
    content: () =>
      h(
        AcitonEditor,
        { type: 'primary' },
        () => '关闭'
      )
  })

}
</script>

<style scoped lang="less">
.layout {
  position: relative;
  height: 100vh;
  .header {
    height: auto;
    position: sticky;
    top: 0;
    z-index: 999
  }
  .body {
    height: 88%;
    overflow-y: auto;
    scrollbar-width: none;
  }
  .footer{
    height: auto;
    position: sticky;
    bottom: 0;
    width: 100%;
    z-index: 999;
  }


}
</style>