<template>
    <div class="dashboard-wrapper">
        <MainShow :date="$route.params.date as string"></MainShow>
    </div>
</template>

<script lang="ts" setup>
import {onMounted} from "vue";
import {useRoute} from "vue-router";
import MainShow from "@/components/common/MainShow/index.vue"
const $route = useRoute();

onMounted(() => {
    console.log($route.params.date);
})



</script>

<style scoped lang="less">
.dashboard-wrapper{
    height: 100%;
}
</style>