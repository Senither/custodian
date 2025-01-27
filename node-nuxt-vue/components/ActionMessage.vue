<script setup lang="ts">
const visible = ref(false)

function show() {
    visible.value = true

    setTimeout(() => {
        visible.value = false
    }, 2500)
}

defineExpose({ show })
</script>

<template>
    <transition name="fade">
        <div v-if="visible" class="mr-4 text-secondary text-sm">
            <slot />
        </div>
    </transition>
</template>

<style scoped>
.fade-enter-active {
    transition: opacity 0.1s ease;
}

.fade-leave-active {
    transition: opacity 0.75s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}
</style>
