<?php

namespace App\Livewire\Forms;

use Livewire\Form;

class FilterForm extends Form
{
    public ?string $query = null;

    public ?string $status = null;

    public ?int $category = null;

    public ?int $priority = null;
}
