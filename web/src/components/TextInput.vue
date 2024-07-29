<template>
    <div class="input">
        <label :class="['filled', disabled ? 'disable' : '']">
            <textarea v-if="multiLine" :placeholder="label" v-model="value" :disabled="disabled" required></textarea>
            <input v-if="!multiLine" required :type="type ? type : 'text'" v-model="value" :disabled="disabled">
            <span v-if="!multiLine" class="label">{{ label }}</span>
            <!-- <span class="helper">Helper Text</span> -->
            <!-- <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
                <path fill="none" d="M0 0h24v24H0V0z" />
                <circle cx="15.5" cy="9.5" r="1.5" />
                <circle cx="8.5" cy="9.5" r="1.5" />
                <path
                    d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zM12 20c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8zm-5-6c.78 2.34 2.72 4 5 4s4.22-1.66 5-4H7z" />
            </svg> -->
            <Spiner class="spiner" v-if="loading"/>
        </label>
    </div>
</template>

<script setup lang="ts">
import Spiner from '@/components/ui/Spiner.vue';

defineProps<{
    label: string,
    type?: string,
    disabled?: boolean
    multiLine?: boolean,
    loading?: boolean
}>()
const value = defineModel<string>({ required: true })
</script>

<style lang="scss" scoped>
.input {
    position: relative;

    .spiner {
        position: absolute;
        width: 1.25rem;
        height: 1.25rem;
        top: 5px;
        transform: translateY(-100%);
        right: 5px;
    }
}

.input textarea {
    border-radius: 10px;
    width: 100%;
    resize: none;
    height: 3.5rem;
    min-height: 3.5rem;
    max-height: 50rem;
    padding: 0.9rem;
    font-size: 1rem;
}

.input .filled>input {
    border: none;
    border-radius: 10px;
    border-bottom: 0.125rem solid var(--dark);
    width: 100%;
    font-size: 1rem;
    padding-left: 0.9rem;
    line-height: 147.6%;
    padding-top: 1rem;
    padding-bottom: 0.5rem;
    height: 4rem;
    color: white;
    background: var(--dark);
}

.input .label {
    cursor: text;
    color: #eee;
}

.filled.disable,
.filled.disable>.label {
    cursor: default;
}


.input .filled>input:focus {
    outline: none;
}

.filled:not(.danger)>input:focus {
    background: var(--dark);
    border-color: var(--primary-color);
}

.filled:not(.danger)>input:valid {
    background: var(--dark);
}

.filled:not(.disable)>input:hover {
    background: var(--dark-alt);
}

.filled>input:focus+.label,
.filled>input:valid+.label {
    top: .25rem;
    font-size: 0.9375rem;
    margin-bottom: 1rem;
}

.filled:not(.danger)>input:focus+.label {
    color: var(--primary-color);
}

.filled:not(.danger)>input:focus~svg {
    fill: var(--primary-color);
}

.filled>.label {
    position: absolute;
    font-size: 0.95rem;
    top: 1.375rem;
    left: 0.9rem;
    line-height: 147.6%;
    color: var(--ink-medium);
    transition: top .2s;
}

.filled>svg {
    position: absolute;
    font-size: 0.95rem;
    top: 1.375rem;
    right: 0.875rem;
    fill: var(--ink-medium);
}

.filled>.helper {
    font-size: 0.9375rem;
    color: white;
    left: 0;
    letter-spacing: 0.0275rem;
    margin: 0.125rem 0.875rem;
}

.filled.danger>.label,
.filled.danger>.helper {
    color: var(--color-danger);
}

.danger>svg {
    fill: var(--color-danger);
}

.danger>input {
    border-color: var(--color-danger);
}
</style>
