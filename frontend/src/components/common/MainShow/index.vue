<template>
    <div class="main-show-wrapper">
        <n-data-table :columns="columns" :bordered="false" :data="data" />
    </div>
</template>

<script lang="ts" setup>
import { ref, h, onMounted } from 'vue';
import { NDataTable, NTag, NButton } from 'naive-ui';
import type { DataTableColumns } from 'naive-ui';
import type { RowData } from './index';
import { GetTasks } from '../../../../wailsjs/go/app/App';
import { model } from '../../../../wailsjs/go/models';


const props = defineProps(
    {
        date:{
            type: String,
            default: new Date().toISOString().slice(0, 10).replace('/', '-')
        }
    }
)


let data = ref<model.TaskView[]>()
onMounted(async() => {
    const a =  await GetTasks([props.date])
    console.log(a);
    data.value = a[props.date]
})



const createColumns = (jp:{StartTask: (row: RowData) => void}) => {
    return [
        {
            title: '任务名',
            key: 'title',
        },
        {
            title: '描述',
            key: 'description',
        },
        {
            title: '标签',
            key: 'tags',
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
            render(row: RowData) {
                return h(
                    NButton,
                    {
                        size: 'small',
                        onClick: () => jp.StartTask(row)
                    },
                    { default: () => '开始' }
                )
            }
        },
    ]
}


import { StartTask } from './index';
const columns: DataTableColumns<RowData> = createColumns({StartTask})
</script>

<style scoped lang="less">

</style>