<?php

namespace App\Listeners;

use App\Events\UserCreated;

class CreateDefaultCategoriesForUser
{
    /**
     * The default categories.
     */
    protected array $defaultCategories = [
        'House Stuff',
        'Work',
        'Learning',
        'Meeting',
    ];

    /**
     * Handle the event.
     */
    public function handle(UserCreated $event): void
    {
        foreach ($this->defaultCategories as $category) {
            $event->user->categories()->create([
                'name' => $category,
            ]);
        }
    }
}
