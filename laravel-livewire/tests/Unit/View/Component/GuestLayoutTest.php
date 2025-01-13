<?php

use App\View\Components\GuestLayout;
use Illuminate\View\Component;
use Illuminate\View\View;

it('extends Illuminate\View\Component', function () {
    expect(GuestLayout::class)->toExtend(Component::class);
});

it('renders the expected view', function () {
    $component = new GuestLayout();

    expect($component->render())
        ->toBeInstanceOf(View::class)
        ->getName()->toBe('layouts.guest');
});
