<?php

namespace App\Observers;

use App\Concerns\InteractsWithTaskRelationshipCache;
use App\Models\Category;

class CategoryObserver
{
    use InteractsWithTaskRelationshipCache;

    /**
     * Handle the Category "created" event.
     */
    public function created(Category $category): void
    {
        $this->clearTaskRelationshipDataCache($category->user_id);
    }

    /**
     * Handle the Category "updated" event.
     */
    public function updated(Category $category): void
    {
        $this->clearTaskRelationshipDataCache($category->user_id);
    }

    /**
     * Handle the Category "deleted" event.
     */
    public function deleted(Category $category): void
    {
        $this->clearTaskRelationshipDataCache($category->user_id);
    }
}
