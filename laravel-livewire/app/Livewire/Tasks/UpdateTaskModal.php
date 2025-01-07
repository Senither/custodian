<?php

namespace App\Livewire\Tasks;

use App\Livewire\Contracts\TaskModalComponent;
use App\Models\Task;
use Illuminate\Contracts\View\View;

class UpdateTaskModal extends TaskModalComponent
{
    /**
     * Mounts the component.
     */
    public function mount(Task $task): void
    {
        $this->form->setTask($task);
    }

    /**
     * Validates the form and saves the task.
     */
    public function saveTask(): void
    {
        $this->form->update();

        $this
            ->closeModalWithEvent('reload-tasks')
            ->to(RenderDashboardTasks::class);
    }

    /**
     * Renders the component.
     */
    public function render(): View
    {
        return view(
            view: 'livewire.tasks.update-task-modal',
            data: $this->getTaskRelationshipData(),
        );
    }
}
