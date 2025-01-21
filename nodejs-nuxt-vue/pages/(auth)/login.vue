<script setup lang="ts">
definePageMeta({
    layout: 'auth',
    middleware: ['unauthenticated'],
})

const { fetch: refreshSession } = useUserSession()

const form = reactive({
    email: '',
    password: '',
})

async function onLogin() {
    $fetch('/api/login', {
        method: 'POST',
        body: form,
    }).then(async () => {
        await refreshSession()
        await navigateTo('/')
    }).catch(() => alert('Bad credentials'))
}
</script>

<template>
    <div class="flex bg-base-100 shadow-2xl w-96 card">
        <form class="card-body" @submit.prevent="onLogin">
            <h3 class="font-semibold text-lg">Sign in to Custodian</h3>

            <div class="form-control mt-4">
                <label class="label">
                    <span class="label-text">Email</span>
                </label>
                <input v-model="form.email" type="email" placeholder="email@company.com" class="input-bordered input"
                    autofocus required />
            </div>

            <div class="form-control">
                <label class="label">
                    <span class="label-text">Password</span>
                </label>
                <input v-model="form.password" type="password" placeholder="********" class="input-bordered input"
                    required />
            </div>

            <div class="flex justify-between gap-4">
                <label class="flex gap-2 cursor-pointer label">
                    <input type="checkbox" class="checkbox checkbox-primary" />
                    <span class="label-text">Remember me</span>
                </label>

                <label class="label">
                    <NuxtLink href="/forgot-password" class="label-text link link-hover link-primary">
                        Forgot password?
                    </NuxtLink>
                </label>
            </div>

            <div class="form-control mt-6">
                <button type="submit" class="btn btn-primary">Login</button>
            </div>

            <p class="text-base-content text-sm">
                Not registered?
                <NuxtLink href="/register" class="link link-hover link-primary">
                    Create account
                </NuxtLink>
            </p>
        </form>
    </div>
</template>
