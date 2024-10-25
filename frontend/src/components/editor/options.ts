import {h} from 'vue'
import type { VNodeChild } from 'vue'
import { NTag } from 'naive-ui'
import type { SelectOption } from 'naive-ui'

export const priorityOptions:Array<SelectOption> = [
    {
      label: '低',
      value: 0,
      type: 'success'
    },
    {
      label: '中',
      value: 1,
      type: 'info'
    },
    {
      label: '高',
      value: 2,
      type: 'warning'
    },
    {
      label: '紧急',
      value: 3,
      type: 'error'
    }
]


export const statusOptions:Array<SelectOption> = [
    {
      label: '未开始',
      value: 0,
      type: 'warning'

    },
    {
      label: '进行中',
      value: 1,
      type: 'info'
    },
    {
      label: '已完成',
      value: 2,
      type: 'success'
    },
    {
      label: '已取消',
      value: 3,
      type: 'error'
    }
  ]

export function renderOption (option: SelectOption): VNodeChild {
    return [
      h(NTag, { 
        type: option.type as 'success' | 'warning' | 'error'|'info'}, 
        option.label?.toString()),
    ]

}