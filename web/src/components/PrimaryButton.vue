<template>
    <button :class="['btn', colorClass, disabled ? 'disabled' : '']" :disabled="disabled">
        {{ label }}
    </button>
</template>

<script setup lang="ts">
import { Color } from '@/models/Color';
import { onMounted, ref } from 'vue';

const props = withDefaults(defineProps<{
    label: string
    color?: Color
    disabled?: boolean
}>(), {
    color: Color.Info
})
const colorClass = ref("")

onMounted(() => {
    switch (props.color) {
        case Color.Danger:
            colorClass.value = "danger"
            break;
        case Color.Success:
            colorClass.value = "success"
            break;
        default:
            colorClass.value = "info"
            break;
    }
})
</script>

<style scoped>
.btn {
    color: white;
    text-align: center;
    padding: 20px 35px;
    border-radius: 7px;
    cursor: pointer;
    font-size: 15px;
    font-weight: 700;
}

.btn.disabled {
    cursor: default;
    opacity: 0.5;
}

.btn.info {
    background-color: var(--primary-color);
}

.btn.success {
    background-color: var(--success-color);
}

.btn.danger {
    background-color: var(--danger-color);
}

.btn:not(.disabled):hover {
    /*background-color: var(--primary-alpha-color);*/
    opacity: 0.9;
}
</style>
