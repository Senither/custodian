<script setup lang="ts">
const tasks = ref([
    {
        completed: false,
        title: 'Buy groceries',
        priority: 'high',
        category: 'house stuff',
    },
    {
        completed: false,
        title: 'Build Custodian using new stacks',
        priority: 'low',
        category: 'learning',
    },
    {
        completed: true,
        title: 'Finish client project',
        priority: 'highest',
        category: 'work',
    },
])
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
                            <input type="text" placeholder="Eg. My shopping list..."
                                class="input-bordered w-full max-w-xs input input-sm" />
                        </label>

                        <div class="-mb-2">
                            <div class="label">
                                <span class="label-text">Filter by status</span>
                            </div>

                            <div class="flex justify-start gap-3">
                                <div class="form-control">
                                    <label class="flex justify-start gap-3 cursor-pointer label">
                                        <input type="radio" name="status-radio" class="checked:radio-primary radio" />
                                        <span class="label-text">Finished</span>
                                    </label>
                                </div>
                                <div class="form-control">
                                    <label class="flex justify-start gap-3 cursor-pointer label">
                                        <input type="radio" name="status-radio" class="checked:radio-secondary radio" />
                                        <span class="label-text">Pending</span>
                                    </label>
                                </div>
                            </div>
                        </div>

                        <label class="form-control w-full max-w-xs">
                            <div class="label">
                                <span class="label-text">Filter by category</span>
                            </div>
                            <select class="select-bordered select-sm select">
                                <option disabled selected>Pick one</option>
                                <option>House Stuff</option>
                                <option>Work</option>
                                <option>Learning</option>
                                <option>Meeting</option>
                            </select>
                        </label>

                        <label class="form-control w-full max-w-xs">
                            <div class="label">
                                <span class="label-text">Filter by priority</span>
                            </div>
                            <select class="select-bordered select-sm select">
                                <option disabled selected>Pick one</option>
                                <option>Low</option>
                                <option>Medium</option>
                                <option>High</option>
                                <option>Highest</option>
                            </select>
                        </label>

                        <button class="mt-4 btn btn-secondary">Reset filters</button>
                    </div>
                    <!-- Filter Dropdown : End -->
                </div>
            </div>
        </div>

        <div class="flex flex-col gap-4 bg-base-100 shadow-lg px-6 py-4 rounded-md">
            <TaskCard v-for="task in tasks" :task="task" @statusChanged="val => task.completed = val" />
        </div>
    </div>
</template>
