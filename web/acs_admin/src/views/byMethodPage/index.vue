<template >
  <div >
    <methodChart :chartName="currTitle"  />
  </div>
  <div style="margin-top: 30px">
    <allMethodChart style="margin-top: 30px"  chartName="请求方式调用次数(取决于已经存在的记录)"></allMethodChart>
  </div>
</template>

<script setup>
import methodChart from './components/methodChart.vue'
import allMethodChart from './components/allMethodChart.vue'
import { onMounted} from 'vue'
import {usePublicStore} from "@/stores/public";

const currTitle ="请求方式概览"

// store
const storePublic = usePublicStore()
// 面包屑导航信息
const breadcrumb = [
  {
    path: '/',
    name: '首页'
  },
  {
    path: '',
    name: currTitle
  }
]
onMounted(()=>{
  storePublic.SetTitle(currTitle);
  storePublic.SetBreadcrumb(breadcrumb);
})

</script>


<style scoped>

</style>