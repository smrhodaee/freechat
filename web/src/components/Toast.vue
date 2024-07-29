<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { useToastStore } from '../stores/toast';
import Timer from "../tools/Timer"
import { storeToRefs } from 'pinia';

const toast = ref<HTMLDivElement>()
const progress = ref<HTMLDivElement>()
const toastStore = useToastStore()
const { show } = storeToRefs(toastStore)

watch(show, (v) => {
    if (v) open()
})

let timer: Timer

onMounted(() => {
    timer = new Timer(5000, (ms) => {
        if (ms == 0) {
            toast.value?.classList.remove("active");
            setTimeout(() => {
                progress.value?.style.setProperty('--progress-right', "0")
                toastStore.unset()
            }, 500);
        } else {
            progress.value?.style.setProperty('--progress-right', `${100 - ((ms / 5000) * 100)}%`)
        }
    })
    toast.value?.addEventListener('mouseenter', () => timer.pause())
    toast.value?.addEventListener('mouseleave', () => timer.resume())
    if (show.value) open()
})

const open = () => {
    toast.value?.classList.add("active");
    timer.start()
}

</script>
<template>
    <div ref="toast" class="toast">
        <div class="toast-content">
            <i v-if="toastStore.type == 'ERROR'" class="fa fa-times icon error"></i>
            <i v-else class="fa fa-check icon success"></i>
            <div class="message">
                <span class="text title">{{ toastStore.title }}</span>
                <span class="text caption">{{ toastStore.message }}</span>
            </div>
        </div>
        <i class="fa fa-times close" aria-hidden="true" @click="timer.stop"></i>
        <div v-if="toastStore.type == 'ERROR'" ref="progress" class="progress error"></div>
        <div v-else ref="progress" class="progress success"></div>
    </div>
</template>

<style scoped>
.toast {
    position: fixed;
    top: 1.25rem;
    right: 2rem;
    border-radius: 10px;
    background: var(--dark);
    padding: 1rem;
    box-shadow: 0 6px 20px -5px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    transform: translateX(calc(100% + 2rem));
    transition: all 0.5s cubic-bezier(0.68, -0.55, 0.265, 1.35);
    z-index: 100;
}

.toast.active {
    transform: translateX(0%);
}

.toast .toast-content {
    display: flex;
    align-items: center;
}

.toast-content .icon {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 2.1875rem;
    min-width: 2.1875rem;
    color: #fff;
    text-align: center;
    font-size: 1.25rem;
    border-radius: 50%;
}

.toast-content .icon.info {
    background-color: #4070f4;
}

.toast-content .icon.error {
    background-color: var(--danger-color);
}

.toast-content .icon.success {
    background-color: var(--success-color);
}

.toast-content .message {
    display: flex;
    flex-direction: column;
    margin: 0 2rem;
    /* padding: 10px 15px; */
}


.message .text {
    font-size: 1rem;
    font-weight: 400;
    color: var(--ink-medium);
}

.message .text.title {
    font-weight: 600;
    text-align: left;
    color: #fff;
}

.toast .close {
    position: absolute;
    top: 10px;
    right: 15px;
    padding: 5px;
    cursor: pointer;
    opacity: 0.7;
    color: #fff;
}

.toast .close:hover {
    opacity: 1;
}

.toast .progress {
    position: absolute;
    bottom: 0;
    left: 0;
    height: 3px;
    width: 100%;
}

.toast .progress:before {
    content: "";
    position: absolute;
    bottom: 0;
    right: var(--progress-right);
    height: 100%;
    width: 100%;
}

.toast .progress.info:before {
    background-color: #4070f4;
}

.toast .progress.error:before {
    background-color: var(--danger-color);
}

.toast .progress.success:before {
    background-color: var(--success-color);
}

.toast.active~button {
    pointer-events: none;
}
</style>