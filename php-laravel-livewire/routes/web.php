<?php

use App\Livewire\Tasks\RenderDashboardTasks;
use Illuminate\Support\Facades\Route;

Route::redirect('/', '/dashboard');

Route::middleware(['auth', 'verified'])->group(function () {
    Route::get('dashboard', RenderDashboardTasks::class)->name('dashboard');
    Route::view('profile', 'profile')->name('profile');
});

require __DIR__.'/auth.php';
