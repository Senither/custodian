<?php

namespace App\Livewire\Tasks;

use App\Models\Task;
use Illuminate\Contracts\View\View;
use Livewire\Attributes\Locked;
use Livewire\Component;

class TaskCard extends Component
{
    /**
     * The task that should be rendered in the card.
     */
    #[Locked]
    public Task $task;

    /**
     * Determines if the task is completed.
     */
    public bool $completed;

    /**
     * Mounts the component with the given task.
     */
    public function mount(Task $task): void
    {
        $this->authorize('view', $task);

        $this->task = $task;

        $this->completed = $task->status;
    }

    /**
     * Updates the status of the task when the completed property changes.
     */
    public function updatedCompleted(): void
    {
        $this->authorize('update', $this->task);

        $this->task->update([
            'status' => $this->completed,
        ]);
    }

    /**
     * Renders the component.
     */
    public function render(): View
    {
        return view('livewire.tasks.task-card');
    }
}
