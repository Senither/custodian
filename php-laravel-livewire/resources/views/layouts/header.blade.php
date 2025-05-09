<header class="flex justify-between items-center px-4 py-2">
    <div class="mx-auto max-w-5xl navbar">
        <div class="navbar-start">
            <a wire:navigate href="{{ route('dashboard') }}">
                <h1 class="font-extrabold text-4xl text-primary">{{ config('app.name', 'Custodian') }}</h1>
                <p class="-mt-1.5 ml-6 font-semibold text-lg text-secondary">{{ config('app.descriptor', 'HTML Template') }}</p>
            </a>
        </div>

        <div class="flex items-center gap-4 navbar-end">
            <x-theme-selector />

            <div class="dropdown dropdown-end">
                <div tabindex="0" role="button" class="bg-base-100 m-1 btn">
                    <svg class="w-5 h-5" data-slot="icon" fill="none" stroke-width="1.5" stroke="currentColor"
                         viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round"
                              d="M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z">
                        </path>
                    </svg>
                </div>
                <ul tabindex="0"
                    class="z-[1] bg-base-100 shadow p-2 border border-base-200 rounded-box w-44 dropdown-content menu">
                    <li><a wire:navigate href="{{ route('profile') }}">Profile</a></li>
                    <livewire:header.logout />
                </ul>
            </div>
        </div>
    </div>
</header>
