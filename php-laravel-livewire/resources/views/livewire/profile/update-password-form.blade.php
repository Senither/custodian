<?php

use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;
use Illuminate\Validation\Rules\Password;
use Illuminate\Validation\ValidationException;
use Livewire\Volt\Component;

new class extends Component {
    public string $current_password = '';
    public string $password = '';
    public string $password_confirmation = '';

    /**
     * Update the password for the currently authenticated user.
     */
    public function updatePassword(): void
    {
        try {
            $validated = $this->validate([
                'current_password' => ['required', 'string', 'current_password'],
                'password' => ['required', 'string', Password::defaults(), 'confirmed'],
            ]);
        } catch (ValidationException $e) {
            $this->reset('current_password', 'password', 'password_confirmation');

            throw $e;
        }

        Auth::user()->update([
            'password' => Hash::make($validated['password']),
        ]);

        $this->reset('current_password', 'password', 'password_confirmation');

        $this->dispatch('password-updated');
    }
}; ?>

<section class="gap-4 grid grid-cols-1 md:grid-cols-5">
    <header class="md:col-span-2 px-6 sm:px-0 prose">
        <h3 class="font-semibold text-xl">Update Password</h3>
        <p>Ensure your account is using a long, random password to stay secure.</p>
    </header>

    <div class="flex flex-col flex-1 md:col-span-3 bg-base-100 border border-base-300 sm:rounded-md">
        <form wire:submit="updatePassword">
            <div class="flex flex-col gap-4 px-4 py-3 w-full">
                <label class="form-control px-4 w-full">
                    <div class="label">
                        <span class="label-text">Current password</span>
                    </div>
                    <input wire:model="current_password" type="password" class="input-bordered w-full max-w-xs input" />
                    <x-input-error :messages="$errors->get('current_password')" class="mt-2" />
                </label>
                <label class="form-control px-4 w-full">
                    <div class="label">
                        <span class="label-text">New password</span>
                    </div>
                    <input wire:model="password" type="password" class="input-bordered w-full max-w-xs input" />
                    <x-input-error :messages="$errors->get('password')" class="mt-2" />
                </label>
                <label class="form-control px-4 w-full">
                    <div class="label">
                        <span class="label-text">Confirm password</span>
                    </div>
                    <input wire:model="password_confirmation" type="password" class="input-bordered w-full max-w-xs input" />
                    <x-input-error :messages="$errors->get('password_confirmation')" class="mt-2" />
                </label>
            </div>

            <div class="my-0 py-0 divider"></div>

            <div class="flex justify-end items-center px-4 py-3">
                <x-action-message class="mr-3" on="profile-updated">
                    {{ __('Saved.') }}
                </x-action-message>

                <button class="btn btn-primary">Save</button>
            </div>
        </form>
    </div>
</section>
