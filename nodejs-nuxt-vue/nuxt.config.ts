// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: ['@prisma/nuxt', '@nuxtjs/tailwindcss', 'nuxt-auth-utils'],
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
    '.prisma/client/index-browser': './node_modules/.prisma/client/index-browser.js'
  },
})
