<?php

namespace App\Livewire\Tasks;

use App\Livewire\Contracts\TaskModalComponent;
use App\Models\Task;
use Illuminate\Contracts\View\View;

class DeleteTaskModal extends TaskModalComponent
{
    /**
     * The task that should be deleted.
     */
    public Task $task;

    /**
     * Mounts the component.
     */
    public function mount(Task $task): void
    {
        $this->task = $task;
    }

    /**
     * Handles deleting the task.
     */
    public function delete(): void
    {
        $this->task->delete();

        $this
            ->closeModalWithEvent('reload-tasks')
            ->to(RenderDashboardTasks::class);
    }

    /**
     * Renders the component.
     */
    public function render(): View
    {
        return view('livewire.tasks.delete-task-modal');
    }
}
