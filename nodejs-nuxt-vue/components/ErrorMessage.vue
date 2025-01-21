<script setup lang="ts">
const error = defineModel()
const message = ref('')

watch(error, (value) => {
    if (value == null) {
        message.value = ''
        return
    }

    message.value = value.message
})
</script>

<template>
    <div class="font-medium text-error text-sm" v-if="message.length > 0">
        <p>{{ message }}</p>
    </div>
</template>
