import {h} from 'vue'
import type { VNodeChild } from 'vue'
import { NTag } from 'naive-ui'
import type { SelectOption } from 'naive-ui'



export const statusOptions:Array<SelectOption> = [
    {
      label: '未开始',
      value: '未开始',
      type: 'warning'

    },
    {
      label: '进行中',
      value: '进行中',
      type: 'info'
    },
    {
      label: '已完成',
      value: '已完成',
      type: 'success'
    },
    {
      label: '已取消',
      value: '已取消',
      type: 'error'
    }
  ]

export function renderOption (option: SelectOption): VNodeChild {
    return [
      h(NTag, { 
        type: option.type as 'success' | 'warning' | 'error'|'info',}, 
        option.label?.toString()),
    ]

}