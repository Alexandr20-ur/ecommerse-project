<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Undo } from '@lucide/vue'

const router = useRouter()
const formRef = ref()

// Начальное состояние формы под нашу таблицу products
const form = reactive({
  name: '',
  description: '',
  price: null,
  stock: 0,
})

// Правила валидации
const rules = reactive({
  name: [
    { required: true, message: 'Введите название товара', trigger: 'blur' },
    { min: 3, max: 100, message: 'От 3 до 100 символов', trigger: 'blur' },
  ],
  price: [
    { required: true, message: 'Укажите цену', trigger: 'blur' },
    { type: 'number', min: 0.01, message: 'Цена должна быть больше 0', trigger: 'blur' },
  ],
  stock: [
    { required: true, message: 'Укажите остаток', trigger: 'blur' },
    { type: 'number', min: 0, message: 'Остаток не может быть отрицательным', trigger: 'blur' },
  ],
})

// Обработка отправки
const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      console.log('Данные для отправки на Go-бэкенд:', JSON.parse(JSON.stringify(form)))
      // TODO: Здесь будет вызов API createProduct(form)
      // После успеха: router.push({ name: 'admin.Products' })
    }
  })
}

const goBack = () => router.back()
</script>

<template>
  <!-- Шапка с кнопкой назад -->
  <el-row class="mb-4" align="middle">
    <el-col :span="22">
      <h2 class="text-xl font-bold text-gray-800">Создание товара</h2>
    </el-col>
    <el-col :span="2" style="display: flex; justify-content: flex-end;">
      <el-button :icon="Undo" plain @click="goBack">Назад</el-button>
    </el-col>
  </el-row>

  <!-- Форма создания -->
  <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="120px"
      style="max-width: 600px"
      @submit.prevent="handleSubmit"
  >
    <el-form-item label="Название" prop="name">
      <el-input v-model="form.name" placeholder="Например: Кроссовки Nike Air" />
    </el-form-item>

    <el-form-item label="Описание" prop="description">
      <el-input
          v-model="form.description"
          type="textarea"
          :rows="4"
          placeholder="Подробное описание характеристик товара..."
      />
    </el-form-item>

    <el-row :gutter="20">
      <el-col :span="12">
        <el-form-item label="Цена (₽)" prop="price">
          <el-input-number
              v-model="form.price"
              :min="0.01"
              :precision="2"
              :step="10"
              controls-position="right"
              style="width: 100%"
          />
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="Остаток" prop="stock">
          <el-input-number
              v-model="form.stock"
              :min="0"
              :step="1"
              controls-position="right"
              style="width: 100%"
          />
        </el-form-item>
      </el-col>
    </el-row>

    <el-form-item style="margin-top: 30px;">
      <el-button type="primary" size="large" @click="handleSubmit" :loading="false">
        Создать товар
      </el-button>
      <el-button size="large" @click="formRef?.resetFields()">Сбросить</el-button>
    </el-form-item>
  </el-form>
</template>

<style scoped>
/* Дополнительные стили для красоты */
:deep(.el-input-number .el-input__inner) {
  text-align: left;
}
</style>