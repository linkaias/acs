<template>
  <el-card class="box-card" shadow="never">
    <template #header>
      <div class="card-header">
        <span>{{props.chartName}}</span>
        <div class="right_content">
          <el-form  :inline="true">
            <el-form-item label="时间粒度">
              <el-select v-model="ChartParams.time_dimension"
                         placeholder="时间粒度"
                         @change="ChangeDimension">
                <el-option
                    v-for="(item,index) in ['min','hour','day']"
                    :key="index"
                    :label="item"
                    :value="item"
                />
              </el-select>
              <TimeInfo/>
            </el-form-item>

            <el-form-item label="时间范围" >
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
              <TimeInfo/>
            </el-form-item>
          </el-form>
          <el-icon @click="onRefreshAll" style="cursor: pointer"><Refresh /></el-icon>
        </div>
      </div>
    </template>
    <div class="chart" v-loading="loading">
      <div ref="ref_chart_container" style="height: 400px;" ></div>
    </div>
  </el-card>
</template>

<script setup>
  import TimeInfo from "@/components/TimeRangeSelectInfo.vue"
  import { Refresh } from "@element-plus/icons-vue";
  import * as echarts  from 'echarts';
  import API from  '@/api/index';
  import dayjs from 'dayjs'
  import { VerifyTimeRange, TimeRangeShortcuts } from '@/utils/time_utils'
  import {  ref, onMounted, reactive } from 'vue'
  import { ElMessage } from "element-plus";

  const ref_chart_container = ref(null)

  const props = defineProps({
    chartName: {
      type: String,
      default: 'chart'
    }
  })


  // 图表加载
  const loading = ref(false)

  // chart obj
  const chartObj = ref(null)
  // chart option
  const chartAllCountOption =  {
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        label: {
          backgroundColor: '#6a7985'
        }
      }
    },
    legend: {
      data: ['总调用次数']
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: []
    },
    yAxis: {
      type: 'value'
    },
    series: [],
  };

  // 时间筛选快捷选项
  const shortcuts = TimeRangeShortcuts();

  // 时间筛选值
  let end = dayjs()
  let start = dayjs()
  start = start.subtract(15,'minute')
  const ChartTimeRange = ref([start.format('YYYY-MM-DD HH:mm:ss'),end.format('YYYY-MM-DD HH:mm:ss')])

  // 图表默认query条件
  const ChartParams = reactive({
    time_dimension:"min",
    begin_time:(ChartTimeRange.value)[0],
    end_time:(ChartTimeRange.value)[1],
  })


  onMounted(()=>{
    chartObj.value =  echarts.init(ref_chart_container.value)
    // 渲染图表
    loadChart();
  })


  // 时间粒度修改
  const ChangeDimension = (value)=>{
    ChartParams.time_dimension = value;
    loadChart();
  }

  // 加载chart数据，渲染chart
  const loadChart = async ()=>{
    let msg = VerifyTimeRange(ChartParams.time_dimension,ChartParams.begin_time,ChartParams.end_time);
    if (msg!==""){
      ElMessage.error(msg)
      return;
    }
    loading.value = true;
    let res = await API.chartApi.reqGetAllCount(ChartParams.time_dimension, ChartParams.begin_time,ChartParams.end_time);
    loading.value = false;
    let data = res.data;
    chartAllCountOption.xAxis.data = data.x_data;
    chartAllCountOption.series =[
      {
        name: '总调用次数',
        type: 'line',
        data: data.y_data_all_count,
        areaStyle: {}
      }
    ];
    // 渲染图表
    chartAllCountOption && chartObj.value.setOption(chartAllCountOption);
    chartObj.value.setOption(chartAllCountOption);
  }

  // 手动刷新图表
  const onRefreshAll = ()=>{
    loadChart();
  }

  // 时间选择改变更新图表
  const onChangeAllCount = (range)=>{
    if(range[0]===""|| range[1]===0){
      return false
    }
    ChartParams.begin_time = range[0];
    ChartParams.end_time = range[1];
    loadChart();
  }

</script>


<style scoped>
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  .right_content{
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  .el-form-item{
    margin-bottom: 0px;
  }
</style>