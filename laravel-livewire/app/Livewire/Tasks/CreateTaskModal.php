<?php

namespace App\Livewire\Tasks;

use App\Livewire\Contracts\TaskModalComponent;
use Illuminate\Contracts\View\View;

class CreateTaskModal extends TaskModalComponent
{
    /**
     * Validates the form and saves the task.
     */
    public function createTask(): void
    {
        $this->form->store();

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
            view: 'livewire.tasks.create-task-modal',
            data: $this->getTaskRelationshipData(),
        );
    }
}
