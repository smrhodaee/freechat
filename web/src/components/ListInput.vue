<script setup lang="ts">
// import TextInput from './TextInput.vue';
import PrimaryButton from './PrimaryButton.vue';
import { ref } from 'vue';
import SelectBox from './ui/selectbox/SelectBox.vue';
import Option from './ui/selectbox/models';

 defineProps<{
    label: string,
    options?: Option[]
    min: number
}>()

defineEmits(["change"])

const items = defineModel<string[]>({ required: true })
const item = ref<Option>()

function addItem() {
    if (item.value) {
        if (!items.value.includes(item.value.value))
            items.value.push(item.value.value)
    }
}

function removeItem(value: string) {
    const i = items.value.indexOf(value)
    items.value.splice(i, 1)
}
</script>

<template>
    <div class="list-input">
        <form @submit.prevent="addItem" class="top">
            <SelectBox :min="min" class="flex-grow-1 me-1" @text-change="(v, callback) => $emit('change', v, callback)" :label="label" v-model="item"
                :options="options" :editable="true" />
            <PrimaryButton label="ADD" />
        </form>
        <ul>
            <li v-for="i in items">
                <span>{{ i }}</span>
                <i class="fa fa-trash" @click="removeItem(i)"></i>
            </li>
        </ul>
    </div>
</template>

<style scoped>
.list-input .top {
    display: flex;
    margin-bottom: 1rem;
}

.list-input ul {
    list-style: none;
    background-color: var(--dark);
    border-radius: 10px;
}

.list-input ul li {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    padding: 1rem;
}

.list-input ul li i {
    cursor: pointer;
}

.list-input ul li i:hover {
    color: var(--ink-medium);
}
</style>
