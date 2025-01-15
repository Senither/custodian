<?php

namespace App\Livewire\Contracts;

use App\Concerns\InteractsWithTaskRelationshipCache;
use App\Livewire\Forms\TaskForm;
use Livewire\Component;
use Livewire\Features\SupportEvents\Event;

abstract class TaskModalComponent extends Component
{
    use InteractsWithTaskRelationshipCache;

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
}
