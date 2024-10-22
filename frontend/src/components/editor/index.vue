<script setup lang="ts">
import { reactive } from 'vue'
import {
  NCard,
  NForm,
  NFormItem,
  NInput,
  NDynamicTags,
  NSelect
} from 'naive-ui';
import type { TaskView } from './index'
import {statusOptions} from './options'



let model = reactive<TaskView>({
  title: '任务名称',
  description: '任务描述',
  status: '未开始',
  priority: '普通',
  dueDate: new Date().toString(),
  tags: [],
  subtasks: [],
  parentTask: null,
})




</script>

<template>
  <n-card>
    <n-form :model="model">
      <n-form-item path="title" label="任务名称">
        <n-input v-model:value="model.title" />
      </n-form-item>
      <n-form-item path="description" label="任务描述">
        <n-input 
        type="textarea" 
        placeholder="请输入任务描述"
        size="small"
        :autosize="{
          minRows: 3,
          maxRows: 5,
        }"
        v-model:value="model.description"/>
      </n-form-item>
      <n-form-item path="tags" label="标签">
        <n-dynamic-tags v-model:value="model.tags"/>
      </n-form-item>
      <n-form-item>
        <n-select :options="statusOptions" ></n-select>
      </n-form-item>
    </n-form>

    {{ JSON.stringify(model, null, 2)}}
  </n-card>


</template>