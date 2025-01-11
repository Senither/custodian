<?php

use App\Livewire\Tasks\UpdateTaskModal;
use App\Models\Category;
use App\Models\Priority;
use App\Models\Task;
use App\Models\User;
use Livewire\Livewire;

it('can render the component', function () {
    $this->actingAs($user = User::factory()->createQuietly());

    $task = Task::factory()->recycle($user)->create();

    Livewire::test(UpdateTaskModal::class, ['task' => $task])
        ->assertStatus(200)
        ->assertSet('form.task', $task)
        ->assertSet('form.message', $task->message)
        ->assertSet('form.priority', $task->priority_id)
        ->assertSet('form.category', $task->category_id);
});

it('can update a task', function () {
    $this->actingAs($user = User::factory()->createQuietly());

    $task = Task::factory()->recycle($user)->create();

    $priority = Priority::factory()->recycle($user)->create();
    $category = Category::factory()->recycle($user)->create();

    Livewire::test(UpdateTaskModal::class, ['task' => $task])
        ->set('form.message', 'Updated Task')
        ->set('form.priority', $priority->id)
        ->set('form.category', $category->id)
        ->call('saveTask')
        ->assertDispatched('reload-tasks');

    expect($task->refresh())
        ->message->toBe('Updated Task')
        ->priority_id->toBe($priority->id)
        ->category_id->toBe($category->id);
});

it('cannot update a task for another user', function () {
    $task = Task::factory()->create();

    $this->actingAs(User::factory()->createQuietly());

    Livewire::test(UpdateTaskModal::class, ['task' => $task])
        ->set('form.message', 'Updated Task')
        ->call('saveTask')
        ->assertForbidden();
});
