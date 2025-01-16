<?php

namespace App\Listeners;

use App\Events\UserCreated;

class CreateDefaultPrioritiesForUser
{
    /**
     * The default priorities.
     */
    protected array $defaultPriorities = [
        'Low',
        'Medium',
        'High',
        'Highest',
    ];

    /**
     * Handle the event.
     */
    public function handle(UserCreated $event): void
    {
        foreach ($this->defaultPriorities as $priority) {
            $event->user->priorities()->create([
                'name' => $priority,
            ]);
        }
    }
}
