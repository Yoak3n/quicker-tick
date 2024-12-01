<template>
    <div class="calendar-button-wrapper">
        <button class="calendar-button" 
        :disabled="!props.current_month">
            <n-flex vertical wrap  align="space-around" style="height: 90%;">
                <div class="calendar-button-date"
                :style="props.isToday?'color: red': ''"
                >
                    {{ index}}
                </div>
                <div class="calendar-button-content">
                    <CalendarContent v-if="props.tasks" :tasks="props.tasks"/>
                </div>
            </n-flex>
        </button>
    </div>
</template>

<script lang="ts" setup>
import { onMounted,ref,PropType } from 'vue';
import {NFlex} from 'naive-ui';
import CalendarContent from '../CalendarContent/index.vue'
import type{TaskView} from '@/types'
const props = defineProps({
    date: {
        type: String,
        default: ''
    },
    current_month:{
        type: Boolean,
        default: false
    },
    isToday: {
        type: Boolean,
        default: false
    },
    tasks: {
        type: Object as PropType<Array<TaskView> | undefined> | undefined,
    }
})
const date =new Date(props.date)

let index = ref('')
onMounted(()=>{
    index.value = date.getDate().toString()
    if (props.tasks != undefined){
        console.log(props.tasks)
    }
    
})

</script>

<style scoped lang="less">
.calendar-button-wrapper{
    .calendar-button{
        border: none;
        background-color: #fff;
        width: 100%;
        height: 8rem;
        border-radius: 5px;
        &:hover{
            background-color: #f0f0f0;
        }
        &:focus{
            outline: none;
        }
        &:active{
            background-color: #117af2;
            color: #fff;
        }

        &:focus + &:active{
            background-color: #117af2;
            color: #fff;

        }
        &:disabled{
            background-color: #fff;
            color: #ccc;
        }
        .calendar-button-date{
            font-size: 1.5rem;
            font-weight: bold;
            text-align: center;
            margin-bottom: 0.5rem;
            line-height: 1.5rem;;
        }
        .calendar-button-content{
            height: 80%;
        }
        //&:not(:disabled)::after{
        //     content: '';
        //    display: block;
        //    width: 50%;
        //   height: 1px;
        //    background-color: #ccc;
        //    margin:0 auto;
        //}
    }
    .active{
        background-color: #117af2;
        color: #fff;
    }


}

</style>