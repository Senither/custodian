<script setup lang="ts">
// TODO: Check if the current user is authenticated, if they
// are we can redirect them to the dashboard, if not we can
// redirect them to the login page

await navigateTo({
    path: '/dashboard'
})
</script>
