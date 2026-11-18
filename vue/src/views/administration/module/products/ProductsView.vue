<script setup>
import ViewGrid from '@/components/ViewGrid.vue'
import { useRouter } from 'vue-router'
import { ref } from 'vue'

const router = useRouter()
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

const handleAction = (command) => {
  switch (command) {
    case 'add':
      router.push({name: 'admin.Products.Create'})
      break
    case 'import':
      console.log(command)
      break
    case 'groups':
      console.log(command)
      break
  }
}
</script>

<template>
  <el-row  class="mb-3">
    <el-col :span="22"><h2 class="grid-content">Товары</h2></el-col>
    <el-col :span="2">
      <el-dropdown placement="top-end" @command="handleAction">
        <el-button> Действия </el-button>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="add">Добавить товар</el-dropdown-item>
            <el-dropdown-item command="import">Импорт товаров</el-dropdown-item>
            <el-dropdown-item command="groups">Посмотреть группы товаров</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
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