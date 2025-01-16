<?php

use App\Livewire\Tasks\DeleteTaskModal;
use App\Livewire\Tasks\TaskCard;
use App\Livewire\Tasks\UpdateTaskModal;
use App\Models\Task;
use App\Models\User;
use Livewire\Livewire;

it('can render the component', function () {
    $this->actingAs($user = User::factory()->createQuietly());

    $task = Task::factory()->recycle($user)->create();

    Livewire::test(TaskCard::class, ['task' => $task])
        ->assertStatus(200)
        ->assertSee($task->message)
        ->assertSee($task->category->name)
        ->assertSee($task->priority->name)
        ->assertSet('completed', $task->status);
});

it('renders the edit and delete components', function () {
    $this->actingAs($user = User::factory()->createQuietly());

    $task = Task::factory()->recycle($user)->create();

    Livewire::test(TaskCard::class, ['task' => $task])
        ->assertSeeLivewire(UpdateTaskModal::class)
        ->assertSeeLivewire(DeleteTaskModal::class);
});

it('updates the task status when completed property changes', function () {
    $this->actingAs($user = User::factory()->createQuietly());

    $task = Task::factory()->recycle($user)->create([
        'status' => false,
    ]);

    Livewire::test(TaskCard::class, ['task' => $task])
        ->set('completed', true);

    expect($task->refresh()->status)->toBeTrue();
});

it('cannot render tasks for other users', function () {
    $task = Task::factory()->create();

    $this->actingAs(User::factory()->createQuietly());

    Livewire::test(TaskCard::class, ['task' => $task])
        ->assertForbidden();
});
