<script setup>
import ViewGrid from '@/components/ViewGrid.vue'
import {
  Check,
  Delete,
  Edit,
  Message,
  Search,
  Star,
} from '@element-plus/icons-vue'

import { ref } from 'vue'

// Твои данные (придут с бэкенда)
const products = ref([
  { id: 1, date: '2016-05-03', name: 'Tom', address: 'Los Angeles' },
  { id: 2, date: '2016-05-02', name: 'Jerry', address: 'New York' },
])

// Конфигурация колонок
const productColumns = [
  { prop: 'date', label: 'Дата', width: 180 },
  { prop: 'name', label: 'Имя', width: 180 },
  { prop: 'address', label: 'Адрес', minWidth: 200 },
]

// Обработчики событий
const onEdit = (row) => console.log('Редактируем товар:', row.id)
const onDelete = (row) => console.log('Удаляем товар:', row.id)
const onView = (row) => console.log('Смотрим товар:', row.id)
</script>

<template>
  <el-row :gutter="20" class="mb-3">
    <el-col :span="16"><h2 class="grid-content">Товары</h2></el-col>
    <el-col :span="6">
      <div class="button-row">
        <el-button :icon="Search" circle />
        <el-button type="primary" :icon="Edit" circle />
        <el-button type="success" :icon="Check" circle />
        <el-button type="info" :icon="Message" circle />
        <el-button type="warning" :icon="Star" circle />
        <el-button type="danger" :icon="Delete" circle />
      </div>
    </el-col>
  </el-row>
  <!-- Использование компонента -->
  <ViewGrid
      :data="products"
      :columns="productColumns"
      :show-edit="true"
      :show-delete="true"
      :show-view="false"
      @edit="onEdit"
      @delete="onDelete"
      @view="onView"
  />
</template>