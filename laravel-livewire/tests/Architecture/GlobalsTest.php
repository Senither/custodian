<?php

arch('global helpers are not used')
    ->expect(['env', 'fake'])
    ->not->toBeUsed();

arch('debug helpers are not used')
    ->expect(['dd', 'dump', 'ray', 'die', 'var_dump', 'sleep', 'usleep'])
    ->not->toBeUsed();

arch('PHP process functions are not used')
    ->expect(['exit', 'die', 'eval', 'exec', 'passthru', 'proc_open', 'shell_exec', 'system'])
    ->not->toBeUsed();

arch('PHP filesystem functions are not used')
    ->expect(['include', 'include_once', 'require', 'require_once'])
    ->not->toBeUsed();

arch('PHP array functions are not used')
    ->expect(['array', 'compact'])
    ->not->toBeUsed();
