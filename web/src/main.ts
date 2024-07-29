import vue from 'vue'
import { createApp } from 'vue'
import { RouteRecordRaw, createMemoryHistory, createRouter } from 'vue-router'
import { useAuthStore } from './stores/auth'
import './style.css'
import App from './App.vue'
import { createPinia } from 'pinia'

import HomeView from './views/HomeView.vue'
import LoginView from './views/LoginView.vue'
import CreateGroup from './views/CreateGroup.vue'


const routes: RouteRecordRaw[] = [
    { path: "/", name: "home", component: HomeView, meta: { requireAuth: true } },
    { path: "/login", name: "login", component: LoginView },
    { path: "/create-group", name: "create-group", component: CreateGroup, meta: { requireAuth: true } }
]

const router = createRouter({
    history: createMemoryHistory(),
    routes
})

router.beforeEach((to) => {
    const authStore = useAuthStore()
    if (authStore.isLoggedIn) {
        if (to.name == "login") return '/'
    } else {
        if (to.meta.requireAuth) {
            return '/login'
        }
    }
})

const helpers: vue.Plugin = {
    install(app: vue.App, options: any) {
    }
}

createApp(App)
    .use(createPinia())
    .use(helpers)
    .use(router)
    .mount('#app')
