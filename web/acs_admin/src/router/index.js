import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import {usePublicStore} from  '@/stores/public'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/home/index'
    },
    {
      path: '/home',
      name: 'home',
      component: HomeView,
      children:[
        {
          path: 'index',
          component: ()=>import('../views/indexPage/index.vue')
        } ,
        {
          path: 'group',
          name:"group",
          component: ()=>import('../views/byGroupPage/index.vue')
        },
        {
          path: 'method',
          name:"method",
          component: ()=>import('../views/byMethodPage/index.vue')
        },
        {
          path: 'top',
          name:"top",
          component: ()=>import('../views/byTopPage/index.vue')
        }
      ],
    },
    {
      path: '/sys',
      name: 'sys',
      component: HomeView,
      children:[
        {
          path: 'group_list',
          name: "group_list",
          component: ()=>import('../views/systemGroupList/index.vue')
        } ,
        {
          path: 'record',
          name:"record",
          component: ()=>import('../views/systemRecordList/index.vue')
        }
      ],
    },
    {
      path: '/setting',
      name: 'setting',
      component: HomeView,
      children:[
        {
          path: 'system',
          name: "system",
          component: ()=>import('../views/settingPage/index.vue')
        }
      ],
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/login/LoginView.vue')
    }
  ],
  scrollBehavior (to, from, savedPosition) {
    return {y: 0}
  }
})

router.beforeEach((to, from)=>{
  usePublicStore().SetPageLoading(true)
  //return false 取消导航
})

router.afterEach((to, from)=>{
  usePublicStore().SetPageLoading(false)
})

export default router
