<template>
  <el-card shadow="never">
    <template #header>
      <div class="card-header">
        <span>请求记录列表</span>
      </div>
    </template>
    <div  class="my_table_header">
      <el-form :model="FormSearch" inline >
        <el-form-item >
          <el-select filterable v-model="FormSearch.group_id"
                     placeholder="分组筛选">
            <el-option
                v-for="(item,index) in groupSearchData"
                :key="index"
                :label="item.name"
                :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item >
          <el-select filterable v-model="FormSearch.method"
                     placeholder="Method筛选">
            <el-option
                v-for="(item,index) in selectMethodData"
                :key="index"
                :label="item.name"
                :value="item.val"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-input v-model="FormSearch.url" placeholder="筛选url"></el-input>
        </el-form-item>
        <el-form-item>
          <el-input v-model="FormSearch.ip" placeholder="筛选IP"></el-input>
        </el-form-item>
        <el-form-item >
          <el-date-picker
              v-model="ChartTimeRange"
              type="datetimerange"
              @change="onChangeAllCount"
              :shortcuts="shortcuts"
              value-format="YYYY-MM-DD HH:mm"
              range-separator="-"
              start-placeholder="开始时间"
              end-placeholder="结束时间"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :icon="Search" @click="search" plain>搜索</el-button>
        </el-form-item>
      </el-form>
      <el-form inline>
        <el-form-item>
          <el-button :icon="Refresh" circle @click="loadData" />
        </el-form-item>
      </el-form>
    </div>
    <el-table size="large" border v-loading="Loading" :data="TableData" style="width: 100%;margin-bottom: 50px">
      <el-table-column prop="id" label="ID" width="100" />
      <el-table-column prop="name" label="请求名称" width="350" />
      <el-table-column prop="url" label="请求URL"  />
      <el-table-column prop="group_id" label="所属分组" width="200" />
      <el-table-column prop="method" label="Method" width="150" />
      <el-table-column prop="ip" label="IP" width="280" />
      <el-table-column prop="add_time" label="创建时间" width="220"  />
    </el-table>
    <div class="page_lk" v-loading="Loading">
      <el-pagination
          :page-size="Size"
          layout="prev, pager, next, total"
          :total="Total"
          v-model:current-page="CurrPage"
      />
    </div>
  </el-card>
</template>

<script setup>
  import {reqGetRecordList} from "@/api/record";
  import { reqGetAllGroup } from  '@/api/group'
  import dayjs from "dayjs";
  import {onMounted, ref, reactive, watch} from "vue";
  import { Refresh,Search } from '@element-plus/icons-vue'
  import { usePublicStore} from '@/stores/public'
  import {TimeRangeShortcuts} from "@/utils/time_utils";


  const storePublic = reactive(usePublicStore())
  const Loading = ref(false)
  const TableData = ref([])
  const breadcrumb = [
    {
      path: '/',
      name: '首页'
    },
    {
      path: '',
      name: '请求记录列表'
    }
  ]
  const FormSearch= ref({
    group_id:'',
    url:'',
    method:'',
    ip:'',
    start_time:'',
    end_time:''
  })
  const selectMethodData = [
    {
      "val": "",
      "name": "全部",
    },
    {
      "val": "GET",
      "name": "GET",
    },
    {
      "val": "POST",
      "name": "POST",
    },
    {
      "val": "PUT",
      "name": "PUT",
    },
    {
      "val": "DELETE",
      "name": "DELETE",
    },
    {
      "val": "PATCH",
      "name": "PATCH",
    },
    {
      "val": "HEAD",
      "name": "HEAD",
    },
    {
      "val": "OPTIONS",
      "name": "OPTIONS",
    },

  ]

  const groupSearchData = ref([{
    "id": "",
    "name": "全部",
  }])
  // 时间筛选快捷选项
  const shortcuts = TimeRangeShortcuts();
  // 时间筛选值
  const ChartTimeRange = ref([])

  const CurrPage = ref(1)
  const Total = ref(0)
  const Size = ref(10)

  onMounted(async ()=>{
    storePublic.SetTitle('请求记录列表');
    storePublic.SetBreadcrumb(breadcrumb);
    await loadGroup();
    await loadData();
  })
  const loadData = async ()=>{
    Loading.value = true;
    let params={
      page: CurrPage.value,
      size: Size.value,
      start_time:FormSearch.value.start_time,
      end_time:FormSearch.value.end_time,
      group_id:FormSearch.value.group_id,
      url:FormSearch.value.url,
      method:FormSearch.value.method,
      ip:FormSearch.value.ip
    }
    let info = await reqGetRecordList(params)
    Loading.value = false
    // 遍历info.data,并且修改数据字段格式
    let data = info.data;
    for (let i = 0; i < data.length; i++) {
      data[i].add_time = dayjs(data[i].add_time).format('YYYY-MM-DD HH:mm:ss')
      // 匹配分组名称
      for (let j = 0; j < groupSearchData.value.length; j++) {
        if(data[i].group_id===groupSearchData.value[j].id){
          data[i].group_id = groupSearchData.value[j].name
        }
      }
    }
    if(CurrPage.value===1){
      Total.value = info.total;
    }
    TableData.value= data
  }
  const loadGroup = async ()=>{
    Loading.value = true
    let info = await reqGetAllGroup()
    // 遍历info.data,并且修改数据字段格式
    let data = info.data;
    for (let i = 0; i < data.length; i++) {
      data[i].created_at = dayjs(data[i].created_at).format('YYYY-MM-DD HH:mm:ss')
      data[i].updated_at = dayjs(data[i].updated_at).format('YYYY-MM-DD HH:mm:ss')
    }
    groupSearchData.value.push(...data)
  }

  const search =()=>{
    if (CurrPage.value===1){
      loadData()
    }else{
      CurrPage.value = 1
    }
  }

  const onChangeAllCount = (range)=>{
    if (range ==null){
      FormSearch.value.start_time = "";
      FormSearch.value.end_time = "";
      return  false
    }
    if(range[0]===""|| range[1]===0){
      return false
    }
    FormSearch.value.start_time = range[0];
    FormSearch.value.end_time = range[1];
  }

  watch(CurrPage, (count, prevCount) => {
    loadData()
  })


</script>

<style scoped>
  .my_table_header{
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  .page_lk{
    display: flex;justify-content: center;margin-bottom: 20px
  }
</style>