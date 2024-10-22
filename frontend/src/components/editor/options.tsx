import {h,defineComponent} from 'vue'
import type { VNodeChild } from 'vue'
import { NIcon } from 'naive-ui'
import type { SelectOption } from 'naive-ui'
import { 
    MusicalNotes as MusicIcon,
    CheckmarkCircleOutline
 } from '@vicons/ionicons5'

const finishedIcon = defineComponent({
    render(){
        return h(NIcon, {component: MusicIcon},"已完成")
    }
})



export const statusOptions:Array<SelectOption> = [
    {
      label: '未开始',
      value: '未开始',

    },
    {
      label: '进行中',
      value: '进行中'
    },
    {
      label: '已完成',
      value: '已完成',
      render: ()=>{
        return h(NIcon, {component: finishedIcon})
      },
    },
    {
      label: '已取消',
      value: '已取消'
    }
  ]
