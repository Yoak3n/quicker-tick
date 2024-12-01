<template>
    <n-form :model="form_model">
        <n-form-item label="操作名称">
            <n-input placeholder="Action Name" v-model:value="form_model.name"/>
        </n-form-item>
        <n-form-item label="描述">
            <n-input placeholder="Description"  v-model:value="form_model.description"/>
        </n-form-item>
        <n-form-item label="图标">
            <n-input placeholder="Icon" v-model:value="form_model.icon"></n-input>
        </n-form-item>
        <n-form-item label="命令">
            <n-select 
            style="width: 40%"
            :options="actionTypeOptions" 
            @update:value="changeArgTip"
            :default-value="defaultAction.type"/>
            <n-input 
            :placeholder="argTipPlaceholder" 
            style="width: 60%"
            v-model:value="form_model.command"></n-input>
        </n-form-item>
        <n-form-item>
            <n-button type="primary" class="button" @click="submitAction">Add</n-button>
        </n-form-item>
    </n-form>

</template>

<script lang="ts" setup>
import { reactive,ref } from 'vue';
import { NForm,NFormItem,NInput,NButton,NSelect } from 'naive-ui';
import type { Action } from '@/types';
import { actionTypeOptions } from '.';
import { AddAction } from '../../../../wailsjs/go/app/App';
import { model } from 'wailsjs/go/models';
const defaultAction:Action = {
    id:'',
    name: '',
    description: '',
    icon: '',
    command: '',
    type:'cmd'
}
let argTipPlaceholder = ref('Command')

const form_model = reactive<Action>(defaultAction)
const submitAction= ()=>{
    console.log(form_model);
    AddAction(form_model as model.Action)
    window.$modal.destroyAll()
}

const changeArgTip = (value:string)=>{
    if (value === 'browser'){
        argTipPlaceholder.value = 'Url'
    }else{
        argTipPlaceholder.value = 'Command'
    }
}

</script>

<style scoped lang="less">
.button{
    width: 50%;
    margin: 0 auto;
}
</style>