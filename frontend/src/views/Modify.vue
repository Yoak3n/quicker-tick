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
        <editor :callback="getOutput" :data="data" />
    </div>
</template>

<script lang="ts" setup>
import { onMounted,ref } from 'vue';
import { useRoute } from 'vue-router';
import { NBreadcrumb, NBreadcrumbItem} from 'naive-ui';
import Editor from '../components/editor/index.vue'
import {AddTask,GetTasks}from '../../wailsjs/go/app/App'
import { TaskView } from '@/types';

const $route = useRoute();
const date = $route.params.date as string;
const getOutput = (data:string)=> {
    AddTask(data,date).then((res)=>{
        if (res != ""){
            alert(res)
        }
    })
}
let data = ref<TaskView[]>()
onMounted(()=>{
    GetTasks([date]).then((v)=>{
        data.value = v[date]
    });
})
</script>

<style scoped lang="less">

</style>