<?php

use App\Livewire\Forms\LoginForm;
use Illuminate\Support\Facades\Session;
use Livewire\Attributes\Layout;
use Livewire\Volt\Component;

new #[Layout('layouts.guest')] class extends Component {
    public LoginForm $form;

    /**
     * Handle an incoming authentication request.
     */
    public function login(): void
    {
        $this->validate();

        $this->form->authenticate();

        Session::regenerate();

        $this->redirectIntended(default: route('dashboard', absolute: false), navigate: true);
    }
}; ?>

<div class="flex bg-base-100 shadow-2xl w-96 card">
    <form class="card-body" wire:submit="login">
        <h3 class="font-semibold text-lg">Sign in to {{ config('app.name') }}</h3>

        <x-auth-session-status class="mb-4" :status="session('status')" />

        <div class="form-control mt-4">
            <label class="label">
                <span class="label-text">Email</span>
            </label>
            <input wire:model="form.email" type="email" placeholder="email@company.com" class="input-bordered input" autofocus required />
            <x-input-error :messages="$errors->get('form.email')" class="mt-2" />
        </div>

        <div class="form-control">
            <label class="label">
                <span class="label-text">Password</span>
            </label>
            <input wire:model="form.password" type="password" placeholder="********" class="input-bordered input" required />
            <x-input-error :messages="$errors->get('form.password')" class="mt-2" />
        </div>

        <div class="flex justify-between gap-4">
            <label class="flex gap-2 cursor-pointer label">
                <input wire:model="form.remember" type="checkbox" class="checkbox checkbox-primary" />
                <span class="label-text">Remember me</span>
            </label>

            @if (Route::has('password.request'))
                <label class="label">
                    <a wire:navigate href="{{ route('password.request') }}" class="label-text link link-hover link-primary">
                        Forgot password?
                    </a>
                </label>
            @endif
        </div>

        <div class="form-control mt-6">
            <button type="submit" class="btn btn-primary">Login</button>
        </div>

        <p class="text-base-content text-sm">
            Not registered? <a wire:navigate href="{{ route('register') }}" class="link link-hover link-primary">Create account</a>
        </p>
    </form>
</div>
