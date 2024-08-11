<template>
    <div class="calendar" style="height: 500px;">
        <n-calendar
            v-model:value="value"
            @update:value="handleUpdateValue"
        >
        </n-calendar>
    </div>

</template>

<script lang="ts" setup>
import {ref,onMounted} from 'vue'
import { NCalendar } from 'naive-ui';
import { useRouter } from 'vue-router';

const $router = useRouter()
let value = ref(Date.now())



const handleUpdateValue = (
    val: number,{year,month,date}:{ year: number, month: number, date: number }
) => {
    console.log(`${year}-${month}-${date}`)
    const d = `${year}-${month}-${date}`
    $router.push({name:'Modify',params:{date:d}})
}


</script>

<style scoped lang="less">

</style>