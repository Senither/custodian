<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="csrf-token" content="{{ csrf_token() }}">

    <title>{{ config('app.name', 'Custodian') }}</title>

    <!-- Scripts -->
    <script src="https://cdn.jsdelivr.net/npm/theme-change@2.0.2/index.js"></script>
    @vite(['resources/css/app.css', 'resources/js/app.js'])
</head>

<body>
    <header class="top-0 fixed flex justify-between items-center px-6 py-4 w-full">
        <div>
            <h1 class="font-extrabold text-4xl text-primary">{{ config('app.name') }}</h1>
            <p class="-mt-1.5 ml-6 font-semibold text-lg text-secondary">{{ config('app.descriptor') }}</p>
        </div>

        <x-theme-selector />
    </header>

    <main class="bg-base-200 min-h-screen hero">
        <div class="flex-col hero-content">
            {{ $slot }}
        </div>
    </main>

    <footer class="flex justify-center items-center bg-base-100 -mt-[3.54rem] px-4 py-3 border-t border-base-300 w-full">
        @include('layouts.footer')
    </footer>
</body>

</html>
