<template>
    <textarea ref="mirror" class="mirror"></textarea>
    <div class="chat-input">
        <span class="top" v-if="file">
            <span>{{ file.name }}</span>
            <i class="fa fa-close" aria-hidden="true" @click="closeFile"></i>
        </span>
        <div class="middle">
            <textarea placeholder="Please enter your message" ref="textarea" class="text" v-model="message"></textarea>
            <div class="ctrl">
                <label for="file-upload">
                    <input id="file-upload" type="file" style="display: none;" @change="fileUpload"
                        accept="image/png,image/jpeg">
                    <i class="fa fa-file" aria-hidden="true"></i>
                </label>
                <i class="fa fa-chevron-right" aria-hidden="true" @click="() => submit()"></i>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
const textarea = ref<HTMLTextAreaElement>()
const mirror = ref<HTMLTextAreaElement>()
let borderTopWidth = 0
let borderBottomWidth = 0
const message = defineModel<string>({ required: true })
const emit = defineEmits(['submit'])
const parseValue = (v: string) => v.endsWith('px') ? parseInt(v.slice(0, -2), 10) : 0;
const file = ref<File>()

function closeFile() {
    file.value = undefined
}

onMounted(() => {
    if (textarea.value) {
        const textareaStyles = window.getComputedStyle(textarea.value);
        [
            'border',
            'width',
            'height',
            'boxSizing',
            'fontFamily',
            'fontSize',
            'fontWeight',
            'letterSpacing',
            'lineHeight',
            'padding',
            'textDecoration',
            'textIndent',
            'textTransform',
            'whiteSpace',
            'wordSpacing',
            'wordWrap',
        ].forEach((property: string) => {
            mirror.value?.style.setProperty(property, textareaStyles.getPropertyValue(property))
        });
        borderTopWidth = parseValue(textareaStyles.borderTopWidth);
        borderBottomWidth = parseValue(textareaStyles.borderBottomWidth);
        textarea.value.addEventListener('input', () => {
            if (textarea.value && mirror.value) {
                mirror.value.textContent = textarea.value.value;
                const newHeight = mirror.value.scrollHeight + borderTopWidth + borderBottomWidth;
                textarea.value.style.height = `${newHeight}px`;
            }
        })
        textarea.value.addEventListener('keydown', e => {
            if (e.key === "Enter") {
                e.preventDefault()
                submit();
            }
        })
    }
})

function submit() {
    emit('submit', file.value)
    if (textarea.value) textarea.value.style.height = "3rem";
    file.value = undefined
}

function fileUpload(e: Event) {
    const elm = (e.target as HTMLInputElement)
    if (elm.files) file.value = elm.files[0]
}
</script>

<style scoped>
.chat-input {
    width: calc(100% - 22rem);
    display: flex;
    flex-direction: column;
    background-color: var(--dark);
    margin: 1rem;
    border-radius: 10px;
    overflow: hidden;


    .top {
        align-self: flex-end;
        display: flex;
        justify-content: center;
        align-items: center;

        i {
            cursor: pointer;
            padding: 1rem;
        }
    }

    .middle {
        display: flex;
        width: 100%;
        justify-content: center;
        align-items: center;

        .ctrl {
            display: flex;

            i {
                display: block;
                padding: 1rem;
                cursor: pointer;
            }
        }

        .text {
            position: relative;
            flex-grow: 1;
            max-height: 320px;
            padding: 0.5rem;
            height: 3rem;
            resize: none;
            outline: 0;
            transition: height 0.2s ease-in-out;
        }
    }


}

.mirror {
    position: fixed;
    top: -9999px;
    left: -9999px;
    visibility: hidden;
}

@media screen and (max-width: 992px) {
    .chat-input {
        width: calc(100% - 2rem);
    }
}
</style>