<script setup lang="ts">
definePageMeta({
    middleware: ['authenticated'],
})

const { clear: clearSession } = useUserSession()

await clearSession()
await navigateTo('/login')
</script>
