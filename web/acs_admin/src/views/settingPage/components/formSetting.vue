<template >
  <el-alert v-if="!isSaveDb" title="当前配置文件未保存到数据库" type="info" style="margin-bottom: 10px" show-icon :closable="false" />
  <el-alert title="注意：此处修改后会覆盖默认配置，删除后会恢复默认配置。修改配置后，需重启服务生效。" type="info" show-icon :closable="false" />
  <el-form v-if="loadingOk" v-loading="loading" label-width="150px" class="form_class" size="large">
    <el-divider content-position="left">基本配置</el-divider>
    <el-row>
      <el-col :span="12">
        <el-form-item label="调试模式" >
          <el-switch
              v-model:model-value="formData.debug"
              :inline-prompt="true"
          ></el-switch>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="后台访问监控" >
          <el-switch
              v-model:model-value="formData.open_admin_record"
              :inline-prompt="true"
          ></el-switch>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-form-item label="HTTP端口">
          <el-input v-model:model-value="formData.http_serve_port" type="number" placeholder="http运行端口"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="GRPC端口">
          <el-input v-model:model-value="formData.grpc_serve_port" type="number" placeholder="GRPC运行端口"></el-input>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-form-item label="记录留存">
          <el-input v-model:model-value="formData.record_save_days" type="number" placeholder="记录留存天数"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12">

      </el-col>
    </el-row>

    <el-divider content-position="left">登陆信息</el-divider>
    <el-row>
      <el-col :span="12">
        <el-form-item label="后台登陆用户名">
          <el-input v-model:model-value="formData.admin_user" type="text" placeholder="后台登陆用户名"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="后台登陆密码">
          <el-input v-model:model-value="formData.admin_password"  show-password type="password" placeholder="后台登陆密码"></el-input>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-form-item label="HTTP User">
          <el-input v-model:model-value="formData.http_user" type="text" placeholder="http服务用户"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="HTTP Pass">
          <el-input v-model:model-value="formData.http_password"  show-password type="password" placeholder="http服务密码"></el-input>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-form-item label="GRPC User">
          <el-input v-model:model-value="formData.grpc_user" type="text" placeholder="grpc服务用户"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="GRPC Pass">
          <el-input v-model:model-value="formData.grpc_password" show-password type="password" placeholder="grpc服务密码"></el-input>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-form-item label="Auth SECRET">
          <el-input v-model:model-value="formData.secret_key" type="text" placeholder="登陆SECRET"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="Server Expire Time">
          <el-input v-model:model-value="formData.jwt_expire_time" type="number" placeholder="服务登陆过期时间"></el-input>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-form-item label="Admin Expire Time">
          <el-input v-model:model-value="formData.jwt_admin_expire_time" type="number" placeholder="后台管理登陆过期时间"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12"></el-col>
    </el-row>


    <el-divider content-position="left">数据入库策略</el-divider>
    <el-row>
      <el-col :span="12">
        <el-form-item label="单次导入条数">
          <el-input v-model:model-value="formData.db_batch_insert_num" type="number" placeholder="数据库单次插入数据条数"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="自动保存间隔">
          <el-input v-model:model-value="formData.db_batch_insert_gap" type="number" placeholder="数据每隔多少秒保存"></el-input>
        </el-form-item>
      </el-col>
    </el-row>

    <el-divider content-position="left">日志策略</el-divider>
    <el-row>
      <el-col :span="12">
        <el-form-item label="日志保存路径">
          <el-input v-model:model-value="formData.log_file_path" type="text" placeholder="日志保存路径"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="日志保存天数">
          <el-input v-model:model-value="formData.log_file_max_age" type="number" placeholder="日志文件最大保存天数"></el-input>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <el-form-item label="日志切割间隔(小时)">
          <el-input v-model:model-value="formData.log_file_gap" type="text" placeholder="日志切割间隔(小时)"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12">

      </el-col>
    </el-row>

    <el-divider content-position="left">数据库连接信息</el-divider>
    <el-alert style="margin: 10px" title="注意：由于读取配置信息需要先连接数据库，所以数据库连接信息需要在配置文件修改。" type="info" show-icon :closable="false" />
    <template v-if="false">
      <el-row >
        <el-col :span="12">
          <el-form-item label="数据库类型">
            <el-select
                v-model:model-value="formData.db_type"
            >
              <el-option
                  v-for="(item,index) in [{'name':'mysql','open':true},{'name':'clickhouse','open':true},{'name':'mongo','open':false}]"
                  :key="index"
                  :label="item.name"
                  :value="item.name"
                  :disabled="!item.open"
              />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="12"></el-col>
      </el-row>
      <template v-if="formData.db_type === 'mysql'">
        <el-row>
          <el-col :span="12">
            <el-form-item label="Mysql Host">
              <el-input v-model:model-value="formData.mysql_host" type="text" placeholder="数据库地址"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="Mysql Port">
              <el-input v-model:model-value="formData.mysql_port" type="number" placeholder="数据库端口"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="Mysql Database">
              <el-input v-model:model-value="formData.mysql_database" type="text" placeholder="数据库名称"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="Mysql User">
              <el-input v-model:model-value="formData.mysql_username" type="text" placeholder="用户名"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="Mysql Password">
              <el-input v-model:model-value="formData.mysql_password" type="password" show-password placeholder="密码"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="测试连接">
              <el-button size="small" type="primary" @click="testMysqlConn" plain>测试连接</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </template>
      <template v-if="formData.db_type === 'clickhouse'">
        <el-row>
          <el-col :span="12">
            <el-form-item label="CH Host">
              <el-input v-model:model-value="formData.ch_host" type="text" placeholder="数据库地址"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="CH Port">
              <el-input v-model:model-value="formData.ch_port" type="number" placeholder="数据库端口"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="CH Database">
              <el-input v-model:model-value="formData.ch_database" type="text" placeholder="数据库名称"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="CH User">
              <el-input v-model:model-value="formData.ch_username" type="text" placeholder="用户名"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="CH Password">
              <el-input v-model:model-value="formData.ch_password" type="text" placeholder="密码"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
          </el-col>
        </el-row>
      </template>
    </template>
    <el-divider></el-divider>
    <el-form-item>
      <el-button type="primary" @click="onSubmit" plain>保存</el-button>
    </el-form-item>

  </el-form>
</template>

<script setup>
  import {reqGetConfig, reqSaveConfig,reqPingDb} from "@/api/system"
  import {ref, onMounted} from 'vue'
  import {cloneDeep} from 'lodash'
  import {ElMessage} from "element-plus"

  const loadingOk = ref(false)
  const loading = ref(false)
  const formData = ref({});
  const isSaveDb = ref(false);


  onMounted(() => {
    loadConfig();
  })
  const onSubmit = () => {
    loading.value = true
    let data= cloneDeep(formData.value)
    reqSaveConfig(data).then(() => {
      loading.value = false
      ElMessage.success("保存成功")
    })
  }

  const loadConfig = async () => {
    let info = await reqGetConfig();
    // json字符串转对象
    let data = JSON.parse(info.data)
    formData.value = data
    loadingOk.value = true
    if(info.is_save ==="1"){
      isSaveDb.value = true;
    }
  }

  const testMysqlConn = ()=>{
    loading.value = true
    let data = cloneDeep(formData.value)
    reqPingDb(data).then(()=>{
      loading.value = false
      ElMessage.success("测试连接成功")
    }).catch(()=>{
      loading.value = false
    })
  }
</script>


<style scoped>
  .form_class .el-form-item{
    width: 80%;
  }
</style>
