import axios from 'axios'

const apiClient = axios.create({
    baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8000'
})

apiClient.interceptors.request.use((config) => {
    const token = localStorage.getItem('accessToken')
    const publicPaths = ['/auth', '/register', '/me']

    // Проверяем, начинается ли путь с одного из публичных
    const isPublic = publicPaths.some(path => config.url?.startsWith(path))
    if (!isPublic) {
        config.url = `/api${config.url}`

    }

    if (token && config.headers) {
        config.headers.Authorization = `Bearer ${token}`
    }

    return config
})

export default apiClient