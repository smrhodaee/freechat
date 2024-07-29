<template>
    <div ref="root" class="select-box">
        <div v-if="!editable" class="combobox" @click="isWrapperVisiable = !isWrapperVisiable">
            <div class="selected-option">
                {{ mappedSelectedOption }}
            </div>
            <i v-show="isWrapperVisiable" class="fa fa-chevron-up"></i>
            <i v-show="!isWrapperVisiable" class="fa fa-chevron-down"></i>
        </div>
        <div v-else class="textbox">
            <TextInput :label="label" v-model="text" :loading="loading" />
        </div>
        <Transition name="slide-fade">
            <div v-if="isWrapperVisiable" class="options-wrapper">
                <div v-for="option in options" class="option" @click="selectOption(option)">
                    {{ option.label }}
                </div>
            </div>
        </Transition>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue';
import TextInput from '@/components/TextInput.vue';
import Option from "./models"
import { watch } from 'vue';

const props = defineProps<{
    label: string,
    options?: Option[],
    editable: boolean,
    min: number
}>()

const emit = defineEmits(["textChange"])

const selectedOption = defineModel<Option>()
const root = ref<HTMLElement>()
const text = ref("")
const loading = ref(false)
const isWrapperVisiable = ref(false)
const mappedSelectedOption = computed(() => {
    return selectedOption.value?.label || "Please select something"
})

watch(text, (v) => {
    isWrapperVisiable.value = false
    let curr = null
    props.options?.forEach((opt) => {
        if(opt.value == v) {
            curr = opt
            return
        }
    })
    if (curr) {
        selectedOption.value = curr
    } else if (v.length >= props.min) {
        loading.value = true
        emit("textChange", v, () => {
            loading.value = false
            isWrapperVisiable.value = true
        })
    }
})

const selectOption = (option: Option) => {
    text.value = option.value
}

const closeWrapper = (element: Event) => {
    if (!root.value?.contains(element.target as Node)) {
        isWrapperVisiable.value = false
    }
}

onMounted(() => {
    window.addEventListener('click', closeWrapper)
})

onBeforeUnmount(() => {
    window.removeEventListener('click', closeWrapper)
})
</script>

<style lang="scss" scoped>
.select-box {
    position: relative;

    .combobox {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 1.5rem;
        background-color: var(--dark);
        border-radius: 10px;
        height: 4rem;
        cursor: pointer;
    }

    .selected-option {
        margin-right: 1rem;
    }

    .options-wrapper {
        width: 100%;
        max-height: 20rem;
        overflow-y: auto;
        display: flex;
        flex-direction: column;
        position: absolute;
        top: 100%;
        left: 0;
        border-radius: 0 0 10px 10px;
        background-color: var(--dark-alt);

        .option {
            text-align: center;
            cursor: pointer;
            padding: 0.5rem;
        }

        .slide-fade-enter-active {
            transition: all 0.3 ease-out;
        }


        .slide-fade-enter-from,
        .slide-fade-leave-to {
            height: 0;
            opacity: 0;
        }
    }
}
</style>