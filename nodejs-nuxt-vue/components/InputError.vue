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

    for (let err of errors) {
        if ((err?.path ?? []).includes(name)) {
            messages.value.push(err.message)
        }
    }
})
</script>

<template>
    <div class="font-medium text-error text-sm" v-if="messages.length > 0">
        <ul>
            <li v-for="message of messages">{{ message }}</li>
        </ul>
    </div>
</template>
