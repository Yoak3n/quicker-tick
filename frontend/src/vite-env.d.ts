/// <reference types="vite/client" />

declare module '*.vue' {
    import type {DefineComponent} from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component
}
// import type {LoadingBarApi,DialogApiInjection,MessageApiInjection,NotificationApiInjection,ModalApiInjection} from "naive-ui"

declare interface Window{
    $loadingBar:LoadingBarInst
    $dialog:DialogApiInjection
    $message:MessageApiInjection
    $notification:NotificationApiInjection
    $modal:ModalApiInjection
} 