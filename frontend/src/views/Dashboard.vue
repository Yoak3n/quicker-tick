<template>
    <div class="dashboard-wrapper">
        <Calendar 
        :date="value"
        />
    </div>
</template>

<script lang="ts" setup>
import Calendar from '@/components/calendar/index.vue'

const date = new Date()
const current_month =  date.getMonth()+ 1
const current_year = date.getFullYear()

const value = `${current_year}-${current_month}`

</script>

<style scoped lang="less">

</style>