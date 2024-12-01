<template>
  <n-flex justify="space-between" style="padding:0.5rem 0.5rem 0 0.5rem;background-color: #f5f5f5;--wails-draggable:drag;height:32px">
    <div class="header">
        <b>滴答清单</b>
  </div>
    <div class="buttons">
      <n-button text @click="()=> $router.go(-1)">
        <template #icon>
            <n-icon ><ArrowBackOutline/></n-icon>
        </template>
      </n-button>
      <n-button text @click="pushToHome">
        <template #icon>
            <n-icon ><HomeOutline/></n-icon>
        </template>
      </n-button>
      <n-button text @click="closeApp">
        <template #icon>
            <n-icon ><Close/></n-icon>
        </template>
      </n-button>
    </div>
  </n-flex>

</template>

<script lang="ts" setup>
import {h,Component} from 'vue';
import {NFlex,NIcon,NButton} from "naive-ui";
import {
  HomeOutline,
  ArrowBackOutline,
  Close
} from '@vicons/ionicons5'

function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}
import { Quit } from '../../../../wailsjs/runtime/runtime'
import { useRouter } from 'vue-router';
const $router = useRouter()

const pushToHome  = (e: MouseEvent) => {
  e.preventDefault()
  $router.push('/')
}
const closeApp=(e: MouseEvent)=>{
  e.preventDefault()
  Quit()
}

</script>

<style scoped lang="less">
.header{
  cursor: default;
}
.buttons{
  button {
    margin: 0 0.5rem;
    &:hover{
      color: #ff3366;
    }
  }
}
</style>