<template>
    <div :class="['room-item', selected ? 'active' : '']" @click="$emit('select', item)">
        <div class="type info-badge">{{  item.type }}</div>
        {{ item.title }}
        <span class="unread success-badge" v-if="item.unreadMessages">{{ item.unreadMessages }}</span>
    </div>
    <hr>
</template>

<script setup lang="ts">
import Room from '@/models/Room';
defineProps<{
    item: Room,
    selected: boolean
}>()
defineEmits(['select'])
</script>

<style scoped>
.room-item {
    padding: 1rem 2rem;
    text-align: center;
    width: 100%;
    cursor: pointer;
    position: relative;
}

.type {
    position: absolute;
    left: 1rem;
}

.unread {
    position: absolute;
    right: 1rem;
}

.room-item:hover {
    background-color: var(--dark-alt);
}

.room-item.active {
    background-color: var(--full-dark);
    cursor: default;
}
</style>