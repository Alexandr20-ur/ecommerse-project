export function setupRouterGuards(router) {
  router.beforeEach((to, from) => {
    const isAuthenticated = !!localStorage.getItem('accessToken')
    if (!isAuthenticated) {
      return { name: 'home'}
    }
  })
}