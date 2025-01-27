<script setup lang="ts">
definePageMeta({
    layout: 'auth',
    middleware: ['unauthenticated'],
})

const { fetch: refreshSession } = useUserSession()

const errors = ref(null)
const form = reactive({
    name: '',
    email: '',
    password: '',
    password_confirmation: '',
})

async function onRegister() {
    errors.value = null

    $fetch('/api/auth/register', {
        method: 'POST',
        body: form,
    }).then(async () => {
        await refreshSession()
        await navigateTo('/')
    }).catch((error) => {
        errors.value = error.data
    })
}
</script>

<template>
    <div class="flex bg-base-100 shadow-2xl w-96 card">
        <form class="card-body" @submit.prevent="onRegister">
            <h3 class="font-semibold text-lg">
                Sign up for Custodian
            </h3>

            <ErrorMessage v-model="errors" />

            <div class="form-control mt-4">
                <label class="label">
                    <span class="label-text">Your name</span>
                </label>
                <input
                    v-model="form.name" type="text" placeholder="Eg. John Doe" class="input-bordered input" autofocus
                    required autocomplete="name"
                >
                <InputError v-model="errors" name="name" />
            </div>

            <div class="form-control mt-4">
                <label class="label">
                    <span class="label-text">Your email</span>
                </label>
                <input
                    v-model="form.email" type="email" placeholder="email@company.com" class="input-bordered input"
                    required autocomplete="username"
                >
                <InputError v-model="errors" name="email" />
            </div>

            <div class="form-control mt-4">
                <label class="label">
                    <span class="label-text">Password</span>
                </label>
                <input
                    v-model="form.password" type="password" placeholder="********" class="input-bordered input"
                    required autocomplete="new-password"
                >
                <InputError v-model="errors" name="password" />
            </div>

            <div class="form-control mt-4">
                <label class="label">
                    <span class="label-text">Confirm Password</span>
                </label>
                <input
                    v-model="form.password_confirmation" type="password" placeholder="********"
                    class="input-bordered input" required autocomplete="new-password"
                >
                <InputError v-model="errors" name="password_confirmation" />
            </div>

            <div class="form-control mt-6">
                <button type="submit" class="btn btn-primary">
                    Create account
                </button>
            </div>

            <p class="text-base-content text-sm">
                Already have an account?
                <NuxtLink href="/login" class="link link-hover link-primary">
                    Login here
                </NuxtLink>
            </p>
        </form>
    </div>
</template>
