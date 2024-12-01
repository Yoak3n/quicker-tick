<script setup lang="ts">
import { reactive,ref } from 'vue'
import {
  NCard,
  NForm,
  NFormItem,
  NInput,
  NDynamicTags,
  NSelect,
  NDatePicker,
  NGrid,
  NGi,
  NButton
} from 'naive-ui';
import type { SelectOption } from 'naive-ui'
import type { Action, TaskView } from '@/types'
import { 
  statusOptions,
  priorityOptions,
  renderOption,
  renderDynamicTag } from './options'

import { GetActions } from '../../../wailsjs/go/app/App';
import { model } from '../../../wailsjs/go/models';
import {showActionEditor} from '.'

const props = defineProps({
  callback: {
    type: Function,
    default: () => {}
  },
  date:{
    type: String,
    required: true
  }
})


let form = reactive<TaskView>({
  title: '任务名称',
  description: '任务描述',
  status: 0,
  priority: 0,
  due_date: props.date,
  actions: [],
  tags: [],
  children: [],
  root: false,
  checked: false,
})

const submitTask = () => {
  props.callback(form)
}
let action_data = reactive<model.Action[]>([])
const getActionsFromDatabase =() => {
  actionsOptions = []
  action_select_loading.value = true
  GetActions().then((res) => {
    action_data = res
    let options_temp:Array<SelectOption> = []
    res.forEach((item) => {
      options_temp.push({
        label: item.name,
        value: item.id
      })
    })
    action_select_loading.value = false
    actionsOptions = options_temp
  })
}
let actionsOptions = reactive<Array<SelectOption>>([])
let action_select_loading = ref<boolean>(false) 

const updateActions = (value:string) => {
  let action_temp:Action[] =[] 
  action_data.findIndex((item) => {
    if (item.id == value) {
      action_temp.push(item)
    }
  })
  form.actions = action_temp
}


</script>

<template>
  <n-card>
    <n-form :model="model">
      <n-form-item path="title" label="任务名称">
        <n-input v-model:value="form.title" />
      </n-form-item>
      <n-form-item path="description" label="任务描述">
        <n-input type="textarea" placeholder="请输入任务描述" size="small" :autosize="{
          minRows: 3,
          maxRows: 5,
        }" v-model:value="form.description" />
      </n-form-item>

      <n-grid :cols="2">
        <n-gi>
          <n-form-item path="tags" label="标签">
            <n-dynamic-tags v-model:value="form.tags" :render-tag="renderDynamicTag" />
          </n-form-item>
        </n-gi>
        <n-gi>
          <n-form-item path="due_date" label="截止日期">
            <n-date-picker :default-calendar-start-time="new Date().getTime()"
              placeholder="请选择截止日期" type="date"
              v-model:formatted-value="form.due_date"
              format="yyyy-M-d">
            </n-date-picker>
          </n-form-item>
        </n-gi>
      </n-grid>
      <n-grid :cols="4">
        <n-gi>
          <n-form-item path="status" label="状态" class="status">
            <n-select :options="statusOptions" :render-label="renderOption" v-model:value="form.status" ></n-select>
          </n-form-item>
        </n-gi>
        <n-gi>
          <n-form-item path="priority" label="优先级" class="priority">
            <n-select v-model:value="form.priority" :options="priorityOptions" :render-label="renderOption" />
          </n-form-item>
        </n-gi>
      </n-grid>
      <n-form-item label="快捷操作" path="action">
        <n-grid :cols="3">
          <n-gi>
            <n-select 
            @update-value="updateActions"
            :loading="action_select_loading"
            @focus="getActionsFromDatabase"
            :options="actionsOptions"/>
          </n-gi>
          <n-gi :span="2">
            <n-button dashed @click="showActionEditor" ghost style="width: 80%;margin:0 auto;" >创建快捷操作</n-button>
          </n-gi>
        </n-grid>
      </n-form-item>
      <n-form-item>
        <n-button type="success" @click="submitTask" style="width: 80%;margin:0 auto;">提交</n-button>
      </n-form-item>
    </n-form>
  </n-card>


</template>

<style scoped>
.status {
  width: 40%;
}

.priority {
  width: 40%;
}
</style>