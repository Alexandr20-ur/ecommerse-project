export function setupRouterGuards(router) {
  router.beforeEach((to, from, next) => {
    const isAuthenticated = !!localStorage.getItem('accessToken')
    if (to.meta.requiresAuth && !isAuthenticated) {
      next({ name: 'home' })
    } else {
      next()
    }
  })
}