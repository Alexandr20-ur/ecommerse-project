import { defineStore } from 'pinia'

import axios from 'axios'

export const useUserStore = defineStore('user', {
    state: () => ({
        user: null,
        token: null,
        isAuthenticated: false
    }),
    actions: {
        async login({login, password}) {
            let response = await axios.post('http://localhost:8000/auth', {login, password})

            console.log(response, "login")
        },
        async register({email, password}) {
            let response = await axios.post('http://localhost:8000/register', {email, password})

            console.log(response)
        }
    }
})