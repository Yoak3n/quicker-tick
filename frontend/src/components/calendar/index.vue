<template>
    <div class="calendar-wrapper">
        <n-card :bordered="false" content-style="padding-top:0;">
            <div class="calendar-header">
                <div class="header-action">
                    <div class="nav">
                        <button text class="pre" 
                        @click="()=>{
                            preMonth()
                            nextTick(()=>{
                                render_view()
                            })
                        }">
                            <n-icon>
                                <ChevronBack/>
                            </n-icon>
                        </button>
                        <button text class="next" 
                        @click="()=>{
                            nextMonth()
                            nextTick(()=>{
                                render_view()
                            })
                        }">
                            <n-icon>
                                <ChevronForward/>
                            </n-icon>
                        </button>
                    </div>
                    <p class="title">{{`${y}年${m}月`}}</p>
                    <button class="today" @click="()=>backToToday()">今天</button>
                </div>
                <div class="day-bar">
                    <n-grid :cols="7" >
                        <n-gi>
                            星期一
                        </n-gi>
                        <n-gi>
                            星期二
                        </n-gi>
                        <n-gi>
                            星期三
                        </n-gi>
                        <n-gi>
                            星期四
                        </n-gi>
                        <n-gi>
                            星期五
                        </n-gi>
                        <n-gi>
                            星期六
                        </n-gi>
                        <n-gi>
                            星期日
                        </n-gi>
                    </n-grid>
                </div>
                
            </div>
            <div class="calendar-body">
                <n-grid :cols="7">
                    <n-gi v-for="i in page_item_count">
                        <CalendarButton  v-if="i < current_month_first_day" 
                        :key="`${y}-${m - 1}-${last_month_count - current_month_first_day + i +1}`"
                        :date="`${y}-${m - 1}-${last_month_count - current_month_first_day + i +1}`"/>
                        <CalendarButton v-else-if="i >= current_month_first_day && i < current_month_count + current_month_first_day"
                            :date="`${y}-${m}-${i - current_month_first_day+1}`" 
                            @click="jump_to_date(i)"
                            :key="`${y}-${m}-${i - current_month_first_day+1}`"
                            :current_month="true"
                            :isToday="computeIsToday(`${y}-${m}-${i - current_month_first_day+1}`)"
                            :tasks="current_month_data.get(`${y}-${m}-${i - current_month_first_day+1}`)" />
                        <CalendarButton v-else
                        :key="`${y}-${m + 1}-${i - current_month_count - current_month_first_day +1}`"
                        :date="`${y}-${m + 1}-${i - current_month_count - current_month_first_day +1}`" />
                    </n-gi>
                </n-grid>
            </div>
        </n-card>
    </div>
</template>

<script lang="ts" setup>
import { ref,onMounted, nextTick,computed } from 'vue';
import { NCard, NGrid, NGi,NIcon } from 'naive-ui';
import { useRouter } from 'vue-router';
import {ChevronBack,ChevronForward} from '@vicons/ionicons5'
import CalendarButton from './CalendarButton/index.vue'
import type { Task} from '@/types';
import {GetTasks}from '../../../wailsjs/go/app/App'
const props = defineProps({
    date: {
        type: String, // String like 2024-01-01 or 2024-01
        default: new Date().toLocaleDateString()
    }
})
const $router = useRouter()
let ds = ref(props.date)
let y = ref<number>(0)
let m = ref<number>(0)
const jump_to_date = (index:number)=>{
    const target = `${y}-${m}-${index - current_month_first_day.value+1 }`
    // $router.push({name:'Modify',params:{date:target}})
    $router.push('/modify/'+target)
}

let current_year = ref<number>(0)
let current_month = ref<number>(0)
let current_date = ref<number>(0)
let last_month_count = ref<number>(0)
let current_month_count = ref<number>(0)
let current_month_first_day = ref<number>(0)
let current_month_data = ref<Map<string,Array<Task>>>(new Map())
onMounted(()=>{
    const ym = ds.value.split('-')
    y.value = Number(ym[0])
    m.value = Number(ym[1])
    render_view()
})
const render_view = async()=> {
    const current = new Date(y.value, m.value - 1)
    current_month.value = current.getMonth() // 获取当前月
    current_year.value = current.getFullYear() // 获取当前年
    current_date.value = current.getDate() // 获取当前日
    last_month_count.value = new Date(current_year.value, current_month.value, 0).getDate() // 获取上个月的天数
    current_month_count.value = new Date(current_year.value, current_month.value + 1, 0).getDate() // 获取当前月的天数
    current_month_first_day.value = new Date(current_year.value, current_month.value, 1).getDay() // 获取当前月的第一天是星期几
    if (current_month_first_day.value ===0)current_month_first_day.value = 7
    
    let dates:Array<string> = []
    for (let i = 0; i < current_month_count.value; i++) {
        dates.push(`${current_year.value}-${current_month.value + 1}-${i + 1}`)
    }
    
    GetTasks(dates).then((res)=>{
        
        dates.forEach((date)=>{
            let views:Array<Task> = []
            const temp =res[date]
            if (temp != undefined) {
                views = temp
                current_month_data.value.set(date,views)
            }
        })
       
    })
}
const backToToday = ()=>{
    const today = new Date()
    y.value = today.getFullYear()
    m.value = today.getMonth() + 1
    nextTick(()=>{
        render_view()
    })
}

const computeIsToday = (date:string)=>{
    const today = new Date()
    const today_str = `${today.getFullYear()}-${today.getMonth() + 1}-${today.getDate()}`
    return today_str === date
}

const page_item_count = computed(() => {
    const remainder = (current_month_count.value + current_month_first_day.value) % 7
    if (remainder !== 0) {
        return (current_month_count.value + current_month_first_day.value) + 7 - remainder
    } else {
        return current_month_count.value + current_month_first_day.value
    }
})

const preMonth = ()=> {
    if(m.value === 1){
        m.value = 12
        y.value -= 1
    }else{
        m.value -= 1
    }

}
const nextMonth = ()=> {
    if(m.value === 12){
        m.value = 1
        y.value += 1
    }else{
        m.value += 1
    }
}
</script>

<style scoped lang="less">
.calendar-wrapper {
    text-align: center;
    .calendar-header{
        margin-bottom: 10px;
        .header-action{
            display: flex;
            justify-content: space-between;
            align-items: center;
            font-size: 16px;
            line-height: 100%;
            .nav{
                .pre:active:not(:disabled),.next:active:not(:disabled){
                    background-color: rgba(0, 0, 0, 0.06);
                }
                .pre{
                    border-top-right-radius: 0;
                    border-bottom-right-radius: 0;
                    border-right: none;
                    box-sizing: border-box;
                    &:hover{
                        border-right:1px solid #e79794;
                    }
                    
                }
                .next{
                    border-top-left-radius: 0;
                    border-bottom-left-radius: 0;
                }
                
            }
            .title{
                font-size: 18px;
                font-weight: 600;
            }
            .today{
                width: 64px;
                font-size: 15px;

            }
        }
        .day-bar {
            text-align: center;
            color:#b3b3b3;
            font-size: 14px;
            cursor:default;
        }
    }
    .calendar-body {
        margin-top: 10px;
    }
}
button{
    width: 42px;
    height: 28px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    border-radius: 14px;
    border: 1px solid currentColor;
    color: #c9c9c9;
    background-color: #fff;
    cursor: pointer;
    transition: all .3s;
    &:disabled{
        background-color: rgba(0, 0, 0, 0.08);
        cursor: not-allowed;
    }
    &:hover{
        color: #e79794;
    }
}
.pre:hover + .next{
    border-left:none;
}
</style>