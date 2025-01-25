<script setup lang="ts">
const emit = defineEmits([
    'statusChanged',
    'taskDeleted',
])

const props = defineProps({
    priorities: Array,
    categories: Array,
})

const { task } = useAttrs()
const status = ref(task.status)

const isEditing = ref(false)
const isDeleting = ref(false)

const errors = ref(null)
const form = ref({
    message: '',
    category_id: null,
    priority_id: null,
})

watch(status, () => {
    emit('statusChanged', status.value)
})

const openEditModal = () => {
    form.value = {
        message: task.message,
        category_id: task.category_id,
        priority_id: task.priority_id,
    }
    isEditing.value = true
}

const closeEditModal = () => {
    errors.value = null
    form.value = {
        message: '',
        category_id: null,
        priority_id: null,
    }

    isEditing.value = false
}

const saveTask = () => {
    $fetch(`/api/tasks/${task.id}`, {
        method: 'PUT',
        body: {
            message: form.value.message,
            category_id: form.value.category_id,
            priority_id: form.value.priority_id
        },
    }).then(() => {
        errors.value = null

        task.category = props.categories.find(c => c.id == form.value.category_id)
        task.priority = props.priorities.find(p => p.id == form.value.priority_id)
        task.message = form.value.message

        closeEditModal()
    }).catch((error) => {
        errors.value = error.data
    })
}

const deleteTask = () => {
    $fetch(`/api/tasks/${task.id}`, {
        method: 'DELETE',
    }).then(() => {
        isDeleting.value = false

        emit('taskDeleted', task.id)
    })
}
</script>

<template>
    <div :class="{
        'border-primary/10 bg-primary/5 shadow px-4 py-3 border rounded-md transition-opacity': true,
        'opacity-45 hover:opacity-80': status,
    }">
        <div class="flex justify-between items-center">
            <div class="flex items-center gap-4">
                <input v-model="status" type="checkbox" class="checkbox checkbox-lg checked:checkbox-primary" />

                <div class="cursor-pointer" @click="status = !status">
                    <h4 class="font-semibold text-lg">{{ task.message }}</h4>
                    <div class="flex gap-3">
                        <span class="py-2 text-secondary-content cursor-pointer badge badge-secondary badge-sm">
                            priority: {{ task.priority.name }}
                        </span>

                        <span class="py-2 text-secondary-content cursor-pointer badge badge-secondary badge-sm">
                            category: {{ task.category.name }}
                        </span>
                    </div>
                </div>
            </div>

            <div class="flex gap-2">
                <button @click="isDeleting = true" class="bg-base-100 btn btn-square hover:btn-error">
                    <svg class="w-5 h-5 text-content" data-slot="icon" fill="none" stroke-width="1.5"
                        stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0">
                        </path>
                    </svg>
                </button>

                <button @click="openEditModal" class="bg-base-100 btn btn-square hover:btn-primary">
                    <svg class="w-5 h-5 text-content" data-slot="icon" fill="none" stroke-width="1.5"
                        stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10">
                        </path>
                    </svg>
                </button>
            </div>
        </div>

        <Modal :open="isEditing" @close="closeEditModal">
            <div class="flex items-center gap-4 p-6">
                <svg class="w-8 h-8 text-primary" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                        viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round"
                            d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10">
                    </path>
                </svg>
                <h3 class="font-bold text-lg">Editing Task</h3>
            </div>

            <form @submit.prevent="saveTask" class="flex flex-col gap-4 px-6">
                <label class="form-control w-full">
                    <div class="label">
                        <span class="label-text">What is the task?</span>
                    </div>
                    <input v-model="form.message" type="text" placeholder="Eg. Do the thing..."
                            class="input-bordered w-full input" />
                    <InputError v-model="errors" name="message" class="mt-2" />
                </label>

                <div class="gap-4 grid grid-cols-1 md:grid-cols-2">
                    <label class="form-control">
                        <div class="label">
                            <span class="label-text">Category</span>
                        </div>
                        <select v-model="form.category_id" class="select-bordered select">
                            <option disabled>Pick one</option>
                            <option v-for="val of categories" :value="val.id" :key="val.id">{{ val.name }}</option>
                        </select>
                        <InputError v-model="errors" name="category_id" class="mt-2" />
                    </label>

                    <label class="form-control">
                        <div class="label">
                            <span class="label-text">Priority</span>
                        </div>
                        <select v-model="form.priority_id" class="select-bordered select">
                            <option disabled>Pick one</option>
                            <option v-for="val of priorities" :value="val.id" :key="val.id">{{ val.name }}</option>
                        </select>
                        <InputError v-model="errors" name="priority_id" class="mt-2" />
                    </label>
                </div>
            </form>

            <div method="dialog" class="modal-backdrop">
                <div class="gap-4 grid grid-cols-2 mt-4 px-6 py-4 border-t border-base-300">
                    <button @click="saveTask" class="text-primary-content btn btn-primary">Save Task</button>
                    <a @click="isEditing = false" class="bg-base-200 border border-base-300 btn" tabindex="0">Close</a>
                </div>
            </div>
        </Modal>

        <Modal :open="isDeleting" @close="isDeleting = false">
            <div class="flex items-center gap-4 p-6">
                <svg class="w-8 h-8 text-warning" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                        viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round"
                            d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z">
                    </path>
                </svg>
                <h3 class="font-bold text-lg">Are you sure you want to continue?</h3>
            </div>
            <div class="px-6 prose">
                <p>
                    You're about the delete the "{{ task.message }}" task, this is a permanent action and cannot be reversed.
                </p>
                <p>Are you sure you want to continue?</p>
            </div>
            <form @submit.prevent="deleteTask" class="modal-backdrop">
                <div class="gap-4 grid grid-cols-2 mt-4 px-6 py-4 border-t border-base-300">
                    <button type="submit" class="btn btn-warning">Delete Task</button>
                    <a @click="isDeleting = false" class="bg-base-200 border border-base-300 btn" tabindex="0">Close</a>
                </div>
            </form>
        </Modal>
    </div>
</template>
