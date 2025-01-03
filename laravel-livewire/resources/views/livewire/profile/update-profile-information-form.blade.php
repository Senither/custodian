<?php

use App\Models\User;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Session;
use Illuminate\Validation\Rule;
use Livewire\Volt\Component;

new class extends Component {
    public string $name = '';
    public string $email = '';

    /**
     * Mount the component.
     */
    public function mount(): void
    {
        $this->name = Auth::user()->name;
        $this->email = Auth::user()->email;
    }

    /**
     * Update the profile information for the currently authenticated user.
     */
    public function updateProfileInformation(): void
    {
        $user = Auth::user();

        $validated = $this->validate([
            'name' => ['required', 'string', 'max:255'],
            'email' => ['required', 'string', 'lowercase', 'email', 'max:255', Rule::unique(User::class)->ignore($user->id)],
        ]);

        $user->fill($validated);

        if ($user->isDirty('email')) {
            $user->email_verified_at = null;
        }

        $user->save();

        $this->dispatch('profile-updated', name: $user->name);
    }

    /**
     * Send an email verification notification to the current user.
     */
    public function sendVerification(): void
    {
        $user = Auth::user();

        if ($user->hasVerifiedEmail()) {
            $this->redirectIntended(default: route('dashboard', absolute: false));

            return;
        }

        $user->sendEmailVerificationNotification();

        Session::flash('status', 'verification-link-sent');
    }
}; ?>

<section class="gap-4 grid grid-cols-1 md:grid-cols-5">
    <header class="md:col-span-2 px-6 sm:px-0 prose">
        <h3 class="font-semibold text-xl">Account details</h3>
        <p>Update your account's profile information and email address.</p>
    </header>

    <div class="flex flex-col flex-1 md:col-span-3 bg-base-100 border border-base-300 sm:rounded-md">
        <form wire:submit="updateProfileInformation">
            <div class="flex flex-col gap-4 px-4 py-3 w-full">
                <label class="form-control px-4 w-full">
                    <div class="label">
                        <span class="label-text">Name</span>
                    </div>
                    <input wire:model="name" type="text" placeholder="Your name" class="input-bordered w-full max-w-xs input" />
                    <x-input-error class="mt-2" :messages="$errors->get('name')" />
                </label>

                <label class="form-control px-4 w-full">
                    <div class="label">
                        <span class="label-text">Email</span>
                    </div>
                    <input wire:model="email" type="email" placeholder="name@company.com"
                           class="input-bordered w-full max-w-xs input" />
                    <x-input-error class="mt-2" :messages="$errors->get('email')" />
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
