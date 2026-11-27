interface ImportMetaEnv {
    readonly VITE_API_URL: string
    readonly BASE_URL: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}

declare module '*.css' {
    const content: Record<string, string>
    export default content
}

// Объявляем, что все .vue файлы являются валидными модулями
declare module '*.vue' {
    import type {DefineComponent} from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component
}
