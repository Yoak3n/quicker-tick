<template>
    <div class="tasks-wrapper">
        <n-grid 
        :collapsed-rows="3" 
        :cols="1" 
        :y-gap="10"
        collapsed>
            <n-gi v-if="tasks && tasks.length > 0" v-for="item in props.tasks" :key="item.id">
                <div class="task-item">
                    {{ item.description }}
                </div>
            </n-gi>
        </n-grid>

    </div>
</template>

<script lang="ts" setup>
import { PropType,toRef,onMounted } from 'vue';
import { NGrid,NGi } from 'naive-ui';
import type { TaskView } from '@/types';
import { ConvertTaskView } from '../../../../wailsjs/go/app/App';
import { model } from 'wailsjs/go/models';
const props = defineProps({
    tasks: {
        type: Object as PropType<Array<TaskView> | undefined> | undefined,
        required: true
    }
})
const tasks = toRef(props, 'tasks')
onMounted(() => {
    ConvertTaskView(tasks.value as Array<model.TaskView>).then((res) => {
        console.log(JSON.stringify(res));
        
    })
})
</script>

<style scoped lang="less">
.task-item{
    width: 100%;
}
</style>