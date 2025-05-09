<?php

namespace App\Livewire\Tasks;

use App\Concerns\InteractsWithTaskRelationshipCache;
use App\Livewire\Forms\FilterForm;
use App\Models\Task;
use Illuminate\Contracts\View\View;
use Illuminate\Database\Eloquent\Builder;
use Livewire\Attributes\Layout;
use Livewire\Attributes\Lazy;
use Livewire\Attributes\On;
use Livewire\Component;

#[Lazy]
#[Layout('layouts.app')]
class RenderDashboardTasks extends Component
{
    use InteractsWithTaskRelationshipCache;

    /**
     * The filter form that helps determine how to filter the tasks.
     */
    public FilterForm $filter;

    /**
     * Resets the filter form back to the default.
     */
    public function resetFilter(): void
    {
        $this->filter->reset();
    }

    /**
     * Renders the placeholder while the component is being loaded.
     */
    public function placeholder(): View
    {
        return view('placeholders.render-dashboard-tasks-spinner');
    }

    /**
     * Renders the component.
     */
    #[On('reload-tasks')]
    public function render(): View
    {
        return view('livewire.tasks.render-dashboard-tasks', [
            ...$this->getTaskRelationshipData(),
            'tasks' => $this->applyFiltersToQuery(
                Task::with('priority', 'category'),
            )->orderBy('status')->latest()->get(),
        ]);
    }

    /**
     * Applies the form filters to the query builder for tasks,
     * allowing for searching and filtering the result.
     */
    protected function applyFiltersToQuery(Builder $query): Builder
    {
        // Apply the query filter.
        $query->when(\mb_strlen($this->filter->query) > 0, function (Builder $query) {
            return $query->where('message', 'like', "%{$this->filter->query}%");
        });

        // Apply the status filter.
        $query->when($this->filter->status != null, function (Builder $query) {
            return match ($this->filter->status) {
                'finished' => $query->where('status', true),
                'pending' => $query->where('status', false),
                default => $query,
            };
        });

        // Apply the category filter.
        $query->when($this->filter->category != null, function (Builder $query) {
            return $query->where('category_id', $this->filter->category);
        });

        // Apply the priority filter.
        $query->when($this->filter->priority != null, function (Builder $query) {
            return $query->where('priority_id', $this->filter->priority);
        });

        return $query;
    }
}
