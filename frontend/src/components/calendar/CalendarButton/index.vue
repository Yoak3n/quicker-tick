<template>
    <div class="calendar-button-wrapper">
        <button class="calendar-button" 
        @click="openDate"
        :disabled="!props.current_month">
            <n-flex vertical wrap :size="0">
                <div class="calendar-button-date"
                :style="props.isToday?'color: red': ''"
                >
                    {{ index}}
                </div>
                <div class="calendar-button-content">
                    <slot></slot>
                </div>
            </n-flex>
        </button>
    </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router';
import {NFlex} from 'naive-ui';
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
    }
})
const date =new Date(props.date)
const index = date.getDate().toString()
const $router = useRouter()

const openDate = () => {
    $router.push({name:'Modify',params:{date:props.date}})
}
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
    }
    .active{
        background-color: #117af2;
        color: #fff;
    }
}

</style>