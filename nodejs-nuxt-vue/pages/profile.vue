<script setup lang="ts">
definePageMeta({
    middleware: ['authenticated'],
})

$fetch('/api/user').then((response) => {
    userForm.name = response.data.name
    userForm.email = response.data.email
})

const profileActionMessage = ref()
const passwordActionMessage = ref()

const deleteAccountModal = ref(false)

const errors = ref(null)
const userForm = reactive({
    name: '',
    email: '',
})

const passwordForm = reactive({
    current_password: '',
    password: '',
    password_confirmation: '',
})

const updateProfileInformation = () => {
    $fetch('/api/user', {
        method: 'PUT',
        body: userForm,
    }).then(() => {
        errors.value = null

        profileActionMessage.value.show()
    }).catch((error) => {
        errors.value = error.data
    })
}

const updatePassword = () => {
    $fetch('/api/user', {
        method: 'PUT',
        body: passwordForm,
    }).then(() => {
        errors.value = null
        passwordForm.current_password = ''
        passwordForm.password = ''
        passwordForm.password_confirmation = ''

        passwordActionMessage.value.show()
    }).catch((error) => {
        passwordForm.current_password = ''
        errors.value = error.data
    })
}
</script>

<template>
    <div class="flex justify-between items-center mb-8 px-6 md:px-0">
        <h3 class="font-extrabold text-3xl">Your profile</h3>

        <NuxtLink href="/dashboard" class="flex items-center gap-2 link link-hover">
            <svg class="w-4 h-4" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 15.75 3 12m0 0 3.75-3.75M3 12h18">
                </path>
            </svg>

            Return to tasks
        </NuxtLink>
    </div>

    <div class="flex flex-col gap-8 sm:px-6 py-4 rounded-md">
        <section class="gap-4 grid grid-cols-1 md:grid-cols-5">
            <header class="md:col-span-2 px-6 sm:px-0 prose">
                <h3 class="font-semibold text-xl">Account details</h3>
                <p>Update your account's profile information and email address.</p>
            </header>

            <div class="flex flex-col flex-1 md:col-span-3 bg-base-100 border border-base-300 sm:rounded-md">
                <form @submit.prevent="updateProfileInformation">
                    <div class="flex flex-col gap-4 px-4 py-3 w-full">
                        <label class="form-control px-4 w-full">
                            <div class="label">
                                <span class="label-text">Name</span>
                            </div>
                            <input v-model="userForm.name" type="text" placeholder="Your name"
                                class="input-bordered w-full max-w-xs input" />
                            <InputError v-model="errors" name="name" />
                        </label>

                        <label class="form-control px-4 w-full">
                            <div class="label">
                                <span class="label-text">Email</span>
                            </div>
                            <input v-model="userForm.email" type="email" placeholder="name@company.com"
                                class="input-bordered w-full max-w-xs input" />
                            <InputError v-model="errors" name="email" />
                        </label>
                    </div>

                    <div class="my-0 py-0 divider"></div>

                    <div class="flex justify-end items-center px-4 py-3">
                        <ActionMessage ref="profileActionMessage">
                            Saved changes
                        </ActionMessage>
                        <button class="btn btn-primary">Save</button>
                    </div>
                </form>
            </div>
        </section>

        <div class="divider"></div>

        <section class="gap-4 grid grid-cols-1 md:grid-cols-5">
            <header class="md:col-span-2 px-6 sm:px-0 prose">
                <h3 class="font-semibold text-xl">Update Password</h3>
                <p>Ensure your account is using a long, random password to stay secure.</p>
            </header>

            <div class="flex flex-col flex-1 md:col-span-3 bg-base-100 border border-base-300 sm:rounded-md">
                <form @submit.prevent="updatePassword">
                    <div class="flex flex-col gap-4 px-4 py-3 w-full">
                        <label class="form-control px-4 w-full">
                            <div class="label">
                                <span class="label-text">Current password</span>
                            </div>
                            <input v-model="passwordForm.current_password" type="password"
                                class="input-bordered w-full max-w-xs input" />
                            <InputError v-model="errors" name="current_password" />
                        </label>
                        <label class="form-control px-4 w-full">
                            <div class="label">
                                <span class="label-text">New password</span>
                            </div>
                            <input v-model="passwordForm.password" type="password"
                                class="input-bordered w-full max-w-xs input" />
                            <InputError v-model="errors" name="password" />
                        </label>
                        <label class="form-control px-4 w-full">
                            <div class="label">
                                <span class="label-text">Confirm password</span>
                            </div>
                            <input v-model="passwordForm.password_confirmation" type="password"
                                class="input-bordered w-full max-w-xs input" />
                            <InputError v-model="errors" name="password_confirmation" />
                        </label>
                    </div>

                    <div class="my-0 py-0 divider"></div>

                    <div class="flex justify-end items-center px-4 py-3">
                        <ActionMessage ref="passwordActionMessage">
                            Password changed successfully
                        </ActionMessage>
                        <button class="btn btn-primary">Save</button>
                    </div>
                </form>
            </div>
        </section>

        <div class="divider"></div>

        <!-- <livewire:profile.delete-user-form /> -->
    </div>

    <div class="gap-4 grid grid-cols-1 md:grid-cols-5">
        <div class="md:col-span-2 px-6 sm:px-0 prose">
            <h3 class="font-semibold text-xl">Delete Account</h3>
            <p>Permanently delete your account.</p>
        </div>

        <div class="flex flex-1 border-error md:col-span-3 bg-error shadow p-2 border sm:rounded-md">
            <div class="bg-base-100 shadow p-4 border border-base-300 rounded-md max-w-none prose">
                <p>
                    Once your account is deleted, all of its resources and data will be permanently deleted.
                    Before deleting your account, please download any data or information that you wish to
                    retain.
                </p>

                <button class="text-error-content btn btn-error" @click="deleteAccountModal = true">
                    Delete Account
                </button>
            </div>
        </div>
    </div>

    <Modal :open="deleteAccountModal" @close="deleteAccountModal = false">
        <div class="flex items-center gap-4 p-6">
            <svg class="w-8 h-8 text-error" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round"
                    d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z">
                </path>
            </svg>

            <h2 class="font-medium text-lg">
                Are you sure you want to delete your account?
            </h2>
        </div>

        <form class="pb-6" method="dialog">
            <p class="mt-1 px-6 text-sm">
                Once your account is deleted, all of its resources and data will be permanently deleted.
                Please enter your password to confirm you would like to permanently delete your account.
            </p>

            <label class="form-control mt-6 px-6 w-full">
                <div class="label">
                    <span class="label-text">Password</span>
                </div>

                <input id="password" name="password" type="password" class="input-bordered w-full max-w-xs input"
                    placeholder="Password" />
            </label>

            <div class="divider"></div>

            <div class="flex justify-end px-6">
                <a @click="deleteAccountModal = false" class="btn btn-primary">
                    Cancel
                </a>

                <button type="submit" class="btn btn-error ms-3">
                    Delete Account
                </button>
            </div>
        </form>
    </Modal>
</template>
