<template>
  <el-form :model="props.FormData" size="large">
    <el-form-item label="分组名称" >
      <el-input v-model="props.FormData.name" placeholder="请输入分组名称"></el-input>
    </el-form-item>
    <el-form-item label="分组描述" >
      <el-input v-model="props.FormData.desc" type="textarea" placeholder="请输入分组描述"></el-input>
    </el-form-item>
  </el-form>
</template>

<script setup>
  const props = defineProps({
    FormData: {
      type: Object,
      default: () => {
        return {}
      }
    }
  })
</script>

<style scoped>

</style>