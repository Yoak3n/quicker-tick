<template>
    <div class="layout" ref="containerRef">
        <div class="header" >
            <app-bar/>
        </div>
        <div class="content">
            <router-view></router-view>
        </div>
        <div class="footer">
            <n-float-button :bottom="5" :right="5"  menu-trigger="hover">
                <n-icon>
                    <MenuOutline/>
                </n-icon>
                <template #menu>
                    <n-float-button @click="$router.push('/calendar')" shape="circle" type="primary">
                        <n-icon>
                            <CalendarOutline/>
                        </n-icon>
                    </n-float-button>
                    <n-float-button @click="()=>$router.push('/modify')" type="primary">

                    </n-float-button>
                    <n-float-button type="primary">
                        <n-icon>
                            <PencilOutline/>
                        </n-icon>
                    </n-float-button>
                </template>

            </n-float-button>

        </div>
    </div>

</template>

<script lang="ts" setup>
import {ref} from "vue";
import { NFloatButton,NIcon } from "naive-ui";
import AppBar from '@/components/common/AppBar/index.vue'
import { useRouter } from 'vue-router';
import {
    MenuOutline,
    CalendarOutline,
    PencilOutline
} from "@vicons/ionicons5"
const $router = useRouter();
const containerRef = ref<HTMLElement | undefined>(undefined)
</script>

<style scoped lang="less">
.layout{
    height: 100vh;
    .header{
        position: sticky;
        top: 0;
        z-index: 999
    }
    .content{
        height: auto;
    }
    .footer{
        position: absolute;
        width: 100%;
        bottom: 0px;
        z-index: 999;
        padding: 10px 0;
        .navigator{
            display: flex;
            justify-content: space-around;
            align-items: center;
            button{
                width: 100px;
                height: 50px;
                border: none;
                border-radius: 5px;
            }
        }
    }
}

</style>