<template >
  <div >
    <groupChart v-if="groupListLoading" :groupList="groupList" :chartName="currTitle"  />
  </div>
  <div style="margin-top: 30px">
    <allGroupChart style="margin-top: 30px"  chartName="按分组调用次数"></allGroupChart>
  </div>
</template>

<script setup>
  import groupChart from './components/groupChart.vue'
  import allGroupChart from './components/allGroupChart.vue'

  import { reqGetAllGroup} from '@/api/group/index'
  import {ref, onMounted} from 'vue'
  import {usePublicStore} from "@/stores/public";

  const currTitle = "分组调用次数"

  // 请求分组列表
  const groupList = ref([])
  const groupListLoading = ref(false)
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
    LoadGroupList();
  })

  const LoadGroupList = ()=>{
    reqGetAllGroup().then(res=>{
      groupList.value = res.data
      groupListLoading.value= true
    })
  }

</script>


<style scoped>

</style>