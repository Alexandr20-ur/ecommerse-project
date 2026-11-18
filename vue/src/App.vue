<script setup>
import { RouterView } from 'vue-router'
import { Notifications } from '@kyvg/vue3-notification'
import { ElConfigProvider } from 'element-plus'
import ru from 'element-plus/es/locale/lang/ru'
</script>

<template>
  <el-config-provider :locale="ru">
    <RouterView />
    <Notifications />
  </el-config-provider>
</template>