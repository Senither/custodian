<?php

use App\Livewire\Actions\Logout;
use Illuminate\Support\Facades\Auth;
use Livewire\Volt\Component;

new class extends Component {
    public bool $showModal = false;
    public string $password = '';

    /**
     * Delete the currently authenticated user.
     */
    public function deleteUser(Logout $logout): void
    {
        $this->validate([
            'password' => ['required', 'string', 'current_password'],
        ]);

        tap(Auth::user(), $logout(...))->delete();

        $this->redirect('/', navigate: true);
    }
}; ?>

<div class="gap-4 grid grid-cols-1 md:grid-cols-5">
    <div class="md:col-span-2 px-6 sm:px-0 prose">
        <h3 class="font-semibold text-xl">Delete Account</h3>
        <p>Permanently delete your account.</p>
    </div>

    <div class="flex flex-1 border-error md:col-span-3 bg-error shadow p-2 border sm:rounded-md">
        <div class="bg-base-100 shadow p-4 border border-base-300 rounded-md max-w-none prose">
            <p>
                Once your account is deleted, all of its resources and data will be permanently deleted.
                Before deleting your account, please download any data or information that you wish to
                retain.
            </p>

            <a x-on:click="$wire.set('showModal', true)" class="text-error-content btn btn-error">Delete Account</a>
        </div>
    </div>

    @teleport('body')
        <dialog id="deleteAccountModal" class="backdrop-blur-sm backdrop-grayscale modal"
                x-bind:open="$wire.showModal"
                x-on:close.stop="$wire.set('showModal', false)"
                x-on:keydown.escape.window="$wire.set('showModal', false)"
                x-on:keydown.tab.prevent="$event.shiftKey || nextFocusable().focus()"
                x-on:keydown.shift.tab.prevent="prevFocusable().focus()">
            <div class="p-0 modal-box">
                <div class="flex items-center gap-4 p-6">
                    <svg class="w-8 h-8 text-primary" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                         viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round"
                              d="M12 18v-5.25m0 0a6.01 6.01 0 0 0 1.5-.189m-1.5.189a6.01 6.01 0 0 1-1.5-.189m3.75 7.478a12.06 12.06 0 0 1-4.5 0m3.75 2.383a14.406 14.406 0 0 1-3 0M14.25 18v-.192c0-.983.658-1.823 1.508-2.316a7.5 7.5 0 1 0-7.517 0c.85.493 1.509 1.333 1.509 2.316V18">
                        </path>
                    </svg>
                    <h2 class="text-lg font-medium">
                        {{ __('Are you sure you want to delete your account?') }}
                    </h2>
                </div>

                <form wire:submit.prevent="deleteUser" class="pb-6">
                    <p class="px-6 mt-1 text-sm">
                        {{ __('Once your account is deleted, all of its resources and data will be permanently deleted. Please enter your password to confirm you would like to permanently delete your account.') }}
                    </p>

                    <label class="px-6 mt-6 form-control w-full">
                        <div class="label">
                            <span class="label-text">Password</span>
                        </div>

                        <input wire:model="password"
                               id="password"
                               name="password"
                               type="password"
                               class="input-bordered w-full max-w-xs input"
                               placeholder="{{ __('Password') }}" />

                        <x-input-error :messages="$errors->get('password')" class="mt-2" />
                    </label>

                    <div class="divider"></div>

                    <div class="px-6 flex justify-end">
                        <a x-on:click="$wire.set('showModal', false)" class="btn btn-primary">
                            {{ __('Cancel') }}
                        </a>

                        <button type="submit" class="btn btn-error ms-3">
                            {{ __('Delete Account') }}
                        </button>

                    </div>
                </form>
            </div>
        </dialog>
    @endteleport
</div>
