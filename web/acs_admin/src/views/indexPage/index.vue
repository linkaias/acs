<template >
  <div>
    <AllCountChart chartName="总调用次数"  />
  </div>
  <div style="margin-top: 30px">
    <InfoTable />
  </div>
</template>

<script setup>
  import AllCountChart from './components/AllCountChart.vue'
  import InfoTable from './components/infoTable.vue'
  import { usePublicStore} from '@/stores/public'
  import {onMounted, reactive } from 'vue'
  const storePublic = reactive(usePublicStore())

  const breadcrumb = [
      {
        path: '/',
        name: '首页'
      },
      {
        path: '',
        name: '总调用次数概览'
      }
  ]
  onMounted(()=>{
    storePublic.SetTitle('总调用次数概览');
    storePublic.SetBreadcrumb(breadcrumb);
  })

</script>


<style scoped>

</style>