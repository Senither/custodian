<script setup lang="ts">
const route = useRoute()

watch(route, () => {
    if (typeof themeChange === 'function') {
        themeChange(false)
    }
})
</script>

<template>
    <NuxtLayout>
        <NuxtPage />
    </NuxtLayout>
</template>
