<script setup lang="ts">
import Chat from '@/models/Chat';
import { APIService } from '@/services/ApiService';
import { onMounted } from 'vue';
defineEmits(['delete'])
defineProps<{
    isFirst: boolean
}>()
const message = defineModel<Chat>({ required: true })

async function loadImage() {
    if (message.value.type == "IMAGE") {
        const res = await APIService.getImageMessage(message.value.id);
        if (!res.error) {
            if (res.data) message.value.url = res.data.url
        }
    }
}

onMounted(loadImage)

</script>

<template>
    <div class="chat-item" :class="{ 'owner': message.isOwner }">
        <div class="username" :style="{ 'display': !message.isOwner && isFirst ? 'block' : 'none' }">
            {{ message.username }}
        </div>
        <div class="value">
            <div class="body">
                <div v-if="message.type == 'IMAGE'" class="image">
                    <img v-if="message.url" :src="message.url" alt="chat image">
                    <div v-else class="loading"></div>
                </div>
                <div class="bottom">
                    <i v-if="message.isOwner" class="fa fa-trash" @click="$emit('delete', message.id)"></i>
                    <div class="text" v-html="message.text"></div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
.chat-item {
    display: flex;
    flex-direction: column;
    padding: 0.5rem;
    border-radius: 10px;
    width: max-content;
    margin-bottom: 1rem;
    background-color: var(--dark-alt);
    max-width: 90%;

    .username {
        margin-bottom: 0.5rem;
    }

    .value {
        display: flex;
        justify-content: center;
        align-items: center;

        .body {
            display: flex;
            flex-direction: column;

            .image {
                display: flex;
                justify-content: center;
                align-items: center;

                img {
                    border-radius: 10px;
                    width: 100%;
                    text-align: center;
                }

                .loading {
                    border: 2px solid white;
                    border-top: 2px solid var(--primary-color);
                    border-radius: 50%;
                    width: 2rem;
                    height: 2rem;
                    animation: spin 2s linear infinite;
                }

                @keyframes spin {
                    0% {
                        transform: rotate(0deg);
                    }

                    100% {
                        transform: rotate(360deg);
                    }
                }
            }


            .bottom {
                widows: 100%;
                display: flex;
                align-items: center;
                padding: 0.5rem;
                justify-content: space-between;

                i {
                    margin-right: 1rem;
                    cursor: pointer;
                }
                .text {
                    word-break: break-all;
                }
            }

        }
    }
}

.chat-item .value .chat-item .value .chat-item .username {
    background-color: var(--dark);
    border-radius: 10px;
    width: min-content;
    padding: 5px 10px;
}

.chat-item.owner {
    background-color: var(--success-color);
    align-self: end;
    color: white;
}
</style>
