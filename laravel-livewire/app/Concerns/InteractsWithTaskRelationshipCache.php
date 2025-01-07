<?php

namespace App\Concerns;

use App\Models\Category;
use App\Models\Priority;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Cache;

trait InteractsWithTaskRelationshipCache
{
    /**
     * The task relationship data cache.
     *
     * This is used to store the relationship data in memory so we
     * can avoid hitting the cache multiple times during a single
     * request when the value is already loaded.
     */
    protected static ?array $taskRelationshipData = null;

    /**
     * Loads the task relationship data from cache, if the data is
     * not already cached it will be fetched from the database
     * and stored for the next fifteen minutes.
     */
    protected static function loadTaskRelationshipData()
    {
        return Cache::remember(
            key: 'task-relationships::'.Auth::id(),
            ttl: fn () => now()->addMinutes(15),
            callback: fn () => [
                'categories' => Category::orderBy('name')->get(),
                'priorities' => Priority::get(),
            ],
        );
    }

    /**
     * Gets the task relationship data.
     */
    protected function getTaskRelationshipData(): array
    {
        return static::$taskRelationshipData ??= static::loadTaskRelationshipData();
    }

    /**
     * Clears the task relationship data cache for the user with
     * the given ID, if no user ID is provided it will default
     * to the currently authenticated user.
     */
    protected function clearTaskRelationshipDataCache(?int $userId): void
    {
        Cache::forget('task-relationships::'.($userId ?? Auth::id()));
    }
}
