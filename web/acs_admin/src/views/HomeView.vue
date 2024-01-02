<template>
  <el-container class="layout-container" style="min-height: 500px">
    <el-header class="header_lk" >
      <div class="header_left">
        <el-image style="height: 80%" :src="LogoImg" title="ACS View Database" ></el-image>
      </div>
      <div class="header_right">
        <div>
          <el-breadcrumb separator="/">
            <template v-for="(item,index) in usePublicStore().breadcrumb" :key="index" >
              <el-breadcrumb-item v-if="item.path!==''" :to="{ path: item.path }">
                {{item.name}}
              </el-breadcrumb-item>
              <el-breadcrumb-item v-else :to="{ path: '/' }">
                {{item.name}}
              </el-breadcrumb-item>
            </template>
          </el-breadcrumb>
        </div>
        <div class="user_info" >
          <el-dropdown trigger="click">
              <span class="el-dropdown-link">
                {{GetUserName()||'未登录'}}
                <el-icon class="el-icon--right">
                  <arrow-down />
                </el-icon>
              </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="logout">退出登陆</el-dropdown-item>
                <el-dropdown-item>
                  <a class="link_class" href="https://funcs.gitbook.io/acs_docs/" target="_blank" >
                    文档
                  </a>
                </el-dropdown-item>
                <el-dropdown-item>
                  <a class="link_class" href="https://github.com/linkaias/acs" target="_blank" >
                    Github
                  </a>
                </el-dropdown-item>
                <el-dropdown-item>
                  <a class="link_class" href="https://gitee.com/tolinkai/acs" target="_blank" >
                    Gitee
                  </a></el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
    </el-header>
    <el-container>
      <el-aside width="201px"  class="left_content" >
        <MenuView :menus="menus" />
      </el-aside>
      <el-container>
        <el-main >
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </el-container>
</template>

<script setup>
  import MenuView from './menu/MenuView.vue'
  import { ArrowDown } from "@element-plus/icons-vue"
  import {GetUserName, RemoveToken} from '@/utils/token_utils'
  import { ElMessage } from 'element-plus'
  import { usePublicStore } from  '@/stores/public'

  import LogoImg from '@/assets/logo.png'

  // index 对应router中的name 用于菜单跳转
  const menus  = [
    {
      index:'home',
      title: '数据图表',
      icon: 'DataLine',
      children: [
        {
          index:'/home/index',
          title: '总览'
        },
        {
          index:'/home/group',
          title: '请求分组概览'
        },
        {
          index:'/home/method',
          title: '请求方式概览'
        },
        // {
        //   index:'/home/top',
        //   title: '请求Top概览'
        // } 饼图展示，暂时不用
      ]
    },
    {
      index:'sys',
      title: '系统管理',
      icon: 'Setting',
      children: [
        {
          index:'/sys/group_list',
          title: '分组管理'
        },
        {
          index:'/sys/record',
          title: '记录管理'
        }
      ]
    },
    {
      index:'/setting/system',
      title: '设置',
      icon: 'Setting',
      children: []
    }
  ]

  const logout = ()=>{
    RemoveToken()
    ElMessage.success("退出成功！")
    setTimeout(function () {
      window.location.href = '/login'
    },1200)
  }
  const handleClickLogo = ()=>{
    console.log(111)
  }
</script>

<style scoped>
  .user_info{
    display: flex;
    justify-content: flex-end;
    align-items: center;
    font-size: 16px;
    height: 60px;
  }
  .header_lk{
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: white; box-shadow: 0 1px 4px rgba(0,21,41,.08);
    padding: 0;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 999;
  }
  .header_left {
    display: flex;
    justify-content: center;
    align-items: center;
    font-weight: bolder;
    flex: 0 0 auto;
    width: 200px;
    height: 60px;
    border-right: solid 1px var(--el-menu-border-color);
  }
  .header_right{
    padding: 0 20px;
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .left_content{
    border-right: solid 1px var(--el-menu-border-color);
    margin-top: 1px;
  }

  .footer_msg{
    text-align: center; font-size: 16px; color: #666666;
    margin-bottom: 20px;
  }
  .left_affix{
    border-right: solid 1px var(--el-menu-border-color);
  }
  .link_class{
    text-decoration: none;
    color: #606266;
  }
</style>