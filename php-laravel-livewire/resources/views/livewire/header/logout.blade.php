<?php

use App\Livewire\Actions\Logout;
use Illuminate\Support\Facades\Auth;
use Livewire\Volt\Component;

new class extends Component
{
    /**
     * Logs out the currently authenticated user.
     */
    public function logout(Logout $logout): void
    {
        tap(Auth::user(), $logout(...));

        $this->redirect('/', navigate: true);
    }
}; ?>

<li><a wire:click="logout">Logout</a></li>
