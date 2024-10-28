<script setup lang="ts">
import { reactive } from 'vue'
import {
  NCard,
  NForm,
  NFormItem,
  NInput,
  NIcon,
  NDynamicTags,
  NSelect,
  NDatePicker,
  NGrid,
  NGi,
  NUpload,
  NUploadDragger,
  NButton
} from 'naive-ui';
import type { TaskView } from '@/types'
import { RandomColor } from '.';
import { statusOptions, priorityOptions, renderOption,renderDynamicTag } from './options'
import { ArchiveOutline } from '@vicons/ionicons5'


const props = defineProps({
  callback: {
    type: Function,
    default: () => {}
  },
  date:{
    type: String,
    required: true,
  }
})


let model = reactive<TaskView>({
  title: '任务名称',
  description: '任务描述',
  status: 0,
  priority: 0,
  due_date: new Date().toISOString().slice(0, 10).replace('/', '-'),
  actions: [],
  tags: [],
  children: [],
  root: false,
  checked: false,
})

const submitTask = () => {
  props.callback(model)
}




</script>

<template>
  <n-card>
    <n-form :model="model">
      <n-form-item path="title" label="任务名称">
        <n-input v-model:value="model.title" />
      </n-form-item>
      <n-form-item path="description" label="任务描述">
        <n-input type="textarea" placeholder="请输入任务描述" size="small" :autosize="{
          minRows: 3,
          maxRows: 5,
        }" v-model:value="model.description" />
      </n-form-item>

      <n-grid :cols="2">
        <n-gi>
          <n-form-item path="tags" label="标签">
            <n-dynamic-tags v-model:value="model.tags" :render-tag="renderDynamicTag" />
          </n-form-item>
        </n-gi>
        <n-gi>
          <n-form-item path="dueDate" label="截止日期">
            <n-date-picker :default-calendar-start-time="new Date().getTime()"
              :default-time="new Date().toLocaleTimeString()" placeholder="请选择截止日期" type="datetime"
              @update:formatted-value="(value) => { model.due_date = value }" format="yyyy-MM-dd HH:mm">
            </n-date-picker>
          </n-form-item>
        </n-gi>
      </n-grid>
      <n-grid :cols="2">
        <n-gi>
          <n-form-item path="status" label="状态" class="status">
            <n-select :options="statusOptions" :render-label="renderOption" v-model:value="model.status"></n-select>
          </n-form-item>
        </n-gi>
        <n-gi>
          <n-form-item path="priority" label="优先级" class="priority">
            <n-select v-model:value="model.priority" :options="priorityOptions" :render-label="renderOption" />
          </n-form-item>
        </n-gi>
      </n-grid>
      <n-form-item label="快捷操作" path="action">
        <n-grid :cols="3">
          <n-gi>
            <n-select />
          </n-gi>
          <n-gi></n-gi>
          <n-gi>
            <n-upload accept=".doc,.docx">
              <n-upload-dragger>
                <div style="margin-bottom: 12px">
                  <n-icon size="48" :depth="3">
                    <ArchiveOutline />
                  </n-icon>
                </div>
              </n-upload-dragger>
            </n-upload>
          </n-gi>
        </n-grid>
      </n-form-item>
      <n-form-item>
        <n-button type="success" @click="submitTask">提交</n-button>
      </n-form-item>
    </n-form>
    <div>
      {{ JSON.stringify(model, null, 2) }}
    </div>
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