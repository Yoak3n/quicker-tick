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
import { PropType,toRef } from 'vue';
import { NGrid,NGi } from 'naive-ui';
import type { TaskView } from '@/types';
const props = defineProps({
    tasks: {
        type: Object as PropType<Array<TaskView> | undefined> | undefined,
        required: true
    }
})
const tasks = toRef(props, 'tasks')
</script>

<style scoped lang="less">
.task-item{
    width: 100%;
}
</style>