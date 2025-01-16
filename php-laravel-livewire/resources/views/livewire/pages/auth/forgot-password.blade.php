<?php

use Illuminate\Support\Facades\Password;
use Livewire\Attributes\Layout;
use Livewire\Volt\Component;

new #[Layout('layouts.guest')] class extends Component {
    public string $email = '';

    /**
     * Send a password reset link to the provided email address.
     */
    public function sendPasswordResetLink(): void
    {
        $this->validate([
            'email' => ['required', 'string', 'email'],
        ]);

        // We will send the password reset link to this user. Once we have attempted
        // to send the link, we will examine the response then see the message we
        // need to show to the user. Finally, we'll send out a proper response.
        $status = Password::sendResetLink($this->only('email'));

        if ($status != Password::RESET_LINK_SENT) {
            $this->addError('email', __($status));

            return;
        }

        $this->reset('email');

        session()->flash('status', __($status));
    }
}; ?>

<div class="flex bg-base-100 shadow-2xl w-96 card">
    <form class="card-body" wire:submit="sendPasswordResetLink">
        <h3 class="font-semibold text-lg">Forgot your password?</h3>
        <p class="text-sm">Don't fret! Just type in your email and we will send you a code to reset your password!</p>

        <x-auth-session-status :status="session('status')" />

        <div class="form-control mt-4">
            <label class="label">
                <span class="label-text">Your email</span>
            </label>
            <input wire:model="email" type="email" placeholder="email@company.com" class="input-bordered input" autofocus required />
            <x-input-error :messages="$errors->get('email')" class="mt-2" />
        </div>

        <div class="form-control mt-6">
            <button type="submit" class="btn btn-primary">Reset password</button>
        </div>

        <p class="text-base-content text-sm">
            Remembered your password?
            <a wire:navigate href="{{ route('login') }}" class="link link-hover link-primary">Login here</a>
        </p>
    </form>
</div>
