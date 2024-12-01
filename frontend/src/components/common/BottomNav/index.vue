<template>
    <n-flex  class="bottom-nav" justify="space-around">
        <button class="nav" @click="$router.push('/')">
            <n-icon>
                <HomeOutline :size="32" />
            </n-icon>        
        </button>
        <button class="nav" @click="$router.push('/calendar')">
            <n-icon>
                <CalendarOutline :size="32" />
            </n-icon>
        </button>
        <button class="nav" @click="showActionEditor">
            <n-icon>
                <Accessibility :size="32" />
            </n-icon>
        </button>
    </n-flex>
</template>

<script lang="ts" setup>
import {h} from 'vue';
import { useRouter } from 'vue-router';
import { NIcon,NFlex } from 'naive-ui';
import {
  HomeOutline,
  CalendarOutline,
  Accessibility
} from '@vicons/ionicons5'
import AcitonEditor from "@/components/editor/action/index.vue"
const $router = useRouter()

const showActionEditor = () => {
  const m = window.$modal.create({
    title: '创建快捷操作',
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
.bottom-nav {
    width: 100%;
    height: 3rem;

    button.nav {
        border-radius: 10px;
        border: .5px outset #aaa;
        background-color: transparent;
        width: 32%;
        font-size: 20px;
        &:hover{
            color: #18A058;

        }
        &:active{
            color: #18A058;
            border: .5px inset #aaa;
            background-color: #eee;
        }
    }
}

</style>