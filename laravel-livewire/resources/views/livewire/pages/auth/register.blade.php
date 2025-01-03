<?php

use App\Models\User;
use Illuminate\Auth\Events\Registered;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;
use Illuminate\Validation\Rules;
use Livewire\Attributes\Layout;
use Livewire\Volt\Component;

new #[Layout('layouts.guest')] class extends Component {
    public string $name = '';
    public string $email = '';
    public string $password = '';
    public string $password_confirmation = '';

    /**
     * Handle an incoming registration request.
     */
    public function register(): void
    {
        $validated = $this->validate([
            'name' => ['required', 'string', 'max:255'],
            'email' => ['required', 'string', 'lowercase', 'email', 'max:255', 'unique:' . User::class],
            'password' => ['required', 'string', 'confirmed', Rules\Password::defaults()],
        ]);

        $validated['password'] = Hash::make($validated['password']);

        event(new Registered(($user = User::create($validated))));

        Auth::login($user);

        $this->redirect(route('dashboard', absolute: false), navigate: true);
    }
}; ?>

<div class="flex bg-base-100 shadow-2xl w-96 card">
    <form class="card-body" wire:submit="register">
        <h3 class="font-semibold text-lg">Sign up for Custodian</h3>

        <div class="form-control mt-4">
            <label class="label">
                <span class="label-text">Your name</span>
            </label>
            <input wire:model="name" type="text" placeholder="Eg. John Doe"
                   class="input-bordered input" autofocus required autocomplete="name" />
        </div>

        <div class="form-control mt-4">
            <label class="label">
                <span class="label-text">Your email</span>
            </label>
            <input wire:model="email" type="email" placeholder="email@company.com"
                   class="input-bordered input" required autocomplete="username" />
        </div>

        <div class="form-control mt-4">
            <label class="label">
                <span class="label-text">Password</span>
            </label>
            <input wire:model="password" type="password" placeholder="********"
                   class="input-bordered input" required autocomplete="new-password" />
        </div>

        <div class="form-control mt-4">
            <label class="label">
                <span class="label-text">Confirm Password</span>
            </label>
            <input wire:model="password_confirmation" type="password" placeholder="********"
                   class="input-bordered input" required autocomplete="new-password" />
        </div>

        <div class="form-control mt-6">
            <button type="submit" class="btn btn-primary">Create account</button>
        </div>

        <p class="text-base-content text-sm">
            Already have an account?
            <a wire:navigate href="{{ route('login') }}" class="link link-hover link-primary">Login here</a>
        </p>
    </form>
</div>
