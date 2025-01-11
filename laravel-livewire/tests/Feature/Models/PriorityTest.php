<?php

use App\Models\Priority;
use App\Models\User;
use Illuminate\Support\Facades\Cache;

describe('relationships', function () {
    it('has user relationship', function () {
        $priority = Priority::factory()->create();

        expect($priority->user)->toBeInstanceOf(User::class);
    });
});

describe('observers', function () {
    it('clears task relationship data cache on created', function () {
        $user = User::factory()->createQuietly();

        Cache::shouldReceive('forget')
            ->once()
            ->with('task-relationships::' . $user->id);

        Priority::factory()->recycle($user)->create();
    });

    it('clears task relationship data cache on updated', function () {
        $priority = Priority::factory()->create();

        Cache::shouldReceive('forget')
            ->once()
            ->with('task-relationships::' . $priority->user_id);

        $priority->update(['name' => 'Updated Name']);
    });

    it('clears task relationship data cache on deleted', function () {
        $priority = Priority::factory()->create();

        Cache::shouldReceive('forget')
            ->once()
            ->with('task-relationships::' . $priority->user_id);

        $priority->delete();
    });
});
