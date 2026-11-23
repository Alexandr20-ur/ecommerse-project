import { defineStore } from 'pinia'
import { notify } from '@kyvg/vue3-notification'
import apiClient from '@/utils/axios.ts'
import {AxiosError} from "axios";

// 1. Интерфейс для ответа сервера при авторизации
interface LoginResponse {
    token: string
    Err?: string // Поле ошибки из твоего Go-бэкенда
}

// 2. Интерфейс для состояния стора
interface UserState {
    user: object | null // Можно заменить на конкретный интерфейс IUser, если есть
    token: string | null
    isAuthenticated: boolean
}

export const useUserStore = defineStore('user', {
    state: (): UserState => ({
        user: null,
        token: localStorage.getItem('accessToken') || null,
        isAuthenticated: !!localStorage.getItem('accessToken'),
    }),

    actions: {
        async login({ email, password }: { email: string; password: string }): Promise<boolean> {
            try {
                const response = await apiClient.post<LoginResponse>('/auth', {
                    email,
                    password,
                })

                this.token = response.data.token
                this.isAuthenticated = true
                localStorage.setItem('accessToken', response.data.token)

                notify({
                    title: 'Успех',
                    text: 'Вы успешно вошли!',
                    type: 'success',
                })

                return true
            } catch (error) {
                let errorMessage = 'Произошла неизвестная ошибка'

                // Типизируем ошибку как AxiosError для доступа к response/request
                if (error instanceof AxiosError) {
                    if (error.response) {
                        const data = error.response.data as LoginResponse
                        if (data?.Err) {
                            errorMessage = data.Err
                        } else {
                            errorMessage = `Ошибка сервера: ${error.response.status}`
                        }
                    } else if (error.request) {
                        errorMessage = 'Сервер недоступен. Проверьте подключение к интернету.'
                    } else {
                        errorMessage = error.message || errorMessage
                    }
                } else if (error instanceof Error) {
                    errorMessage = error.message
                }

                notify({
                    title: 'Ошибка',
                    type: 'error',
                    text: errorMessage,
                })

                console.error('Полная ошибка:', error)
                return false
            }
        },

        logout(): void {
            this.token = null
            this.user = null
            this.isAuthenticated = false // Не забываем сбрасывать флаг!
            localStorage.removeItem('accessToken')
        },

        async register({ email, password }: { email: string; password: string }): Promise<void> {
            try {
                const response = await apiClient.post('/register', {
                    email,
                    password,
                })
                console.log('Регистрация успешна:', response)

                notify({
                    title: 'Успех',
                    text: 'Регистрация прошла успешно! Теперь войдите.',
                    type: 'success'
                })
            } catch (error) {
                // Можно добавить обработку ошибок аналогично login
                console.error('Ошибка регистрации:', error)
                notify({
                    title: 'Ошибка',
                    text: 'Не удалось зарегистрироваться',
                    type: 'error'
                })
            }
        },
    },
})