<script>
import { importImages } from '@/utils/images'
import { useUserStore } from '@/entities/user.ts';
import router from "@/router/index.js";

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
      let success = await userStore.login(this.formData)
      if (success) {
        router.push('/')
      }
    }
  }
}
</script>
<template>
<section class="py-20 bg-lower min-h-100vh">

    <div class="card shadow-lg p-5 lg:p-8 mx-auto" style="max-width: 420px;">
      <div class="mb-4 text-center">
        <h4 class="fs-xl mb-2">Авторизация</h4>
      </div>
      
      <!-- line end.// -->
      <form @submit.prevent="handleSubmit" novalidate>
           <fieldset>
        <div class="mb-4">
          <label class="form-label fw-medium">Почта или логин</label>
          <input v-model="formData.email" type="text" class="form-control" placeholder="Введите почту или логин">
        </div>
        <div class="mb-2">
          <label class="form-label fw-medium">Пароль</label>
          <input v-model="formData.password" type="password" class="form-control" placeholder="Введите ваш пароль">
        </div>
        <div class="mb-6">
          <a href="#" class="fs-sm">Забыли пароль?</a>
        </div>
        <button class="btn btn-primary w-100%">Войдите в свою учетную запись</button>
      </fieldset>
      </form>

      <hr class="my-6">

      <p class="text-center text-secondary">У вас нет аккаунта? <RouterLink to="/register">Зарегестрируйтесь сейчас</RouterLink></p>
    </div>
</section>
</template>