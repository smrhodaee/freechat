<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router';
import { useAuthStore } from './stores/auth';
import Toast from './components/Toast.vue'
import { storeToRefs } from 'pinia';
import {  watch } from 'vue';

const authStore = useAuthStore()
const router = useRouter()
const { isLoggedIn } = storeToRefs(authStore)

watch(isLoggedIn, (v) => {
  if (!v) router.push({ name: "login" })
})

</script>

<template>
  <main>
    <RouterView />
  </main>
  <Toast />
</template>

<style scoped>
header {
  margin-bottom: 1rem;
  padding: 0.5rem;
  width: 100%;
  display: flex;
  justify-content: end;
  align-items: center;
}

header .brand {
  flex-grow: 1;
  padding: 0.25rem;
  margin-right: 4.5rem;
}

.menu {
  padding: 1rem;
  color: var(--ink-medium);
  cursor: pointer;
  font-size: 2rem;
}

.menu i {
  display: block;
}

.menu:hover {
  color: white;
}
</style>
