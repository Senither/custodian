<script setup>
import { Dialog, DialogPanel } from '@headlessui/vue'

defineProps({
    open: Boolean,
})
</script>

<template>
    <Dialog :open="open" @close="$emit('close')">
        <div class="fixed inset-0 backdrop-blur-sm backdrop-grayscale" aria-hidden="true" />

        <div class="fixed inset-0 flex justify-center items-center p-4 w-screen">
            <DialogPanel class="p-0 modal-box">
                <slot />
            </DialogPanel>
        </div>
    </Dialog>
</template>
