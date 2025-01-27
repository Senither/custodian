<script setup lang="ts">
const error = defineModel()
const message = ref('')

watch(error, (value) => {
    if (value == null) {
        message.value = ''
        return
    }

    message.value = value.reason
})
</script>

<template>
    <div v-if="message.length > 0" class="font-medium text-error text-sm">
        <p>{{ message }}</p>
    </div>
</template>
