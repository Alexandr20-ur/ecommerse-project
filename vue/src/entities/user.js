import { defineStore } from 'pinia'
import { notify } from "@kyvg/vue3-notification";

import axios from 'axios'

export const useUserStore = defineStore('user', {
    state: () => ({
        user: null,
        // При старте проверяем, есть ли токен в localStorage
        token: localStorage.getItem('accessToken') || null,
        // Если токен есть, значит пользователь авторизован
        isAuthenticated: !!localStorage.getItem('accessToken')
    }),
    actions: {
        async login({email, password}) {
            try {
                let response = await axios.post('http://localhost:8000/auth', {email, password})
                this.token = response.data.token
                this.isAuthenticated = true

                localStorage.setItem('accessToken', response.data.token)

                notify({
                    title: "Успех",
                    text: "Вы успешно вошли!",
                    type: "success"
                })

                return true
            } catch (error) {
                let errorMessage = "Произошла неизвестная ошибка"

                if (error.response) {
                    let data = error.response.data
                    if (data && data.Err) {
                        errorMessage = data.Err
                    } else {
                        // Если сервер вернул просто статус без тела
                        errorMessage = `Ошибка сервера: ${error.response.status}`
                    }
                } else if(error.request) {
                    // Запрос ушел, но ответа нет (сервер лежит, нет сети)
                    errorMessage = "Сервер недоступен. Проверьте подключение к интернету."
                } else {
                    errorMessage = error.message
                }

                notify({
                    title: "Ошибка",
                    type: "error",
                    text: errorMessage
                })

                console.error("Полная ошибка:", error)
               return false
            }
        },
        async logout() {
            this.token = null
            this.user = null
            localStorage.removeItem('accessToken')
        },
        async register({email, password}) {
            let response = await axios.post('http://localhost:8000/register', {email, password})

            console.log(response)
        }
    }
})