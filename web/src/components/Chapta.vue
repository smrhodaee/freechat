<script setup lang="ts">
import { APIService } from "../services/ApiService";
import { ref, onMounted, defineModel } from "vue";

const image = ref<HTMLImageElement>()
const error = ref(false)
const emit = defineEmits(["error"])

const uuid = defineModel()

const fetchImage = async () => {
  let res = await APIService.getChapta()
  if (res.error) {
    error.value = true
    emit('error', res.message)
  } else {
    if (res.data) {
      uuid.value = res.data.uuid
      if (image.value) {
        image.value.src = res.data.url
        error.value = false
      }
    }
  }
}

onMounted(fetchImage)

</script>

<template>
  <div class="chapta">
    <img v-show="!error" ref="image">
    <span v-show="error" class="error">
      <i class="fa fa-exclamation-triangle "></i>
    </span>
    <div class="refresh" @click="fetchImage">
      <i class="fa fa-refresh" aria-hidden="true"></i>
    </div>
  </div>
</template>

<style scoped>
.chapta {
  width: 300px;
  height: 4rem;
  display: flex;
  justify-content: end;
  border-radius: 7px;
  overflow: hidden;
  background-color: var(--dark);
}

.refresh {
  width: 60px;
  font-size: 1.5rem;
  cursor: pointer;
  background-color: var(--primary-color);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center
}

.refresh:hover {
  background-color: var(--primary-alpha-color);
}

.error {
  display: flex;
  flex-grow: 1;
  justify-content: center;
  align-items: center;
  font-size: 2.5rem;
  text-align: center;
  color: white;
}
</style>
