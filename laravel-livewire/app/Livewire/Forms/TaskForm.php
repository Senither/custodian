<?php

namespace App\Livewire\Forms;

use App\Models\Task;
use Livewire\Attributes\Validate;
use Livewire\Form;

class TaskForm extends Form
{
    /**
     * The task that should be used when updating the form.
     */
    public ?Task $task;

    #[Validate('required|string|max:255')]
    public ?string $message = null;

    #[Validate('required|integer|exists:priorities,id')]
    public ?int $priority = null;

    #[Validate('required|integer|exists:categories,id')]
    public ?int $category = null;

    /**
     * Sets the task for the form.
     */
    public function setTask(Task $task): void
    {
        $this->task = $task;

        $this->message = $task->message;
        $this->priority = $task->priority_id;
        $this->category = $task->category_id;
    }

    /**
     * Validates the form and creates a new task.
     */
    public function store(): void
    {
        $this->validate();

        Task::create([
            'status' => false,
            'message' => $this->message,
            'category_id' => $this->category,
            'priority_id' => $this->priority,
        ]);

        $this->reset();
    }

    /**
     * Saves the form by updating the task.
     */
    public function update(): void
    {
        $this->validate();

        $this->task->update([
            'message' => $this->message,
            'category_id' => $this->category,
            'priority_id' => $this->priority,
        ]);
    }
}
