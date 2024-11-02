<template>
    <div class="header">
        <app-bar />
      </div>
        <router-view v-slot="{ Component }">
            <component :is="Component" />
        </router-view>
        <div class="footer">
            <n-float-button :bottom="5" :right="5" menu-trigger="hover">
              <n-icon>
                <MenuOutline />
              </n-icon>
              <template #menu>
                <n-float-button @click="$router.push('/calendar')" shape="circle" type="primary">
                  <n-icon>
                    <CalendarOutline />
                  </n-icon>
                </n-float-button>
                <n-float-button type="primary"  @click="() =>{
                    showActionEditor()
                }">
                  <n-icon>
                    <Accessibility />
                  </n-icon>
                </n-float-button>
                <n-float-button type="primary" @click="() => {
                  if ($route.params.date != '' || $route.params.date != undefined) {
                    $router.push({ name: 'Modify', params: { date: $route.params.date } });
                  }
                }">
                  <n-icon>
                    <PencilOutline />
                  </n-icon>
                </n-float-button>
              </template>
      
            </n-float-button>
      
          </div>
</template>

<script lang="ts" setup>
import { ref,h } from "vue";
import { NFloatButton, NIcon, NModal, NButton,useModal } from "naive-ui";
import AppBar from '@/components/common/AppBar/index.vue'
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
const containerRef = ref<HTMLElement | undefined>(undefined)
const modal = useModal()

const showActionEditor = () => {

    const m = modal.create({
        title: 'Card 预设',
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
    height: 100vh;

    .header {
        position: sticky;
        top: 0;
        z-index: 999
    }

   
}
</style>