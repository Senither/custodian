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

            <livewire:tasks.delete-task-modal :task="$task" :key="'delete-task-modal-' . $task->id" />
        </div>
    </div>
</div>
