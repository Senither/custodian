<?php

use App\Events\UserCreated;
use App\Models\Category;
use App\Models\Priority;
use App\Models\Task;
use App\Models\User;
use Illuminate\Support\Facades\Event;

describe('events', function () {
    it('dispatches created event when creating new user', function () {
        Event::fake();

        $user = User::factory()->create();

        Event::assertDispatched(
            UserCreated::class,
            fn($event) => $event->user->id === $user->id,
        );
    });
});

describe('serialization', function () {
    it('hides password and remember_token', function () {
        $user = User::factory()->createQuietly([
            'password' => 'password',
            'remember_token' => 'token',
        ]);

        expect($user->toArray())
            ->not->toHaveKeys(['password', 'remember_token']);
    });
});

describe('relationships', function () {
    it('has task relationship', function () {
        $user = User::factory()->createQuietly();
        $task = Task::factory()->recycle($user)->create();

        expect($user->tasks)->toHaveCount(1);
        expect($user->tasks->first()->id)->toBe($task->id);
    });

    it('has priority relationship', function () {
        $user = User::factory()->createQuietly();
        $priority = Priority::factory()->recycle($user)->create();

        expect($user->priorities)->toHaveCount(1);
        expect($user->priorities->first()->id)->toBe($priority->id);
    });

    it('has category relationship', function () {
        $user = User::factory()->createQuietly();
        $category = Category::factory()->recycle($user)->create();

        expect($user->categories)->toHaveCount(1);
        expect($user->categories->first()->id)->toBe($category->id);
    });
});
