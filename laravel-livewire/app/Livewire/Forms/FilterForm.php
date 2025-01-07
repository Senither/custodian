<?php

namespace App\Livewire\Forms;

use Livewire\Form;

class FilterForm extends Form
{
    public ?string $query = null;

    public ?string $status = null;

    public ?int $category = null;

    public ?int $priority = null;

    /**
     * Checks if the form has any filters applied.
     */
    public function hasFilters(): bool
    {
        return $this->query !== null
            || $this->status !== null
            || $this->category !== null
            || $this->priority !== null;
    }
}
