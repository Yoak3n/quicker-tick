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
  NUploadDragger
} from 'naive-ui';
import type { TaskView } from './index'
import { statusOptions, renderOption } from './options'
import { ArchiveOutline } from '@vicons/ionicons5'


let model = reactive<TaskView>({
  title: '任务名称',
  description: '任务描述',
  status: '未开始',
  priority: '普通',
  dueDate: new Date().toISOString().slice(0, 10).replace('/', '-'),
  action: [],
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
        <n-input type="textarea" placeholder="请输入任务描述" size="small" :autosize="{
          minRows: 3,
          maxRows: 5,
        }" v-model:value="model.description" />
      </n-form-item>

      <n-grid :cols="2">
        <n-gi>
          <n-form-item path="tags" label="标签">
            <n-dynamic-tags v-model:value="model.tags" />
          </n-form-item>
        </n-gi>
        <n-gi>
          <n-form-item path="dueDate" label="截止日期">
            <n-date-picker :default-calendar-start-time="new Date().getTime()"
              :default-time="new Date().toLocaleTimeString()" placeholder="请选择截止日期" type="datetime"
              @update:formatted-value="(value) => { model.dueDate = value }">
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
            <n-input v-model:value="model.priority" placeholder="优先级" />
          </n-form-item>
        </n-gi>
      </n-grid>
      <n-form-item label="快捷操作" path="action">
        <n-upload>
          <n-upload-dragger>
            <div style="margin-bottom: 12px">
              <n-icon size="48" :depth="3">
                <ArchiveOutline />
              </n-icon>
            </div>
          </n-upload-dragger>
        </n-upload>
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