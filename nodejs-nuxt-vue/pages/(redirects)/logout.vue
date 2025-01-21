<script setup lang="ts">
const { clear: clearSession } = useUserSession()

await clearSession()
await navigateTo('/login')
</script>
