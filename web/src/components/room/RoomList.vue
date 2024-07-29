<template>
    <div class="room-list">
        <RoomItem v-for="item in items" :item="item" :selected="item.name == selected.name" @select="select" />
    </div>
</template>

<script setup lang="ts">
import Room from '@/models/Room';
import RoomItem from './RoomItem.vue';
defineProps<{
    items: Array<Room>,
}>()
const selected = defineModel<Room>({ required: true })
const emit = defineEmits(["select"])

function select(room: Room) {
    if(room.name != selected.value.name) {
        emit('select', room)
        selected.value = room
    }
}
</script>

<style scoped></style>