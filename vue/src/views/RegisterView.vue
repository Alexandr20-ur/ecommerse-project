<script>
import { importImages } from '@/utils/images'
import { useUserStore } from '@/entities/user';
const images = importImages()

export default {
  data() {
    return {
      formData: {
        email: "",
        password: ""
      }
    }
  },
  methods: {
    async handleSubmit() {
      let userStore = useUserStore()
      userStore.register(this.formData)
    }
  }
}
</script>
<template>
<section class="py-20 bg-lower min-h-100vh">

    <div class="card shadow-lg p-5 lg:p-8 mx-auto" style="max-width: 420px;">
      <div class="mb-4 text-center">
        <h4 class="fs-xl mb-2">Регистрация</h4>
      </div>
      
      <!-- line end.// -->
      <form @submit.prevent="handleSubmit" novalidate>
           <fieldset>
        <div class="mb-4">
          <label class="form-label fw-medium">Почта или логин</label>
          <input v-model="formData.email" type="text" class="form-control" placeholder="Введите почту">
        </div>
        <div class="mb-2">
          <label class="form-label fw-medium">Пароль</label>
          <input v-model="formData.password" type="password" class="form-control" placeholder="Введите ваш пароль">
        </div>
        <div class="mb-6">
          <a href="#" class="fs-sm">Забыли пароль?</a>
        </div>
        <button class="btn btn-primary w-100%">Регистрация</button>
      </fieldset>
      </form>

      <hr class="my-6">

    </div>
</section>
</template>