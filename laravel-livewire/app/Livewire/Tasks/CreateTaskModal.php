<?php

namespace App\Livewire\Tasks;

use App\Livewire\Forms\TaskForm;
use App\Models\Category;
use App\Models\Priority;
use Illuminate\Contracts\View\View;
use Livewire\Component;

class CreateTaskModal extends Component
{
    /**
     * The form for update modal.
     */
    public TaskForm $form;

    /**
     * Determines if the modal is shown to the end-user or not.
     */
    public bool $showModal = false;

    /**
     * Validates the form and saves the task.
     */
    public function createTask(): void
    {
        $this->form->store();

        $this
            ->dispatch('task-created')
            ->to(RenderDashboardTasks::class);

        $this->showModal = false;
    }

    /**
     * Renders the component.
     */
    public function render(): View
    {
        return view('livewire.tasks.create-task-modal', [
            'categories' => Category::orderBy('name')->get(),
            'priorities' => Priority::get(),
        ]);
    }
}
