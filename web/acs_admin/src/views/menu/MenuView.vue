<template>
  <el-menu
      :default-active="defaultMenuIndex"
      class="el-menu-vertical-demo"
      :collapse="false"
      :router="true"
  >
    <template v-for="(menu, index) in menus" :key="index">
      <el-sub-menu  :index="menu.index" v-if="menu.children.length > 0">
        <template #title>
          <span>{{ menu.title }}</span>
        </template>
        <el-menu-item
            v-for="(child, childIndex) in menu.children"
            :key="childIndex"
            :index="child.index"
        >
          {{ child.title }}
        </el-menu-item>
      </el-sub-menu>
      <el-menu-item  :index="menu.index" v-else>
        <span>{{ menu.title }}</span>
      </el-menu-item>
    </template>
  </el-menu>
</template>

<script  setup>
  import {onMounted, ref} from "vue";
  import { useRoute } from 'vue-router';
  // 获取路由实例
  const route = useRoute();

  const props  = defineProps({
    menus: {
      type: Array,
      default: []
    }
  })
  const defaultMenuIndex = ref(null)

  onMounted(()=>{
    updateSelectedMenu(route.path)
  })

  const updateSelectedMenu = (path)=> {
    for (let i = 0; i < props.menus.length; i++) {
      if ((props.menus[i].children).length>0) {
        for (let j = 0; j < props.menus[i].children.length; j++) {
          const child = props.menus[i].children[j];
          if (child.index === path) {
            defaultMenuIndex.value = path;
            return;
          }
        }
      } else if (props.menus[i].index === path) {
        defaultMenuIndex.value = `${i}`;
        return;
      }
    }
  }

</script>


<style scoped>
.el-menu-vertical-demo{
  border: none!important;
}
</style>