<template>
    <n-form :model="form_model">
        <n-form-item label="Name">
            <n-input placeholder="Action Name" v-model:value="form_model.name"/>
        </n-form-item>
        <n-form-item label="Description">
            <n-input placeholder="Description"  v-model:value="form_model.description"/>
        </n-form-item>
        <n-form-item label="Icon">
            <n-input placeholder="Icon" v-model:value="form_model.icon"></n-input>
        </n-form-item>
        <n-form-item label="Command">
            <n-input placeholder="Command" v-model:value="form_model.command"></n-input>
        </n-form-item>
        <n-form-item>
            <n-button type="primary" class="button" @click="submitAction">Add</n-button>
        </n-form-item>
    </n-form>

</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import { NForm,NFormItem,NInput,NButton } from 'naive-ui';
import type { Action } from '@/types';
import { AddAction } from '../../../../wailsjs/go/app/App';
import { model } from 'wailsjs/go/models';
const defaultAction:Action = {
    name: '',
    description: '',
    icon: '',
    command: ''
}


const form_model = reactive<Action>(defaultAction)
const submitAction= ()=>{
    console.log(form_model);
    AddAction(form_model as model.Action)
    window.$modal.destroyAll()
}
</script>

<style scoped lang="less">
.button{
    width: 50%;
    margin: 0 auto;
}
</style>