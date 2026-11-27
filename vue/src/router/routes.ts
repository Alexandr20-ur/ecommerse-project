import RegisterView from '@/views/RegisterView.vue'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import AdminLayout from "@/views/administration/layouts/AdminLayout.vue"; // Проверь опечатку в latouts -> layouts
import DashboardView from "@/views/administration/DashboardView.vue";
import ProductsView from "@/views/administration/module/products/ProductsView.vue";
import CreateProductsView from "@/views/administration/module/products/CreateProductsView.vue";

export const routes = [
    {
        path: '/',
        name: 'home',
        component: HomeView
    },
    {
        path: '/login',
        name: 'login',
        component: LoginView,
        meta: {isGuestOnly: true} // <-- Добавлено
    },
    {
        path: '/register',
        name: 'register',
        component: RegisterView,
        meta: {isGuestOnly: true} // <-- Добавлено
    },
    //===========================
    // ГРУППА: Админка
    //===========================
    {
        path: '/administration',
        component: AdminLayout,
        meta: {requiresAuth: true, requiresAdmin: true},
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
            },
            {
                path: 'products/create',
                name: 'admin.Products.Create',
                component: CreateProductsView
            }
        ]
    }
]