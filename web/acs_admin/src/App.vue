<template>
  <RouterView v-loading.fullscreen.lock="usePublicStore().pageLoading" />
</template>

<script setup>
  import { usePublicStore } from '@/stores/public'
  import {onMounted, watch} from 'vue'

  onMounted(()=>{
    // 监听标题的变化并设置浏览器的标题
    watch(
        () => usePublicStore().title,
        (newTitle) => {
          document.title = newTitle;
        }
    );
  })
</script>