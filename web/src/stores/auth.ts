import { defineStore } from "pinia";
import { computed, ref } from "vue";


export const useAuthStore = defineStore('auth', () => {
    const token = ref(localStorage.getItem("token"))
    const username = ref(localStorage.getItem("username"))

    const isLoggedIn = computed(() => token.value != null && token.value.length == 64 && username.value != null)

    function login(t: string, id: string) {
        token.value = t
        username.value = id
        localStorage.setItem("token", t)
        localStorage.setItem("username", id)
    }

    function logout() {
        token.value = null
        username.value = null
        localStorage.removeItem("token")
        localStorage.removeItem("username")
    }

    return { token, username, isLoggedIn, login, logout }
})