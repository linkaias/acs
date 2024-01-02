<template>
    <el-card shadow="never">
      <template #header>
        <div class="card-header">
          <span>请求分组列表</span>
        </div>
      </template>
      <div  class="my_table_header">
        <el-form inline>
          <el-form-item>
            <el-button   :icon="Plus" plain @click="addGroup" >新增分组</el-button>
          </el-form-item>
        </el-form>
        <el-form inline>
          <el-form-item>
            <el-button :icon="Refresh" circle @click="loadGroup" />
          </el-form-item>
        </el-form>
      </div>
      <el-table size="large" border v-loading="groupLoading" :data="groupTableData" style="width: 100%;margin-bottom: 50px">
        <el-table-column type="index" label="序号" width="100" />
        <el-table-column prop="id" label="分组ID" width="100" />
        <el-table-column prop="name" label="分组名称" width="250" />
        <el-table-column prop="desc" label="分组描述" />
        <el-table-column prop="updated_at" label="更新时间" width="250"  />
        <el-table-column prop="created_at" label="创建时间" width="250"  />
        <el-table-column  label="操作" width="250" >
          <template #default="{row, $index}">
            <el-tooltip
                v-if="row.id===1 || row.id===2"
                class="box-item"
                effect="light"
                content="系统预设分组不能修改"
                placement="bottom-start"
            >
              <el-button disabled  size="small" type="primary" @click="editGroup(row)">编辑</el-button>
            </el-tooltip>
            <el-button v-else  size="small" type="primary" @click="editGroup(row)">编辑</el-button>
            <el-tooltip
                v-if="row.id===1 || row.id===2"
                class="box-item"
                effect="light"
                content="系统预设分组不能删除"
                placement="bottom-start"
            >
              <el-button disabled size="small" type="danger" @click="deleteGroup(row)">删除</el-button>
            </el-tooltip>
            <el-button v-else size="small" type="danger" @click="deleteGroup(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>


  <el-dialog @closed="closedForm" v-model="FormVisible" :title="FormData.id ? '编辑分组': '新增分组'" width="600px">
    <TableForm :FormData.sync="FormData"/>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="closeForm">取消</el-button>
        <el-button type="primary" @click="confirmForm">
          确认
        </el-button>
      </span>
    </template>
  </el-dialog>

</template>

<script setup>
  import TableForm from './components/form.vue'
  import {reqGetAllGroup, reqAddGroup, reqUpdateGroup, reqDelGroup} from "@/api/group";
  import {cloneDeep} from 'lodash'
  import dayjs from "dayjs";
  import {onMounted, ref, reactive} from "vue";
  import { Refresh,Plus } from '@element-plus/icons-vue'
  import { usePublicStore} from '@/stores/public'
  import {ElMessage, ElMessageBox} from "element-plus";

  const storePublic = reactive(usePublicStore())
  const groupLoading = ref(false)
  const groupTableData = ref([])
  const breadcrumb = [
    {
      path: '/',
      name: '首页'
    },
    {
      path: '',
      name: '请求分组管理'
    }
  ]
  const FormVisible = ref(false)
  const FormData = ref({
    id:0,
    name:'',
    desc:''
  })


  onMounted(()=>{
    storePublic.SetTitle('请求分组管理');
    storePublic.SetBreadcrumb(breadcrumb);
    loadGroup();
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

  const editGroup=(row)=>{
    FormData.value={
      id: row.id,
      name: row.name,
      desc: row.desc
    }
    FormVisible.value = true
  }

  const deleteGroup=(row)=>{
    ElMessageBox.confirm(
        '确定删除['+row.name+']分组吗?',
        '提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
    )
        .then(() => {
          apiDelGroup(row.id)
        })
        .catch(() => {})
  }
  const closeForm = ()=>{
    FormVisible.value = false
  }
  const closedForm = ()=>{
    FormData.value={
      id: 0,
      name: '',
      desc: ''
    }
  }
  const confirmForm = ()=>{
    apiSaveGroup()
  }
  const addGroup=()=>{
    FormVisible.value = true
  }

  const apiSaveGroup = async ()=>{
    let data = cloneDeep(FormData.value)
    if (data.name===""){
      ElMessage.error("分组名称不能为空")
      return
    }
    if(data.id>0){ // 修改
      let res = await reqUpdateGroup(data)
      await loadGroup();
      ElMessage.success(res.msg);
    }else{// 新增
      let res = await reqAddGroup(data)
      await loadGroup();
      ElMessage.success(res.msg);
    }
    FormVisible.value = false
  }

  const apiDelGroup = async (id)=>{
    let res = await reqDelGroup(id)
    await loadGroup();
    ElMessage.success(res.msg);
  }

</script>

<style scoped>
  .my_table_header{
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
</style>