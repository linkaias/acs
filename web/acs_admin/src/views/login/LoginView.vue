<template>
  <div class="main_container">
      <div class="login">
       <div class="title">
         <el-image style="height: 50%" :src="LogoImg" title="ACS View Database" ></el-image>
         <p>Admin Login </p>
       </div>
        <div class="login_body">
          <el-form :model="formData"
                   :rules="rules"
                   ref="formRef"
          >
            <el-form-item prop="user" >
              <el-input  :prefix-icon="User" size="large" class="logon_input" v-model="formData.user" placeholder="username" />
            </el-form-item>
            <el-form-item prop="password" >
              <el-input  :prefix-icon="Lock" type="password"  show-password size="large" class="logon_input" v-model="formData.password" placeholder="password" />
            </el-form-item>

            <el-form-item>
              <el-button @click="loginSubmit()" type="primary" plain size="large" style="height: 45px; width: 100%">Login</el-button>
            </el-form-item>

            <div class="login_text">
              <el-link :underline="false" type="info" href="https://funcs.gitbook.io/acs_docs/" target="_blank">文档</el-link>
              <el-link :underline="false" type="info" href="https://github.com/linkaias/acs" target="_blank">Github</el-link>
              <el-link :underline="false" type="info" href="https://gitee.com/tolinkai/acs" target="_blank">Gitee</el-link>
            </div>
          </el-form>

        </div>
      </div>
  </div>
</template>

<script setup>
  import { User, Lock } from '@element-plus/icons-vue'
  import { reactive, ref} from 'vue'
  import { reqGetToken } from '@/api/user/index'
  import { SetToken, SetUserName } from  '@/utils/token_utils'
  import { ElMessage } from 'element-plus'
  import LogoImg from '@/assets/logo.png'


  const formRef = ref(null)
  const rules = reactive({
    user: [
      { required: true, message: 'Please input username', trigger: 'blur' },
      { min: 5, max: 20, message: 'Length should be 5 to 20', trigger: 'blur' },
    ],
    password: [
      { required: true, message: 'Please input password', trigger: 'blur' },
      { min: 6, max: 20, message: 'Length should be 6 to 20', trigger: 'blur' },
    ],
  })
  const formData = reactive({
    user: 'admin',
    password: '123456',
  })


  const submitLogin = async () => {
    let res = await reqGetToken(formData);
    if (res.data.token!==""){
      ElMessage.success("登录成功！")
      SetToken(res.data.token)
      SetUserName(res.data.username)
      //跳转到首页
      setTimeout(function () {
        window.location.href = '/'
      },1500)
    }else{
      ElMessage.error("登录失败！")
    }
  }

  const loginSubmit = async () => {
    await formRef.value.validate((valid, fields) => {
      if (valid) {
        submitLogin()
      } else {
        //console.log('error submit!', fields)
      }
    })
  }


</script>

<style scoped>
  .main_container {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .login {
    margin-top: 8%;
    width: 400px;
    height: 500px;
    background-color: #fff;
  }
  .login .title{
    width: 100%;
    height: 80px;
    display: flex;
    justify-content: center;
    align-items: center;
    color: #333;
    font-weight: bolder;
    font-size: 30px;
  }
  .login_body{
    margin-top: 20px;
  }
  .logon_input {
    height: 45px;
  }

  .login_text{
    text-align: center;
  }

  .login_text a{
    margin: 0 10px;
  }
</style>