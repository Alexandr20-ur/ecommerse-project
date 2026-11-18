<script setup>
import { RouterLink, useRoute, useRouter } from 'vue-router'
import {
  Store,
  LayoutDashboard,
  Truck,
  Package2,
  FileSpreadsheet,
  CircleDollarSign,
  Users,
  Tag,
  BarChart3,
  Settings
} from '@lucide/vue'
import { computed } from 'vue'

const route = useRoute()
const router = useRouter()

// Массив пунктов меню для чистоты кода
const menuItems = [
  { name: 'Дашбоард', routeName: 'admin.Dashboard', icon: LayoutDashboard },
  // { name: 'Все доставки', routeName: '/administration/deliveries', icon: Truck },
  { name: 'Товары', routeName: 'admin.Products', icon: Package2 },
  // { name: 'Заказы', routeName: '/administration/orders', icon: FileSpreadsheet },
  // { name: 'Финансы', routeName: '/administration/finance', icon: CircleDollarSign },
  // { name: 'Клиенты', routeName: '/administration/customers', icon: Users },
  // { name: 'Скидки', routeName: '/administration/discounts', icon: Tag },
  // { name: 'Отчеты', routeName: '/administration/reports', icon: BarChart3 },
]

// Вычисляем активный пункт на основе текущего роута
const activeIndex = computed(() => route.name)
const resolvePath = (routeName) => {
  return router.resolve({name: routeName}).href
}
</script>

<template>
  <aside class="sidebar">
    <!-- Логотип -->
    <div class="sidebar-header">
      <RouterLink to="/administration/dashboard" class="logo-link">
        <Store size="24" />
        <span class="logo-text">Админка</span>
      </RouterLink>
    </div>

    <!-- Меню Element Plus -->
    <el-menu
        :default-active="activeIndex"
        class="sidebar-menu"
    >
      <el-menu-item
          v-for="item in menuItems"
          :key="item.routeName"
          :index="item.routeName"
          @click="$router.push({ name: item.routeName })"
      >
        <component :is="item.icon" size="20" class="menu-icon" />
        <template #title>{{ item.name }}</template>
      </el-menu-item>

      <!-- Разделитель и настройки -->
      <div class="menu-divider"></div>

      <el-menu-item index="/admin/settings">
        <Settings size="20" class="menu-icon" />
        <template #title>Настройки</template>
      </el-menu-item>
    </el-menu>
  </aside>
</template>

<style scoped>
.sidebar {
  width: 240px; /* Фиксированная ширина как в твоем дизайне */
  height: 100vh;
  background-color: #f8fafc; /* Светлый фон как на скриншоте */
  border-right: 1px solid #e2e8f0;
  display: flex;
  flex-direction: column;
}

.sidebar-header {
  height: 64px;
  display: flex;
  align-items: center;
  padding: 0 20px;
  border-bottom: 1px solid #e2e8f0;
}

.logo-link {
  display: flex;
  align-items: center;
  gap: 12px;
  text-decoration: none;
  color: #0f172a;
  font-weight: 600;
  font-size: 18px;
}

/* Стилизация el-menu под твой дизайн */
.sidebar-menu {
  border-right: none !important; /* Убираем стандартную границу el-menu */
  flex-grow: 1;
  overflow-y: auto;
}

/* Переопределяем стили элементов меню */
:deep(.el-menu-item) {
  height: 48px !important;
  line-height: 48px !important;
  margin: 4px 8px !important;
  border-radius: 8px !important;
  font-size: 14px !important;
}

:deep(.el-menu-item.is-active) {
  background-color: #e2e8f0 !important;
  color: #0f172a !important;
  font-weight: 500 !important;
}

:deep(.el-menu-item:hover:not(.is-active)) {
  background-color: #f1f5f9 !important;
}

.menu-icon {
  margin-right: 12px !important;
  color: #64748b;
}

:deep(.el-menu-item.is-active .menu-icon) {
  color: #0f172a;
}

.menu-divider {
  height: 1px;
  background-color: #e2e8f0;
  margin: 8px 16px;
}
</style>