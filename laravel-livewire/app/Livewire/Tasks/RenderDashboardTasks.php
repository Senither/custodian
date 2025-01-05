<?php

namespace App\Livewire\Tasks;

use App\Models\Task;
use Illuminate\Contracts\View\View;
use Livewire\Attributes\Layout;
use Livewire\Attributes\On;
use Livewire\Component;

#[Layout('layouts.app')]
class RenderDashboardTasks extends Component
{
    /**
     * Renders the component.
     */
    #[On('task-updated')]
    public function render(): View
    {
        return view('livewire.tasks.render-dashboard-tasks', [
            'tasks' => Task::with('priority', 'category')
                ->orderBy('status')
                ->latest()
                ->get(),
        ]);
    }
}
