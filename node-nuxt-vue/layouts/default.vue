<script setup lang="ts">
useHead({
    title: 'Custodian',
    meta: [
        {
            name: 'description',
            content: 'NodeJS + Nuxt + Vue version of Custodian',
        },
    ],
    bodyAttrs: {
        class: 'flex flex-col justify-between bg-base-200 min-h-screen',
    },
})
</script>

<template>
    <header class="flex justify-between items-center px-4 py-2">
        <div class="mx-auto max-w-5xl navbar">
            <AppHeader />

            <div class="flex items-center gap-4 navbar-end">
                <ThemeSelector />

                <div class="dropdown dropdown-end">
                    <div tabindex="0" role="button" class="bg-base-100 m-1 btn">
                        <svg
                            class="w-5 h-5" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                            viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true"
                        >
                            <path
                                stroke-linecap="round" stroke-linejoin="round"
                                d="M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z"
                            />
                        </svg>
                    </div>
                    <ul
                        tabindex="0"
                        class="z-[1] bg-base-100 shadow p-2 border border-base-200 rounded-box w-44 dropdown-content menu"
                    >
                        <li>
                            <NuxtLink href="/profile">
                                Profile
                            </NuxtLink>
                        </li>
                        <li>
                            <NuxtLink href="/logout">
                                Logout
                            </NuxtLink>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </header>

    <main class="flex flex-col gap-4 mx-auto mt-8 md:mt-12 mb-auto max-w-3xl container">
        <slot />
    </main>

    <footer class="items-center bg-base-100 mt-12 p-4 border-t border-base-300 text-base-content footer">
        <AppFooter />
    </footer>
</template>

<style>
#custodian {
    @apply flex flex-col flex-auto;
}
</style>
