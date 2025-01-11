<?php

use App\Livewire\Tasks\CreateTaskModal;
use App\Models\Category;
use App\Models\Priority;
use App\Models\Task;
use App\Models\User;
use Livewire\Livewire;

it('can render the component', function () {
    $this->actingAs(User::factory()->createQuietly());

    Livewire::test(CreateTaskModal::class)
        ->assertStatus(200);
});

it('can create a new task', function () {
    $this->actingAs($user = User::factory()->createQuietly());

    $category = Category::factory()->recycle($user)->create();
    $priority = Priority::factory()->recycle($user)->create();

    expect(Task::count())->toBe(0);

    Livewire::test(CreateTaskModal::class)
        ->set('form.message', 'A new task')
        ->set('form.category', $category->id)
        ->set('form.priority', $priority->id)
        ->call('createTask')
        ->assertDispatched('reload-tasks');

    expect(Task::count())->toBe(1);
    expect(Task::first())
        ->message->toBe('A new task')
        ->category_id->toBe($category->id)
        ->priority_id->toBe($priority->id);
});
