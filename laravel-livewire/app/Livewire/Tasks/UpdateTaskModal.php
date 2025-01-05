<?php

namespace App\Livewire\Tasks;

use App\Livewire\Forms\TaskForm;
use App\Models\Category;
use App\Models\Priority;
use App\Models\Task;
use Illuminate\Contracts\View\View;
use Livewire\Component;

class UpdateTaskModal extends Component
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
            ->dispatch('task-updated')
            ->to(RenderDashboardTasks::class);

        $this->showModal = false;
    }

    /**
     * Renders the component.
     */
    public function render(): View
    {
        return view('livewire.tasks.update-task-modal', [
            'categories' => Category::orderBy('name')->get(),
            'priorities' => Priority::get(),
        ]);
    }
}
