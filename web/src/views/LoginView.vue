<template>
    <div class="login">
        <div class="wrapper">
            <form @submit="submit">
                <TextInput class="mb-1" label="Username" v-model="data.username" />
                <TextInput class="mb-1" label="Password" type="password" v-model="data.password" />
                <TextInput class="mb-1" label="Chapta" v-model="data.code" />
                <Chapta class="mb-1" v-model="data.uuid" @error="chaptaErr" />
                <PrimaryButton label="Login/Register" />
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import Chapta from '../components/Chapta.vue'
import PrimaryButton from '../components/PrimaryButton.vue'
import TextInput from '../components/TextInput.vue'
import { APIService, RegisterOrLoginRequest } from '@/services/ApiService';
import { useAuthStore } from '../stores/auth';
import { useRouter } from 'vue-router';
import { useToastStore } from '../stores/toast';

const data = ref<RegisterOrLoginRequest>({
    username: "",
    password: "",
    uuid: "",
    code: "",
})
const authStore = useAuthStore()
const toastStore = useToastStore()
const router = useRouter()

const submit = async (e: Event) => {
    e.preventDefault()
    let res = await APIService.registerOrLogin(data.value)
    if (res.error) {
        toastStore.setError("Register Or Login Error", res.message)
    } else {
        toastStore.setSuccess("Register Or Login", res.message)
        authStore.login(res.data?.token || "", data.value.username)
        router.push({ name: "home" })
    }
}

const chaptaErr = (err: string) => {
    toastStore.setError("Chapta Error", err)
}

</script>

<style scoped>
.wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
}

form {
    display: flex;
    flex-direction: column;
}
</style>
