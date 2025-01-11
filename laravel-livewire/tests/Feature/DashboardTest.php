<?php

use App\Models\Task;
use App\Models\User;

it('redirects to login for unauthenticated requests', function () {
    $response = $this->get('/dashboard');

    $response
        ->assertStatus(302)
        ->assertRedirect(route('login'));
});

it('can render the dashboard', function () {
    $this->actingAs(User::factory()->create());

    $response = $this->get('/dashboard');

    $response->assertStatus(200);
});

it('renders the tasks components', function () {
    $this->actingAs($user = User::factory()->create());
    Task::factory()->recycle($user)->create();

    $response = $this->get('/dashboard');

    $response
        ->assertSeeLivewire('tasks.task-card')
        ->assertSeeLivewire('tasks.create-task-modal')
        ->assertSeeLivewire('tasks.render-dashboard-tasks');
});
