<script setup lang="ts">
import AppBar from '@/components/common/AppBar/index.vue'
import Layout from './layout/index.vue'
import NaiveProvider from "@/components/provider/index.vue"

</script>

<template>
    <NaiveProvider>
      <Layout />
    </NaiveProvider>
</template>

<style lang="less" scoped>
</style>