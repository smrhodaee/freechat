<template>
    <div ref="root" class="chat-list">
        <ChatItem v-for="(item, i) in items" v-model="items[i]" @delete="id => $emit('delete', id)"
            :isFirst="i == 0 ? true : items[i - 1].username != item.username" />
    </div>
</template>

<script setup lang="ts">
import Chat from '@/models/Chat';
import ChatItem from './ChatItem.vue';
defineProps<{
    items: Array<Chat>
}>()
defineEmits(['delete'])
</script>

<style scoped lang="scss">
.chat-list {
    width: 100%;
    display: flex;
    flex-direction: column;
}
</style>