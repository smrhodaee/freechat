<script setup lang="ts">
import ListInput from '@/components/ListInput.vue';
import PrimaryButton from '@/components/PrimaryButton.vue';
import TextInput from '@/components/TextInput.vue';
import Option from '@/components/ui/selectbox/models';
import { Color } from '@/models/Color';
import { APIService } from '@/services/ApiService';
import { useToastStore } from '@/stores/toast';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const title = ref("")
const roomName = ref("")
const usernames = ref<string[]>([])
const options = ref<Option[]>()
const toastStore = useToastStore()
const router = useRouter()

async function addRoom() {
    const res = await APIService.createGroup({
        name: roomName.value,
        title: title.value,
        usernames: usernames.value,
    })
    if (res.error) {
        toastStore.setError("Create Group", res.message)
    } else {
        toastStore.setSuccess("Create Group", res.message)
        router.push({ name: 'home' })
    }
}


const listInputChange = async (v: string, callback: Function) => {
    const res = await APIService.findUsers(v)
    if (res.error) {
        toastStore.setError("Create Group", res.message)
    } else {
        options.value = res.data?.map<Option>((u) => {
            return {
                "value": u.username,
                "label": u.username,
            }
        })
    }
    callback()
}
</script>

<template>
    <div class="add-room">
        <form class="mb-1" @submit.prevent="addRoom">
            <TextInput label="Name" v-model="roomName" />
        </form>
        <div class="middle mb-1">
            <TextInput label="Title" class="flex-grow-1 me-1" v-model="title" />
            <PrimaryButton class="me-1" @click="addRoom" :color="Color.Success" label="Create Group" />
            <RouterLink class="cancel" :to="{ name: 'home' }">Cancel</RouterLink>
        </div>
        <ListInput :min="3" :options="options" class="mb-1" label="Username" v-model="usernames"
            @change="listInputChange" />
    </div>
</template>

<style scoped>
.add-room {
    padding: 0 1rem;
    margin-top: 2rem;
}

.add-room form {
    display: flex;
    flex-direction: column;
}

.add-room .middle {
    display: flex;
    flex-direction: row;
}

.cancel {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1rem 2rem;
    background-color: var(--dark);
    border-radius: 10px;
    color: white;
    font-size: 15px;
    font-weight: 700;
}

.cancel:hover {
    background-color: var(--dark-alt);
}
</style>
