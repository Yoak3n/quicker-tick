<template>
    <div class="main-show-wrapper">
        <n-data-table :columns="columns" :bordered="false" :data="data" :rowProps/>
        <n-button style="width: 100%;"  dashed class="add-task-button" ghost @click="addTaskButton">添加任务</n-button>
        <n-dropdown
        placement="bottom-start"
        trigger="manual"
        :x="x"
        :y="y"
        :options="options"
        :show="showDropdownRef"
        :on-clickoutside="onClickoutside"
        @select="handleSelect"
        />
    </div>
</template>

<script lang="ts" setup>
import { ref, h, onMounted,nextTick } from 'vue';
import { NDataTable, NTag, NButton,NDropdown,NCheckbox } from 'naive-ui';
import type{ DropdownOption,DataTableColumns  } from 'naive-ui';
import { useRouter } from 'vue-router';
import type { RowData } from './index';
import { GetTasks,RunAction } from '../../../../wailsjs/go/app/App';
import { model } from '../../../../wailsjs/go/models';

const $router = useRouter()
const addTaskButton = () =>{
    $router.push({name:'Modify',params:{date:props.date}})
}
let x = ref(0)
let y = ref(0)
let showDropdownRef = ref(false)
const options: DropdownOption[] = [
  {
    label: '编辑',
    key: 'edit'
  },
  {
    label: () => h('span', { style: { color: 'red' } }, '删除'),
    key: 'delete'
  }
]
const handleSelect = ()=>{
    showDropdownRef.value = false
}
const onClickoutside=() =>{
    showDropdownRef.value = false
}

const props = defineProps(
    {
        date:{
            type: String,
            default: ()=>{
                const date = new Date()
                return `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`
            }
        },
    }
)


let data = ref<model.TaskView[]>()
onMounted(async() => {
    const a =  await GetTasks([props.date])
    data.value = a[props.date]
})



const createColumns = (jp:{StartTask: (row: RowData) => void}) => {
    return [
        {
            title:'',
            width:25,
            key:'checked',
            render(row: RowData) {
                return h(
                    NCheckbox,
                    {
                        onUpdateChecked: (checked) => {
                            row.checked = checked
                        }
                    },
                    {
                        default: () => row.checked
                    }
                )
            }
        },
        {
            title: '任务名',
            key: 'title',
        },
        {
            title: '描述',
            key: 'description',
            width:150,
            ellipsis: {
                tooltip: true
            }
        },
        {
            title: '标签',
            key: 'tags',
            width: 200,
            render(row: RowData) {
                const tags = row.tags.map((tagKey) => {
                    return h(
                        NTag,
                        {
                            style: {
                                marginRight: '6px'
                            },
                            type: 'info',
                            bordered: false
                        },
                        {
                            default: () => tagKey
                        }
                    )
                })
                return tags
            }
        },
        {
            title: '操作',
            key: 'action',
            width: 100,
            render(row: RowData) {
                return h(
                    NButton,
                    {
                        size: 'small',
                        onClick: () => {
                            const actions:Array<Action> = row.actions
                                if (actions.length > 0){
                                    const a = actions[0]
                                    RunAction(a)
                                }
                            jp.StartTask(row)
                        }
                    },
                    { default: () => '开始' }
                )
            }
        },
    ]
}
const rowProps = (row:RowData)=>{
    return {
        onContextmenu: (e: MouseEvent) => {
            // window.$message.info(JSON.stringify(row, null, 2))
            e.preventDefault()
            showDropdownRef.value = false
            nextTick().then(() => {
              showDropdownRef.value = true
              x.value = e.clientX
              y.value = e.clientY
            })
    }
}
}
import { StartTask } from './index';
import { Action } from '@/types';
const columns: DataTableColumns<RowData> = createColumns({StartTask})
</script>

<style scoped lang="less">
.main-show-wrapper{
    height: 100%;
    .add-task-button{
        color: #18A058;
    }
}
</style>