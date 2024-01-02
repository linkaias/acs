<template>
  <el-row :gutter="20">
    <el-col :span="12">
      <el-card shadow="never">
        <template #header>
          <div class="card-header">
            <span>请求分组信息</span>
          </div>
        </template>
        <el-table v-loading="groupLoading" :data="groupTableData" style="width: 100%; height: 300px">
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="name" label="分组名称" width="180" />
          <el-table-column prop="desc" label="分组描述" />
          <el-table-column prop="updated_at" label="更新时间" width="180"  />
          <el-table-column prop="created_at" label="创建时间" width="180"  />
        </el-table>
      </el-card>
    </el-col>
    <el-col :span="12">
      <el-card shadow="never">
        <template #header>
          <div class="card-header">
            <span>系统参数</span>
          </div>
        </template>
        <el-table v-loading="systemLoading" :data="systemInfoData" style="width: 100%;height: 300px">
          <el-table-column type="index" label="序号" width="100" />
          <el-table-column prop="name" label="参数名称" width="280" />
          <el-table-column prop="value" label="参数值"  />
        </el-table>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
  import {onMounted, ref} from "vue"
  import {reqGetAllGroup} from '@/api/group'
  import {reqGetSystemInfo} from '@/api/system'
  import dayjs from "dayjs";

  const groupLoading = ref(false)
  const groupTableData = ref([])

  const systemLoading = ref(false)
  const systemInfoData = ref([])

  onMounted(()=>{
    loadGroup();
    loadSystemInfo();
  })
  const loadGroup = async ()=>{
    groupLoading.value = true
    let info = await reqGetAllGroup()
    groupLoading.value = false
    // 遍历info.data,并且修改数据字段格式
    let data = info.data;
    for (let i = 0; i < data.length; i++) {
      data[i].created_at = dayjs(data[i].created_at).format('YYYY-MM-DD HH:mm:ss')
      data[i].updated_at = dayjs(data[i].updated_at).format('YYYY-MM-DD HH:mm:ss')
    }
    groupTableData.value= data
  }

  const loadSystemInfo = async ()=>{
    systemLoading.value = true
    let info = await reqGetSystemInfo()
    systemLoading.value = false
    let data = info.data;
    // data是一个对象， 遍历data，将对象转换为数组
    let dataArray = []
    Object.keys(data).forEach((key)=>{
      //  遍历data[key]对象
      let obj = data[key];
      Object.keys(obj).forEach((k)=>{
        dataArray.push({
          name: k,
          value: obj[k]
        })
      })
    })
    systemInfoData.value = dataArray
  }

</script>

<style scoped>

</style>