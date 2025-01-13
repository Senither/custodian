<?php

use App\View\Components\AppLayout;
use Illuminate\View\Component;
use Illuminate\View\View;

it('extends Illuminate\View\Component', function () {
    expect(AppLayout::class)->toExtend(Component::class);
});

it('renders the expected view', function () {
    $component = new AppLayout();

    expect($component->render())
        ->toBeInstanceOf(View::class)
        ->getName()->toBe('layouts.app');
});
