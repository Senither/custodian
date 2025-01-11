<?php

use App\Models\Category;
use App\Models\Priority;
use App\Models\Task;
use App\Models\User;

describe('relationships', function () {
    it('has user relationship', function () {
        $task = Task::factory()->create();

        expect($task->user)->toBeInstanceOf(User::class);
    });

    it('has priority relationship', function () {
        $task = Task::factory()->create();

        expect($task->priority)->toBeInstanceOf(Priority::class);
    });

    it('has category relationship', function () {
        $task = Task::factory()->create();

        expect($task->category)->toBeInstanceOf(Category::class);
    });
});
