<script setup lang="ts">
const runtime = useRuntimeConfig()
</script>

<template>
    <div class="navbar-start">
        <NuxtLink href="/">
            <h1 class="font-extrabold text-4xl text-primary">{{ runtime.public.appName }}</h1>
            <p class="-mt-1.5 ml-6 font-semibold text-lg text-secondary">{{ runtime.public.appDescriptor }}</p>
        </NuxtLink>
    </div>
</template>
