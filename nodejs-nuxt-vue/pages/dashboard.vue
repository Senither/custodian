<script setup lang="ts">
definePageMeta({
    middleware: ['authenticated'],
})

const isLoadingInitialPage = ref(true)

const isCreating = ref(false)

const tasks = ref([])
const categories = ref([])
const priorities = ref([])

const errors = ref(null)
const form = reactive({
    message: '',
    category_id: null,
    priority_id: null,
})

const filters = reactive({
    query: null,
    status: null,
    category: null,
    priority: null,
})

onMounted(() => {
    loadTasks()
})

$fetch('/api/categories').then((res) => {
    categories.value = res.data
})

$fetch('/api/priorities').then((res) => {
    priorities.value = res.data
})

const updateTaskStatus = (task: Task, status: boolean) => {
    task.status = status
    $fetch(`/api/tasks/${task.id}`, {
        method: 'PUT',
        body: { status },
    })
}

const loadTasks = () => {
    $fetch('/api/tasks').then((res) => {
        tasks.value = res.data
        isLoadingInitialPage.value = false
    })
}

const filteredTasks = computed(() => {
    return tasks.value.filter((task) => {
        return (!filters.query || task.message.toLowerCase().includes(filters.query.toLowerCase()))
            && (filters.status == null || task.status === filters.status)
            && (!filters.category || task.category.id === filters.category)
            && (!filters.priority || task.priority.id === filters.priority)
    })
        .sort((a, b) => a.message < b.message ? 1 : -1)
        .sort((a, b) => a.status > b.status ? 1 : -1)
})

const hasFilters = computed(() => {
    return filters.query
        || filters.status != null
        || filters.category
        || filters.priority
})

const resetFilter = () => {
    filters.query = null
    filters.status = null
    filters.category = null
    filters.priority = null
}

const createTask = () => {
    $fetch('/api/tasks', {
        method: 'POST',
        body: form,
    }).then(() => {
        loadTasks()
        closeCreateModal()
    }).catch((error) => {
        errors.value = error.data
    })
}

const closeCreateModal = () => {
    isCreating.value = false

    errors.value = null

    form.message = ''
    form.category_id = null
    form.priority_id = null
}
</script>

<template>
    <div class="flex flex-col gap-4">
        <div class="flex justify-between items-center px-6 md:px-0">
            <button @click="isCreating = true" class="btn btn-primary">Add Task</button>

            <div class="dropdown dropdown-end">
                <div tabindex="0" role="button" class="bg-base-100 m-1 btn">
                    <svg class="w-5 h-5 text-base-content" data-slot="icon" fill="none" stroke-width="1.5"
                        stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M12 3c2.755 0 5.455.232 8.083.678.533.09.917.556.917 1.096v1.044a2.25 2.25 0 0 1-.659 1.591l-5.432 5.432a2.25 2.25 0 0 0-.659 1.591v2.927a2.25 2.25 0 0 1-1.244 2.013L9.75 21v-6.568a2.25 2.25 0 0 0-.659-1.591L3.659 7.409A2.25 2.25 0 0 1 3 5.818V4.774c0-.54.384-1.006.917-1.096A48.32 48.32 0 0 1 12 3Z">
                        </path>
                    </svg>
                </div>
                <div tabindex="0" class="z-[1] bg-base-100 shadow p-2 rounded-box dropdown-content menu">
                    <!-- Filter Dropdown : Start -->
                    <div class="flex flex-col gap-4 px-3 py-2 min-w-52">
                        <label class="form-control w-full max-w-xs">
                            <div class="label">
                                <span class="label-text">Search for task</span>
                            </div>
                            <input v-model="filters.query" type="text" placeholder="Eg. My shopping list..."
                                class="input-bordered w-full max-w-xs input input-sm" />
                        </label>

                        <div class="-mb-2">
                            <div class="label">
                                <span class="label-text">Filter by status</span>
                            </div>

                            <div class="flex justify-start gap-3">
                                <div class="form-control">
                                    <label class="flex justify-start gap-3 cursor-pointer label">
                                        <input v-model="filters.status" type="radio" name="status-radio" :value="true"
                                            class="checked:radio-primary radio" />
                                        <span class="label-text">Finished</span>
                                    </label>
                                </div>
                                <div class="form-control">
                                    <label class="flex justify-start gap-3 cursor-pointer label">
                                        <input v-model="filters.status" type="radio" name="status-radio" :value="false"
                                            class="checked:radio-secondary radio" />
                                        <span class="label-text">Pending</span>
                                    </label>
                                </div>
                            </div>
                        </div>

                        <label class="form-control w-full max-w-xs">
                            <div class="label">
                                <span class="label-text">Filter by category</span>
                            </div>
                            <select class="select-bordered select-sm select" v-model="filters.category">
                                <option :value="null" disabled selected>Pick one</option>
                                <option v-for="val of categories" :value="val.id" :key="val.id">
                                    {{ val.name }}
                                </option>
                            </select>
                        </label>

                        <label class="form-control w-full max-w-xs">
                            <div class="label">
                                <span class="label-text">Filter by priority</span>
                            </div>
                            <select class="select-bordered select-sm select" v-model="filters.priority">
                                <option :value="null" disabled selected>Pick one</option>
                                <option v-for="val of priorities" :value="val.id" :key="val.id">
                                    {{ val.name }}
                                </option>
                            </select>
                        </label>

                        <button class="mt-4 btn btn-secondary" @click="resetFilter">Reset filters</button>
                    </div>
                    <!-- Filter Dropdown : End -->
                </div>
            </div>
        </div>

        <div v-if="isLoadingInitialPage"
            class="flex flex-col items-center gap-4 bg-base-100 shadow-lg px-6 py-4 rounded-md">
            <span class="py-8 loading loading-bars loading-lg"></span>
        </div>

        <div v-else class="flex flex-col gap-4 bg-base-100 shadow-lg px-6 py-4 rounded-md">
            <div v-if="filteredTasks.length == 0">
                <div v-if="hasFilters" class="flex flex-col justify-center items-center gap-4 py-8">
                    <svg class="w-12 h-12 text-base-content" data-slot="icon" fill="none" stroke-width="1.5"
                        stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M12 3c2.755 0 5.455.232 8.083.678.533.09.917.556.917 1.096v1.044a2.25 2.25 0 0 1-.659 1.591l-5.432 5.432a2.25 2.25 0 0 0-.659 1.591v2.927a2.25 2.25 0 0 1-1.244 2.013L9.75 21v-6.568a2.25 2.25 0 0 0-.659-1.591L3.659 7.409A2.25 2.25 0 0 1 3 5.818V4.774c0-.54.384-1.006.917-1.096A48.32 48.32 0 0 1 12 3Z">
                        </path>
                    </svg>

                    <p class="font-semibold text-center text-lg">No tasks found</p>
                </div>

                <div v-else class="flex flex-col justify-center items-center gap-4 py-8">
                    <svg class="opacity-25 w-12 h-12 text-base-content" data-slot="icon" fill="none" stroke-width="1.5"
                        stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12"></path>
                    </svg>

                    <p class="font-semibold text-center text-lg">You have no tasks yet</p>
                </div>
            </div>

            <TaskCard v-else v-for="task in filteredTasks" :task="task" :key="task.id" :categories="categories"
                :priorities="priorities" @statusChanged="val => updateTaskStatus(task, val)" @taskDeleted="loadTasks" />
        </div>

        <Modal :open="isCreating" @close="closeCreateModal">
            <div class="flex items-center gap-4 p-6">
                <svg class="w-8 h-8 text-primary" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                    viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round"
                        d="M12 18v-5.25m0 0a6.01 6.01 0 0 0 1.5-.189m-1.5.189a6.01 6.01 0 0 1-1.5-.189m3.75 7.478a12.06 12.06 0 0 1-4.5 0m3.75 2.383a14.406 14.406 0 0 1-3 0M14.25 18v-.192c0-.983.658-1.823 1.508-2.316a7.5 7.5 0 1 0-7.517 0c.85.493 1.509 1.333 1.509 2.316V18">
                    </path>
                </svg>
                <h3 class="font-bold text-lg">Creating new Task</h3>
            </div>

            <form @submit.prevent="createTask" class="flex flex-col gap-4 px-6">
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
                            <option :value="null">Pick one</option>
                            <option v-for="val of categories" :value="val.id" :key="val.id">{{ val.name }}</option>
                        </select>
                        <InputError v-model="errors" name="category_id" class="mt-2" />
                    </label>

                    <label class="form-control">
                        <div class="label">
                            <span class="label-text">Priority</span>
                        </div>
                        <select v-model="form.priority_id" class="select-bordered select">
                            <option :value="null">Pick one</option>
                            <option v-for="val of priorities" :value="val.id" :key="val.id">{{ val.name }}</option>
                        </select>
                        <InputError v-model="errors" name="priority_id" class="mt-2" />
                    </label>
                </div>
            </form>

            <div method="dialog" class="modal-backdrop">
                <div class="gap-4 grid grid-cols-2 mt-4 px-6 py-4 border-t border-base-300">
                    <button @click="createTask" class="text-primary-content btn btn-primary">Create Task</button>
                    <a @click="closeCreateModal" class="bg-base-200 border border-base-300 btn" tabindex="0">Close</a>
                </div>
            </div>
        </Modal>
    </div>
</template>
