<script setup lang="ts">
useHead({
    title: 'Custodian',
    meta: [
        {
            name: 'description',
            content: 'NodeJS + Nuxt + Vue version of Custodian',
        },
    ],
})
</script>

<template>
    <header class="top-0 fixed flex justify-between items-center px-6 py-4 w-full">
        <AppHeader />
        <ThemeSelector />
    </header>

    <main class="bg-base-200 min-h-screen hero">
        <div class="flex-col hero-content">
            <slot />
        </div>
    </main>

    <footer
        class="flex justify-center items-center bg-base-100 -mt-[3.54rem] px-4 py-3 border-t border-base-300 w-full">
        <AppFooter />
    </footer>
</template>
