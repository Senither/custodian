<?php

namespace App\Livewire\Forms;

use Livewire\Form;

class FilterForm extends Form
{
    /**
     * The search query to filter the results by.
     */
    public ?string $query = null;

    /**
     * The task status to filter the results by.
     *
     * The value should be "finished" or "unfinished".
     */
    public ?string $status = null;

    /**
     * The category ID to filter the results by.
     */
    public ?int $category = null;

    /**
     * The priority ID to filter the results by.
     */
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
