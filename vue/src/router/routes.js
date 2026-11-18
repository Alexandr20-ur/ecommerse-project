import RegisterView from '@/views/RegisterView.vue'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import AdminLayout from "@/views/administration/latouts/AdminLayout.vue";
import DashboardView from "@/views/administration/DashboardView.vue";
import ProductsView from "@/views/administration/ProductsView.vue";

export const routes = [
    {
        path: '/',
        name: 'home',
        component: HomeView
    },
    {
        path: '/login',
        name: 'login',
        component: LoginView
    },
    {
        path: '/register',
        name: 'register',
        component: RegisterView
    },

    //===========================
    // ГРУППА: Админка
    //===========================
    {
        path: '/administration',
        component: AdminLayout,
        meta: { requiresAuth: true, requiresAdmin: true},
        children: [
            {
                path: '',
                redirect: '/administration/dashboard'
            },
            {
                path: 'dashboard',
                name: 'admin.Dashboard',
                component: DashboardView
            },
            {
                path: 'products',
                name: 'admin.Products',
                component: ProductsView
            }
        ]
    }
]