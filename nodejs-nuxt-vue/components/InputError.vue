<script setup lang="ts">
const error = defineModel()
const messages = ref([])

const { name } = useAttrs()

watch(error, (value) => {
    messages.value = []

    if (value == null) {
        return
    }

    const errors = error.value?.data?.issues ?? []

    for (const err of errors) {
        if ((err?.path ?? []).includes(name)) {
            messages.value.push(err.message)
        }
    }
})
</script>

<template>
    <div v-if="messages.length > 0" class="font-medium text-error text-sm">
        <ul>
            <li v-for="message of messages" :key="message">
                {{ message }}
            </li>
        </ul>
    </div>
</template>
