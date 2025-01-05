<div class="flex flex-col gap-4">
    <div class="flex justify-between px-6 md:px-0">
        <livewire:tasks.create-task-modal />

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
                        <select class="select-bordered select">
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
                        <select class="select-bordered select">
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
        @foreach ($tasks as $task)
            <livewire:tasks.task-card :task="$task" :key="$task->id . '-' . $task->updated_at" />
        @endforeach
    </div>

    {{-- <x-slot name="modals">
        <!-- Delete Task Modal : Start -->
        <dialog id="deleteTaskModal" class="backdrop-blur-sm backdrop-grayscale modal">
            <div class="p-0 modal-box">
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
                        You're about the delete the "Buy groceries" task, this is a permanent action and cannot be reversed.
                    </p>
                    <p>Are you sure you want to continue?</p>
                </div>
                <form method="dialog" class="modal-backdrop">
                    <div class="gap-4 grid grid-cols-2 mt-4 px-6 py-4 border-t border-base-300">
                        <button class="btn btn-warning">Delete Task</button>
                        <button class="bg-base-200 border border-base-300 btn">Close</button>
                    </div>
                </form>
            </div>
        </dialog>
        <!-- Delete Task Modal : End -->
    </x-slot> --}}
</div>
