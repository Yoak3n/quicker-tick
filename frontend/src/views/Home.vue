<template>
    <div class="home-wrapper">
        <main-show :date="today"></main-show>
    </div>

</template>

<script lang="ts" setup>
import MainShow from "@/components/common/MainShow/index.vue"


const today = new Date().toISOString().slice(0, 10).replace('/', '-');





</script>

<style scoped lang="less"></style>