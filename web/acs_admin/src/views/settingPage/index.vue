<template>
  <el-card shadow="never">
    <template #header>
      <div class="card-header">
        <span>系统设置</span>
      </div>
    </template>
    <div class="config_item">
      <FormSetting/>
    </div>
  </el-card>
</template>
<script setup>
  import  FormSetting from './components/formSetting.vue'
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
      name: '系统设置'
    }
  ]
  onMounted(()=>{
    storePublic.SetTitle('系统设置');
    storePublic.SetBreadcrumb(breadcrumb);
  })
</script>

<style scoped>

</style>