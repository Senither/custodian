<?php

namespace App\Observers;

use App\Concerns\InteractsWithTaskRelationshipCache;
use App\Models\Priority;

class PriorityObserver
{
    use InteractsWithTaskRelationshipCache;

    /**
     * Handle the Priority "created" event.
     */
    public function created(Priority $priority): void
    {
        $this->clearTaskRelationshipDataCache($priority->user_id);
    }

    /**
     * Handle the Priority "updated" event.
     */
    public function updated(Priority $priority): void
    {
        $this->clearTaskRelationshipDataCache($priority->user_id);
    }

    /**
     * Handle the Priority "deleted" event.
     */
    public function deleted(Priority $priority): void
    {
        $this->clearTaskRelationshipDataCache($priority->user_id);
    }
}
