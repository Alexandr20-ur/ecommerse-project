<script setup>
import AdminSidebar from '@/components/administration/AdminSidebar.vue'
import AdminHeader from '@/components/administration/AdminHeader.vue'
</script>

<template>
  <div class="lg:d-flex w-100% min-h-100vh">
    <!-- 1. Боковое меню -->
    <AdminSidebar />

    <main class="flex-grow">
      <!-- 2. Верхняя панель -->
      <AdminHeader />

      <!-- 3. Содержимое конкретной страницы -->
      <div class="p-3 lg:p-5">
        <RouterView />
      </div>
    </main>
  </div>
</template>