<?php

use App\Livewire\Tasks\RenderDashboardTasks;
use App\Models\Category;
use App\Models\Priority;
use App\Models\Task;
use App\Models\User;
use Livewire\Livewire;

it('can render the component', function () {
    $this->actingAs(User::factory()->createQuietly());

    Livewire::test(RenderDashboardTasks::class)
        ->assertStatus(200);
});

describe('filter', function () {
    beforeEach(function () {
        $this->actingAs($user = User::factory()->createQuietly());

        $category = Category::factory()->recycle($user)->create();
        $priority = Priority::factory()->recycle($user)->create();

        Task::factory()->recycle($user)->createMany([
            ['message' => 'Task 1', 'status' => false, 'priority_id' => $priority->id],
            ['message' => 'Task 2', 'status' => true, 'priority_id' => $priority->id],
            ['message' => 'Task 3', 'status' => false, 'category_id' => $category->id],
            ['message' => 'Task 4', 'status' => true, 'category_id' => $category->id],
            ['message' => 'Something unique', 'status' => false, 'category_id' => $category->id],
        ]);
    });

    describe('query', function () {
        it('can filter tasks by query', function (string $query, int $count, array $expected) {
            Livewire::test(RenderDashboardTasks::class)
                ->set('filter.query', $query)
                ->assertViewHas('tasks', function ($tasks) use ($count, $expected) {
                    expect($tasks->count())->toBe($count);
                    expect($tasks->pluck('message')->sort()->values()->all())->toEqual($expected);

                    return true;
                });
        })->with([
            'single' => [
                'query' => 'Task 2',
                'count' => 1,
                'expected' => ['Task 2'],
            ],
            'multiple' => [
                'query' => 'Task',
                'count' => 4,
                'expected' => ['Task 1', 'Task 2', 'Task 3', 'Task 4'],
            ],
            'none' => [
                'query' => 'Nothing',
                'count' => 0,
                'expected' => [],
            ],
        ]);
    });

    describe('status', function () {
        it('can filter tasks by status', function (string $status, int $count, array $expected) {
            Livewire::test(RenderDashboardTasks::class)
                ->set('filter.status', $status)
                ->assertViewHas('tasks', function ($tasks) use ($count, $expected) {
                    expect($tasks->count())->toBe($count);
                    expect($tasks->pluck('message')->sort()->values()->all())->toEqual($expected);

                    return true;
                });
        })->with([
            'finished' => [
                'status' => 'finished',
                'count' => 2,
                'expected' => ['Task 2', 'Task 4'],
            ],
            'unfinished' => [
                'status' => 'pending',
                'count' => 3,
                'expected' => ['Something unique', 'Task 1', 'Task 3'],
            ],
            'invalid' => [
                'status' => 'invalid',
                'count' => 5,
                'expected' => ['Something unique', 'Task 1', 'Task 2', 'Task 3', 'Task 4'],
            ],
        ]);
    });

    describe('category', function () {
        it('can filter tasks by category', function (int $category, int $count, array $expected) {
            Livewire::test(RenderDashboardTasks::class)
                ->set('filter.category', $category)
                ->assertViewHas('tasks', function ($tasks) use ($count, $expected) {
                    expect($tasks->count())->toBe($count);
                    expect($tasks->pluck('message')->sort()->values()->all())->toEqual($expected);

                    return true;
                });
        })->with([
            'single' => [
                'category' => 2,
                'count' => 1,
                'expected' => ['Task 1'],
            ],
            'multiple' => [
                'category' => 1,
                'count' => 3,
                'expected' => ['Something unique', 'Task 3', 'Task 4'],
            ],
            'none' => [
                'category' => 999,
                'count' => 0,
                'expected' => [],
            ],
        ]);
    });

    describe('priority', function () {
        it('can filter tasks by priority', function (int $priority, int $count, array $expected) {
            Livewire::test(RenderDashboardTasks::class)
                ->set('filter.priority', $priority)
                ->assertViewHas('tasks', function ($tasks) use ($count, $expected) {
                    expect($tasks->count())->toBe($count);
                    expect($tasks->pluck('message')->sort()->values()->all())->toEqual($expected);

                    return true;
                });
        })->with([
            'single' => [
                'priority' => 2,
                'count' => 1,
                'expected' => ['Task 3'],
            ],
            'multiple' => [
                'priority' => 1,
                'count' => 2,
                'expected' => ['Task 1', 'Task 2'],
            ],
            'none' => [
                'priority' => 999,
                'count' => 0,
                'expected' => [],
            ],
        ]);
    });

    describe('reset', function () {
        it('can reset the filter', function () {
            Livewire::test(RenderDashboardTasks::class)
                ->set('filter.query', 'Nothing')
                ->set('filter.status', 'invalid')
                ->set('filter.category', 999)
                ->set('filter.priority', 999)
                ->call('resetFilter')
                ->assertViewHas('tasks', function ($tasks) {
                    expect($tasks->count())->toBe(5);

                    return true;
                });
        });
    });
});
