<?php

use App\Models\Category;
use App\Models\User;
use Illuminate\Support\Facades\Cache;

describe('relationships', function () {
    it('has user relationship', function () {
        $category = Category::factory()->create();

        expect($category->user)->toBeInstanceOf(User::class);
    });
});

describe('observers', function () {
    it('clears task relationship data cache on created', function () {
        $user = User::factory()->createQuietly();

        Cache::shouldReceive('forget')
            ->once()
            ->with('task-relationships::' . $user->id);

        Category::factory()->recycle($user)->create();
    });

    it('clears task relationship data cache on updated', function () {
        $category = Category::factory()->create();

        Cache::shouldReceive('forget')
            ->once()
            ->with('task-relationships::' . $category->user_id);

        $category->update(['name' => 'Updated Name']);
    });

    it('clears task relationship data cache on deleted', function () {
        $category = Category::factory()->create();

        Cache::shouldReceive('forget')
            ->once()
            ->with('task-relationships::' . $category->user_id);

        $category->delete();
    });
});
