<template>
    <n-breadcrumb separator=">">
        <n-breadcrumb-item href="/">
             首页
          </n-breadcrumb-item>
          <n-breadcrumb-item @click="()=>$router.push('/dashboard')">
            日历视图
          </n-breadcrumb-item>
          <n-breadcrumb-item>
            {{ date }}任务清单
          </n-breadcrumb-item>
    </n-breadcrumb>
    <div class="modify-wrapper">
        <editor :callback="getOutput" :date="date"/>
    </div>
</template>

<script lang="ts" setup>
import { onMounted,ref,onBeforeMount } from 'vue';
import { useRoute,useRouter } from 'vue-router';
import { NBreadcrumb, NBreadcrumbItem} from 'naive-ui';
import Editor from '../components/editor/index.vue'
import {AddTask}from '../../wailsjs/go/app/App'
import { TaskView } from '@/types';

const $router = useRouter();
const $route = useRoute();
const date = $route.params.date as string;
const getOutput = (data:string)=> {
    AddTask(data,date).then((res)=>{
        if (res != ""){
            console.log(res);
            
        }
    })
}



</script>

<style scoped lang="less">

</style>