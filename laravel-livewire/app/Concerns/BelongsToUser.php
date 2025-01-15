<?php

namespace App\Concerns;

use App\Models\Scopes\BelongsToUserScope;
use App\Models\User;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Support\Facades\Auth;

trait BelongsToUser
{
    /**
     * Boot the belongs to user trait for a model.
     *
     * @return void
     */
    public static function bootBelongsToUser()
    {
        static::addGlobalScope(new BelongsToUserScope);

        static::saving(function ($model) {
            return tap($model, function ($model) {
                if (Auth::check()) {
                    $model->user_id = Auth::id();
                }
            });
        });
    }

    /**
     * Get the relationship for the user that owns the model.
     */
    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class);
    }

    /**
     * Determines if the model belongs to the given user.
     */
    public function belongsToUser(int|User $user): bool
    {
        if ($user instanceof User) {
            $user = $user->id;
        }

        return $this->user_id === $user;
    }
}
