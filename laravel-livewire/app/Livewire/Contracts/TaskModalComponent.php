<?php

namespace App\Livewire\Contracts;

use App\Livewire\Forms\TaskForm;
use App\Models\Category;
use App\Models\Priority;
use Livewire\Component;
use Livewire\Features\SupportEvents\Event;

abstract class TaskModalComponent extends Component
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
     * Closes the modal.
     */
    protected function closeModal(): void
    {
        $this->showModal = false;
    }

    /**
     * Closes the modal and dispatches a browser event.
     */
    public function closeModalWithEvent(string $event): Event
    {
        $this->closeModal();

        return $this->dispatch($event);
    }

    /**
     * Gets the view data for the task.
     */
    protected function getTaskViewData(): array
    {
        if (! $this->showModal) {
            return [
                'categories' => [],
                'priorities' => [],
            ];
        }

        return \once(fn () => [
            'categories' => Category::orderBy('name')->get(),
            'priorities' => Priority::get(),
        ]);
    }
}
