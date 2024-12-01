<template>
    <div class="home-wrapper">
        <main-show :date="today"></main-show>
    </div>

</template>

<script lang="ts" setup>
import MainShow from "@/components/common/MainShow/index.vue"


const date = new Date()
const today = `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`





</script>

<style scoped lang="less"></style>