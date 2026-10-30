import { defineStore } from 'pinia'
import { notify } from "@kyvg/vue3-notification";

import axios from 'axios'

export const useUserStore = defineStore('user', {
    state: () => ({
        user: null,
        token: null,
        isAuthenticated: false
    }),
    actions: {
        async login({login, password}) {
            try {
                let response = await axios.post('http://localhost:8000/auth', {login, password})
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