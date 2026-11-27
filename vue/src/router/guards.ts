import {type Router} from 'vue-router' // <-- Добавь ключевое слово type
import {useUserStore} from '@/entities/user' // Исправил путь на стандартный для Pinia stores

export function setupRouterGuards(router: Router) {
    router.beforeEach(async (to) => {
        const userStore = useUserStore()

        // 1. Ждем завершения первоначальной проверки токена, если она еще идет
        if (userStore.isLoading) {
            await userStore.checkAuth()
        }

        // 2. Если страница только для гостей (Логин/Регистрация)
        if (to.meta.isGuestOnly) {
            if (userStore.isAuthenticated) {
                return {name: 'home'}
            }
        }

        // 3. Если страница требует авторизации
        if (to.meta.requiresAuth) {
            if (!userStore.isAuthenticated) {
                return {name: 'login'}
            }

            // 4. Дополнительная проверка на роль админа
            if (to.meta.requiresAdmin && userStore.user?.role !== 'admin') {
                // Можно редиректить на главную или страницу "Доступ запрещен"
                return {name: 'home'}
            }
    }
  })
}