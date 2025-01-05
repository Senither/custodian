<div
     x-bind:class="{
         'border-primary/10 bg-primary/5 shadow px-4 py-3 border rounded-md transition-opacity': true,
         'opacity-45 hover:opacity-80': $wire.completed,
     }">
    <div class="flex justify-between items-center">
        <div class="flex items-center gap-4">
            <input type="checkbox" class="checkbox checkbox-lg checked:checkbox-primary" wire:model.live="completed" />

            <div>
                <h4 class="font-semibold text-lg">{{ $task->message }}</h4>
                <div class="flex gap-3">
                    @if ($task->priority)
                        <span class="py-2 text-secondary-content cursor-pointer badge badge-secondary badge-sm">
                            priority: {{ mb_strtolower($task->priority->name) }}
                        </span>
                    @endif

                    @if ($task->category)
                        <span class="py-2 text-secondary-content cursor-pointer badge badge-secondary badge-sm">
                            category: {{ mb_strtolower($task->category->name) }}
                        </span>
                    @endif
                </div>
            </div>
        </div>

        <div class="flex gap-2">
            <livewire:tasks.update-task-modal :task="$task" :key="'update-task-modal-' . $task->id" />

            <button onclick="deleteTaskModal.showModal()"
                    class="bg-base-100 btn btn-square hover:btn-error">
                <svg class="w-5 h-5 text-content" data-slot="icon" fill="none" stroke-width="1.5"
                     stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"
                     aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round"
                          d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0">
                    </path>
                </svg>
            </button>
        </div>
    </div>
</div>
