<script setup lang="ts">
definePageMeta({
    middleware: ['authenticated'],
})

const isLoadingInitialPage = ref(true)
const tasks = ref([])
const filters = reactive({
    query: null,
    status: null,
    category: null,
    priority: null,
})

$fetch('/api/tasks').then((res) => {
    tasks.value = res.data
    isLoadingInitialPage.value = false
})

const updateTaskStatus = (task: Task, status: boolean) => {
    task.status = status
    $fetch(`/api/tasks/${task.id}`, {
        method: 'PUT',
        body: { status },
    })
}

const filteredTasks = computed(() => {
    return tasks.value.filter((task) => {
        return (!filters.query || task.message.toLowerCase().includes(filters.query.toLowerCase()))
            && (filters.status == null || task.status === filters.status)
            && (!filters.category || task.category.id === filters.category)
            && (!filters.priority || task.priority.id === filters.priority)
    })
})

const hasFilters = computed(() => {
    return filters.query
        || filters.status
        || filters.category
        || filters.priority
})

const resetFilter = () => {
    filters.query = null
    filters.status = null
    filters.category = null
    filters.priority = null
}
</script>

<template>
    <div class="flex flex-col gap-4">
        <div class="flex justify-between items-center px-6 md:px-0">
            <button class="btn btn-primary">Add Task</button>

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
                                <option :value="1">House Stuff</option>
                                <option :value="2">Work</option>
                                <option :value="3">Learning</option>
                                <option :value="4">Meeting</option>
                            </select>
                        </label>

                        <label class="form-control w-full max-w-xs">
                            <div class="label">
                                <span class="label-text">Filter by priority</span>
                            </div>
                            <select class="select-bordered select-sm select" v-model="filters.priority">
                                <option :value="null" disabled selected>Pick one</option>
                                <option :value="1">Low</option>
                                <option :value="2">Medium</option>
                                <option :value="3">High</option>
                                <option :value="4">Highest</option>
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

            <TaskCard v-else v-for="task in filteredTasks" :task="task" :key="task.id"
                @statusChanged="val => updateTaskStatus(task, val)" />
        </div>
    </div>
</template>
