import process from 'node:process'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    compatibilityDate: '2024-11-01',
    devtools: { enabled: true },
    modules: ['@prisma/nuxt', '@nuxtjs/tailwindcss', 'nuxt-auth-utils', '@nuxt/eslint'],
    runtimeConfig: {
        public: {
            appName: process.env.APP_NAME ?? 'Custodian',
            appDescriptor: process.env.APP_DESCRIPTOR ?? 'Nuxt app',
        },
    },
    app: {
        rootAttrs: {
            id: 'custodian',
        },
        head: {
            script: [
                { src: 'https://cdn.jsdelivr.net/npm/theme-change@2.0.2/index.js' },
            ],
        },
    },
    alias: {
        '.prisma/client/index-browser': './node_modules/.prisma/client/index-browser.js',
    },
    eslint: {
        config: {
            standalone: false,
        },
    },
})
