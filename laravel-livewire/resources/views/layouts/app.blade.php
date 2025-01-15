<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>{{ config('app.name', 'Custodian') }}</title>

    <!-- Scripts -->
    <script src="https://cdn.jsdelivr.net/npm/theme-change@2.0.2/index.js"></script>
    @vite(['resources/css/app.css', 'resources/js/app.js'])
</head>

<body class="flex flex-col justify-between bg-base-200 min-h-screen">
    @include('layouts.header')

    <main class="flex flex-col gap-4 mx-auto mt-8 md:mt-12 mb-auto max-w-3xl container">
        {{ $slot }}
    </main>

    <footer class="items-center bg-base-100 mt-12 p-4 border-t border-base-300 text-base-content footer">
        @include('layouts.footer')
    </footer>

    @if (isset($modals))
        {{ $modals }}
    @endif
</body>

</html>
