/**
 * Базовый тип товара, совпадает с таблицей Products
 */
export interface Product {
    id: string;
    name: string;
    description: string | null;
    price: number;
    stock: number;
    createdAt: string;
    updatedAt: string;
}

/**
 * Тип создания товара
 */
export interface CreateProductRequest {
    name: string;
    description?: string;
    price: number;
    stock: number;
}

/**
 * Тип обновления товара
 */
export type UpdateProductRequest = Partial<CreateProductRequest>;

/**
 * Форма по умолчанию
 */
export const createEmptyProductForm = (): CreateProductRequest => ({
    name: '',
    description: '',
    price: 0,
    stock: 0,
})

/**
 * Параметры запросов, для получения списка товаров
 */
export interface ProductListParams {
    page?: number;
    limit?: number;
    search?: string;
    minPrice?: number;
    maxPrice?: number;
    inStock?: number;
}

/**
 * Ответ сервера при запросах GET
 */
export interface ProductListResponse {
    data: Product[];
    total: number;
    page: number;
    limit: number;
}