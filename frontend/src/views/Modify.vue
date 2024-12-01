<template>
    <n-breadcrumb separator=">">
        <n-breadcrumb-item href="/">
             首页
          </n-breadcrumb-item>
          <n-breadcrumb-item @click="()=>$router.push('/calendar')">
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
import { model } from 'wailsjs/go/models';

const $router = useRouter();
const $route = useRoute();

let date = ref($route.params.date as string);

onMounted(()=>{
    if ($route.params.date == undefined){
        date.value = new Date().toISOString().slice(0,10).replace('/',"-");
    }else{
        date.value = $route.params.date as string;
    }
})

const getOutput = (data:TaskView)=> {
    AddTask(data as model.TaskView).then((res)=>{
        if (res != ""){
            window.$message.error(res,{duration: 2000})
        }

    })
    $router.push('/')
}



</script>

<style scoped lang="less">

</style>