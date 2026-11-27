import {defineStore} from 'pinia'
import {notify} from '@kyvg/vue3-notification'
import apiClient from '@/utils/axios.ts'
import {AxiosError} from 'axios'

// Интерфейс, который 1-в-1 совпадает с Go-структурой UserResponse
export interface IUser {
    id: string
    email: string
    phone?: string
    full_name: string
    avatar_url?: string
    role: string
}

interface UserState {
    user: IUser | null
    token: string | null
    isAuthenticated: boolean
    isLoading: boolean
}

export const useUserStore = defineStore('user', {
    state: (): UserState => ({
        user: null,
        token: localStorage.getItem('accessToken') || null,
        isAuthenticated: false,
        isLoading: true,
    }),

    actions: {
        // Вспомогательная функция для красивых ошибок
        _handleError(error: unknown, defaultMessage: string) {
            let errorMessage = defaultMessage

            if (error instanceof AxiosError) {
                if (error.response) {
                    // Бэкенд всегда присылает { error: "..." }
                    const data = error.response.data
                    if (data?.error) {
                        errorMessage = data.error
                    } else {
                        errorMessage = `Ошибка сервера: ${error.response.status}`
                    }
                } else if (error.request) {
                    errorMessage = 'Сервер недоступен. Проверьте подключение к интернету.'
                } else {
                    errorMessage = error.message
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
        },

        async checkAuth(): Promise<void> {
            if (!this.token) {
                this.isLoading = false
                return
            }

            try {
                // Бэкенд отдает плоский объект IUser
                const response = await apiClient.get<IUser>('/me')
                this.user = response.data
                this.isAuthenticated = true
            } catch (error) {
                console.warn('Token invalid or expired')
                this.logout()
            } finally {
                this.isLoading = false
            }
        },

        async login({email, password}: { email: string; password: string }): Promise<boolean> {
            try {
                // Бэкенд отдает { token: "..." }
                const response = await apiClient.post<{ token: string }>('/auth', {
                    email,
                    password,
                })

                this.token = response.data.token
                localStorage.setItem('accessToken', response.data.token)

                // Сразу проверяем токен и подтягиваем юзера
                await this.checkAuth()

                notify({
                    title: 'Успех',
                    text: 'Вы успешно вошли!',
                    type: 'success',
                })

                return true
            } catch (error) {
                this._handleError(error, 'Не удалось войти')
                return false
            }
        },

        async register({email, password}: { email: string; password: string }): Promise<boolean> {
            try {
                await apiClient.post('/register', {email, password})

                notify({
                    title: 'Успех',
                    text: 'Регистрация прошла успешно! Теперь войдите.',
                    type: 'success'
                })
                return true
            } catch (error) {
                this._handleError(error, 'Не удалось зарегистрироваться')
                return false
            }
        },

        logout(): void {
            this.token = null
            this.user = null
            this.isAuthenticated = false
            localStorage.removeItem('accessToken')
        },
    },
})