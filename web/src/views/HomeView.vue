<template>
    <div class="home">
        <nav ref="nav" :class="{ 'active': showMenu }">
            <div class="top">
                <div class="brand">
                    <img class="logo"src="@/assets/brand-logo.png" alt="brand_logo">
                    <span class="title">
                        Free Chat
                    </span>
                </div>
                <i class="fa fa-close" @click="showMenu = false"></i>
            </div>
            <div class="links">
                <RouterLink class="mb-1" :to="{ name: 'create-group' }">Create Group</RouterLink>
                <a class="logout" @click="logout">Logout</a>
            </div>
            <p><strong>Rooms:</strong></p>
            <RoomList :items="rooms" v-model="selectedRoom" @select="selectRoom" />
        </nav>
        <main>
            <div class="top">
                <i ref="openNav" class="fa fa-bars" :class="{ 'active': !showMenu }" @click="openSidebar"></i>
                <span v-if="selectedRoom" class="title">{{ selectedRoom.title }}</span>
            </div>
            <div ref="middle" class="middle" v-if="selectedRoom.name">
                <ChatList ref="chatList" :items="chats" @delete="deleteChat" />
            </div>
            <p class="nowrap" v-else>
                Please select room to start
            </p>
            <div class="bottom" v-if="selectedRoom.name">
                <ChatInput v-model="message" @submit="createChat" />
            </div>
        </main>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { APIService } from '@/services/ApiService';
import ChatInput from '@/components/chat/ChatInput.vue';
import ChatList from '@/components/chat/ChatList.vue';
import { useAuthStore } from '@/stores/auth';
import { useToastStore } from '@/stores/toast';
import Chat from '@/models/Chat';
import { GlobalResponse } from '@/services/common';
import { CreateMessageResponse, WsChat } from '@/services/WsService';
import RoomList from '@/components/room/RoomList.vue';
import Room from '@/models/Room';

const authStore = useAuthStore()
const toastStore = useToastStore()
const chats = ref<Array<Chat>>([])
const chatList = ref<InstanceType<typeof ChatList>>()
const message = ref("");
const showMenu = ref(true);
const nav = ref<HTMLElement>()
const middle = ref<HTMLElement>()

const rooms = ref<Array<Room>>([]);
const selectedRoom = ref<Room>(new Room())

async function selectRoom(room: Room) {
    chats.value = [];
    const res = await APIService.findChatsOfRoom(room.name)
    handleError("Select Room", res)
    if (!res.error) {
        if (res.data) {
            for (let i = 0; i < res.data.length; i++) addChat(res.data[i]);
            selectedRoom.value.unreadMessages = 0;
        }
    } else {
        selectedRoom.value = new Room();
    }
    showMenu.value = false
}

async function createChat(file?: File) {
    handleError("Chat Submit", await APIService.submitChat({
        "room_name": selectedRoom.value?.name || "",
        "text": message.value
    }, file))
    message.value = ""
}

function deleteChat(id: string) {
    APIService.deleteChat(id)
}

function handleError<T>(label: string, res: GlobalResponse<T>) {
    if (res.error) {
        toastStore.setError(label, res.message)
        if (res.status == 401) authStore.logout()
    }
}

async function addChat(data: any) {
    chats.value.push(
        new Chat(
            data.id,
            data.username,
            data.text.replaceAll("\n", "<br>"),
            data.value_id,
            data.type,
            (data.username == authStore.username),
        ))
    setTimeout(scrollDown, 300)
}

const scrollDown = () => {
    middle.value?.scrollTo({
        behavior: "smooth",
        top: middle.value.scrollHeight + 100,
    })
}

const onChatMessage = (msg: CreateMessageResponse) => {
    if (msg.room_name == selectedRoom.value.name) {
        addChat(msg)
    } else {
        rooms.value.forEach((v) => {
            if (v.name == msg.room_name) {
                v.unreadMessages++;
            }
        })
    }
}

function onChatDelete(id: string) {
    const index = chats.value.findIndex((el) => el.id == id)
    if (index != -1)
        chats.value.splice(index, 1)
    else console.error("chat not found")
}

const onChatError = (msg: string) => {
    toastStore.setError("Chat Message", msg)
}

onMounted(async () => {
    const res = await APIService.findRoomsOfUser()
    handleError("Get Rooms", res)
    if (!res.error && res.data) {
        rooms.value = res.data.map((v) => new Room(v.name, v.title, v.type, 0))
    }
    new WsChat(onChatMessage, onChatDelete, onChatError)
})

async function logout() {
    const res = await APIService.logout()
    if (res.error) {
        toastStore.setError("Logout", res.message)
    } else {
        toastStore.setSuccess("Logout", res.message)
    }
    authStore.logout()
}
const openSidebar = () => showMenu.value = true;
</script>


<style scoped lang="scss">
.home {
    position: relative;

    .brand {
        width: 100%;
        display: flex;
        align-items: center;

        .logo {
            width: 4rem;
            height: 4rem;
            margin-right: 2rem;
        }
        .title {
            font-family: beautiful;
            font-size: 3rem;
        }
    }

    nav {
        width: 20rem;
        position: fixed;
        height: 100%;
        background-color: var(--dark);
        display: flex;
        flex-direction: column;

        .top {
            display: flex;
            padding: 1rem;
            justify-content: space-between;
            align-items: center;

            i {
                font-size: 2rem;
                cursor: pointer;
                display: none;
            }
        }

        .links {
            display: flex;
            flex-direction: column;
            padding: 0.5rem 1rem;
        }

        p {
            padding: 1rem;
        }
    }


    main {
        display: flex;
        margin-left: 20rem;
        flex-direction: column;
        flex-grow: 1;
        height: 100vh;
        position: relative;

        .top {
            display: flex;
            padding: 1rem;
            width: 100%;

            i {
                padding-right: 1rem;
                font-size: 2rem;
                cursor: pointer;
                display: none;
            }

            // i.active {}

            .title {
                font-size: 1.5rem;
            }
        }

        .middle {
            width: 100%;
            margin-top: 4rem;
            position: absolute;
            padding: 1rem;
            overflow: hidden;
            top: 0;
            overflow-y: auto;
            bottom: 5rem;
            flex-grow: 1;
        }

        .bottom {
            position: fixed;
            bottom: 0;
            width: 100%;
        }

        p {
            position: absolute;
            transform: translate(-50%, -50%);
            top: 50%;
            left: 50%;
        }
    }

    .logout {
        cursor: pointer;
    }


    @media screen and (max-width: 992px) {
        nav {
            width: 100%;
            position: fixed;
            z-index: 20;
            transform: translateX(-100%);
            transition: transform 0.5s ease-out;

            &.active {
                transform: translateX(0);
            }

            .top {
                i {
                    display: block;
                }
            }
        }

        main {
            margin-left: 0;

            .top {
                i {
                    display: block;
                }
            }
        }
    }
}

.nowrap {
    text-wrap: nowrap;
}
</style>
