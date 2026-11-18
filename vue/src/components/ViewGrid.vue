<script setup>
import { computed } from 'vue'
import { Edit, Delete, Search } from '@element-plus/icons-vue'
import { SettingsIcon } from '@lucide/vue'

// Определяем пропсы с типами и значениями по умолчанию
const props = defineProps({
  data: { type: Array, required: true },
  columns: { type: Array, default: () => [] }, // Конфигурация колонок
  showEdit: { type: Boolean, default: false },
  showDelete: { type: Boolean, default: false },
  showView: { type: Boolean, default: false },
  actionsWidth: { type: [Number, String], default: 120 },
})

// Эмитим события наверх, чтобы родитель мог их обрабатывать
const emit = defineEmits(['edit', 'delete', 'view'])

const handleEdit = (row) => emit('edit', row)
const handleDelete = (row, index) => emit('delete', row, index)
const handleView = (row) => emit('view', row)

// Вычисляем, нужно ли показывать колонку действий
const hasActions = computed(() =>
    props.showEdit || props.showDelete || props.showView
)
</script>

<template>
  <el-table :data="data" style="width: 100%" border>
    <!-- Динамический рендер колонок из конфига -->
    <el-table-column
        v-for="col in columns"
        :key="col.prop"
        :prop="col.prop"
        :label="col.label"
        :width="col.width"
        :min-width="col.minWidth"
        :align="col.align || 'left'"
    >
      <!-- Если передан слот для этой колонки - рендерим его, иначе просто текст -->
      <template #default="scope">
        <slot :name="col.prop" :row="scope.row" :index="scope.$index">
          {{ scope.row[col.prop] }}
        </slot>
      </template>
    </el-table-column>

    <!-- Колонка действий (показывается только если включена хотя бы одна кнопка) -->
    <el-table-column
        v-if="hasActions"
        label="Действия"
        :width="actionsWidth"
        align="center"
        fixed="right"
    >
      <template #header>
        <div class="header-icon-wrapper">
          <SettingsIcon size="16" />
        </div>
      </template>

      <template #default="{ row, $index }">
        <div class="action-buttons">
          <el-button
              v-if="showView"
              type="info"
              :icon="Search"
              circle
              size="small"
              @click="handleView(row)"
              title="Просмотр"
          />
          <el-button
              v-if="showEdit"
              type="primary"
              :icon="Edit"
              circle
              size="small"
              @click="handleEdit(row)"
              title="Редактировать"
          />
          <el-button
              v-if="showDelete"
              type="danger"
              :icon="Delete"
              circle
              size="small"
              @click="handleDelete(row, $index)"
              title="Удалить"
          />
        </div>
      </template>
    </el-table-column>
  </el-table>
</template>

<style scoped>
.header-icon-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}

.action-buttons {
  display: flex;
  gap: 8px;
  justify-content: center;
  align-items: center;
}
</style>