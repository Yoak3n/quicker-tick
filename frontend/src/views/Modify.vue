<template>
    <n-breadcrumb separator=">">
        <n-breadcrumb-item href="/">
             首页
          </n-breadcrumb-item>
          <n-breadcrumb-item>
            {{ date }}任务清单
          </n-breadcrumb-item>
    </n-breadcrumb>
    <div class="modify-wrapper">
        <editor :callback="getOutput" />
    </div>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { NBreadcrumb, NBreadcrumbItem} from 'naive-ui';
import Editor from '../components/editor/index.vue'
import {AddTask,GetTasks}from '../../wailsjs/go/app/App'

const $route = useRoute();
const date = $route.params.date as string;
const getOutput = (data:string)=> {
    AddTask(data,date).then((res)=>{console.log(res);})
}
</script>

<style scoped lang="less">

</style>