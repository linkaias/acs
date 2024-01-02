
<template>
  <el-popover
      placement="top-start"
      title="提示"
      :width="300"
      trigger="hover"
  >
    <template #reference>
      <el-icon class="icon">
        <warning></warning>
      </el-icon>
    </template>
    <p style="font-weight: bold">关于时间筛选限制说明</p>
    <p>1.如果时间粒度是分钟, 那么开始时间和结束时间的间隔必须大于1分钟，并且间隔最大不能超过1天</p>
    <p>2.如果时间粒度是小时, 那么开始时间和结束时间的间隔必须大于1小时，并且间隔最大不能超过31天</p>
    <p>3.如果时间粒度是天, 那么开始时间和结束时间的间隔必须大于1天，并且间隔最大不能超过1年</p>

  </el-popover>

</template>

<script setup>
  import { Warning } from "@element-plus/icons-vue"
</script>

<style scoped>
  .icon{
    font-size: 20px;
    margin: 0 5px;
    cursor: pointer;
  }
</style>