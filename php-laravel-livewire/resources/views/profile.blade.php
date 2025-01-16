<x-app-layout>
    <div class="flex justify-between items-center mb-8 px-6 md:px-0">
        <h3 class="font-extrabold text-3xl">Your profile</h3>

        <a wire:navigate href="{{ route('dashboard') }}" class="flex items-center gap-2 link link-hover">
            <svg class="w-4 h-4" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                 viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 15.75 3 12m0 0 3.75-3.75M3 12h18">
                </path>
            </svg>
            Return to tasks
        </a>
    </div>

    <div class="flex flex-col gap-8 sm:px-6 py-4 rounded-md">
        <livewire:profile.update-profile-information-form />

        <div class="divider"></div>

        <livewire:profile.update-password-form />

        <div class="divider"></div>

        <livewire:profile.delete-user-form />
    </div>
</x-app-layout>
