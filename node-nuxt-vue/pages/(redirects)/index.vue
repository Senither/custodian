<script setup lang="ts">
const { loggedIn } = useUserSession()

if (loggedIn.value) {
    await navigateTo('/dashboard')
}
else {
    await navigateTo('/login')
}
</script>
