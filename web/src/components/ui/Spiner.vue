<template>
    <div class="spiner">
    </div>
</template>

<script setup lang="ts">

</script>

<style lang="scss" scoped>
.spiner {
    border-radius: 100%;
    box-sizing: border-box;
    border: 2px solid #FFF;
    border-bottom-color: transparent;
    animation: spin infinite 1s linear;
}

@keyframes spin {
    0% {
        transform: rotate(0deg);
    }

    100% {
        transform: rotate(360deg);
    }
}
</style>