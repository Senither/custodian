<?php

use App\Livewire\Tasks\DeleteTaskModal;
use App\Models\Task;
use App\Models\User;
use Livewire\Livewire;

it('can render the component', function () {
    $this->actingAs($user = User::factory()->createQuietly());

    $task = Task::factory()->recycle($user)->create();

    Livewire::test(DeleteTaskModal::class, ['task' => $task])
        ->assertStatus(200)
        ->assertSet('task', $task);
});

it('can delete a task', function () {
    $this->actingAs($user = User::factory()->createQuietly());

    $task = Task::factory()->recycle($user)->create();

    Livewire::test(DeleteTaskModal::class, ['task' => $task])
        ->call('delete')
        ->assertDispatched('reload-tasks');

    expect($task->exists())->toBeFalse();
});

it('cannot delete a task for another user', function () {
    $task = Task::factory()->create();

    $this->actingAs(User::factory()->createQuietly());

    Livewire::test(DeleteTaskModal::class, ['task' => $task])
        ->call('delete')
        ->assertForbidden();
});
