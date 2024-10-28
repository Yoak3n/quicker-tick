<template>
    <div class="dashboard-wrapper">

    </div>
</template>

<script lang="ts" setup>
import {onMounted} from "vue";
import {useRoute} from "vue-router";

const $route = useRoute();

onMounted(() => {
    console.log($route);
})



</script>

<style scoped lang="less">

</style>