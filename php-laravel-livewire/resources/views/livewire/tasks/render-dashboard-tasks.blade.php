<div class="flex flex-col gap-4">
    <div class="flex justify-between items-center px-6 md:px-0">
        <livewire:tasks.create-task-modal />

        <div wire:loading class="animate-pulse">
            <svg class="animate-spin w-6 h-6 text-base-content opacity-70" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                 viewBox="0 0 24 24"
                 xmlns="http://www.w3.org/2000/svg"
                 aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round"
                      d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99">
                </path>
            </svg>
        </div>

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
                        <input wire:model.live.debounce.250ms="filter.query" type="text" placeholder="Eg. My shopping list..."
                               class="input-bordered w-full max-w-xs input input-sm" />
                    </label>
                    <div class="-mb-2">
                        <div class="label">
                            <span class="label-text">Filter by status</span>
                        </div>
                        <div class="flex justify-start gap-3">
                            <div class="form-control">
                                <label class="flex justify-start gap-3 cursor-pointer label">
                                    <input wire:model.live="filter.status" type="radio" name="status-radio" value="finished"
                                           class="checked:radio-primary radio" />
                                    <span class="label-text">Finished</span>
                                </label>
                            </div>
                            <div class="form-control">
                                <label class="flex justify-start gap-3 cursor-pointer label">
                                    <input wire:model.live="filter.status" type="radio" name="status-radio" value="pending"
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
                        <select wire:model.live="filter.category" class="select-bordered select">
                            <option :value="0">Pick one</option>
                            @foreach ($categories as $category)
                                <option value="{{ $category->id }}">{{ $category->name }}</option>
                            @endforeach
                        </select>
                    </label>
                    <label class="form-control w-full max-w-xs">
                        <div class="label">
                            <span class="label-text">Filter by priority</span>
                        </div>
                        <select wire:model.live="filter.priority" class="select-bordered select">
                            <option :value="0">Pick one</option>
                            @foreach ($priorities as $priority)
                                <option value="{{ $priority->id }}">{{ $priority->name }}</option>
                            @endforeach
                        </select>
                    </label>
                    <button wire:click="resetFilter" class="mt-4 btn btn-secondary">Reset filters</button>
                </div>
                <!-- Filter Dropdown : End -->
            </div>
        </div>
    </div>

    <div class="flex flex-col gap-4 bg-base-100 shadow-lg px-6 py-4 rounded-md">
        @if ($tasks->isEmpty())
            @if ($filter->hasFilters())
                <div class="py-8 flex flex-col items-center justify-center gap-4">
                    <svg class="w-12 h-12 text-base-content" data-slot="icon" fill="none" stroke-width="1.5"
                         stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round"
                              d="M12 3c2.755 0 5.455.232 8.083.678.533.09.917.556.917 1.096v1.044a2.25 2.25 0 0 1-.659 1.591l-5.432 5.432a2.25 2.25 0 0 0-.659 1.591v2.927a2.25 2.25 0 0 1-1.244 2.013L9.75 21v-6.568a2.25 2.25 0 0 0-.659-1.591L3.659 7.409A2.25 2.25 0 0 1 3 5.818V4.774c0-.54.384-1.006.917-1.096A48.32 48.32 0 0 1 12 3Z">
                        </path>
                    </svg>

                    <p class="text-lg font-semibold text-center">No tasks found</p>
                </div>
            @else
                <div class="py-8 flex flex-col items-center justify-center gap-4">
                    <svg class="w-12 h-12 text-base-content opacity-25" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                         viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12"></path>
                    </svg>

                    <p class="text-lg font-semibold text-center">You have no tasks yet</p>
                </div>
            @endif
        @else
            @foreach ($tasks as $task)
                <livewire:tasks.task-card :task="$task" :key="$task->id . '-' . $task->updated_at" />
            @endforeach
        @endif
    </div>
</div>
