import { defineStore } from "pinia";
import { computed, ref } from "vue";
import ToastModel from "../models/ToastModel";


export const useToastStore = defineStore('toast', () => {
    const toastStr = localStorage.getItem("toast")
    const toast = ref<ToastModel | null>(toastStr != null ? JSON.parse(toastStr) : null)

    function set(title: string, value: string, type: string) {
        toast.value = {
            message: value,
            type: type,
            title: title
        }
        localStorage.setItem("toast", JSON.stringify(toast.value))
    }

    function setError(title: string, value: string) {
        set(title, value, "ERROR")
    }

    function setSuccess(title: string, value: string) {
        set(title, value, "SUCCESS")
    }

    function unset() {
        toast.value = null
        localStorage.removeItem("toast")
    }

    const title = computed(() => toast.value?.title)
    const type = computed(() => toast.value?.type)
    const message = computed(() => toast.value?.message)
    const show = computed(() => toast.value != null)

    return { show, type, title, message, setError, setSuccess, unset }
})