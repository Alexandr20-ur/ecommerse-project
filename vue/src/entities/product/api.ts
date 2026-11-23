import apiClient from '@/utils/axios.ts'
import type {
    CreateProductRequest,
    UpdateProductRequest,
    ProductListResponse,
    ProductListParams
} from './model'

export const productService = {
    // POST /products
    create: (data: CreateProductRequest) => apiClient.post('/products', data),

    // PUT /products/{id}
    update: (id: number | string, data: UpdateProductRequest) => apiClient.put(`/products/${id}`, data),

    // GET /products
    getList: (params: ProductListParams) => apiClient.get<ProductListResponse>('/products', { params }),

    // Бонус: если у вас есть ShowHandler
    getById: (id: number | string) => apiClient.get(`/products/${id}`),

    // Бонус: если у вас есть DeleteHandler
    delete: (id: number | string) => apiClient.delete(`/products/${id}`)
}