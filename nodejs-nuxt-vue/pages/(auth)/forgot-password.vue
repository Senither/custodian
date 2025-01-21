<script setup lang="ts">
definePageMeta({
    layout: 'auth',
    middleware: ['unauthenticated'],
})

const form = ref({
    email: '',
})

async function onSendPasswordResetLink() {
    console.log('send password reset link', form)
}
</script>

<template>
    <div class="flex bg-base-100 shadow-2xl w-96 card">
        <form class="card-body" @submit.prevent="onSendPasswordResetLink">
            <h3 class="font-semibold text-lg">Forgot your password?</h3>
            <p class="text-sm">Don't fret! Just type in your email and we will send you a code to reset your password!
            </p>

            <div class="form-control mt-4">
                <label class="label">
                    <span class="label-text">Your email</span>
                </label>
                <input v-model="form.email" type="email" placeholder="email@company.com" class="input-bordered input"
                    autofocus required />
            </div>

            <div class="form-control mt-6">
                <button type="submit" class="btn btn-primary">Reset password</button>
            </div>

            <p class="text-base-content text-sm">
                Remembered your password?
                <NuxtLink href="/login" class="link link-hover link-primary">Login here</NuxtLink>
            </p>
        </form>
    </div>

</template>
