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
            } catch (error) {
                let data = error.response.data
                if (!data) {
                    notify({
                        title: "Ошибка",
                        type: "error",
                        text: error
                    })
                    return
                }

                notify({
                    title: "Ошибка",
                    text: data.Err,
                    type: "error"
                })
                console.log(error.response.data)
            }
        },
        async register({email, password}) {
            let response = await axios.post('http://localhost:8000/register', {email, password})

            console.log(response)
        }
    }
})