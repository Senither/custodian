<?php

namespace Database\Seeders;

use App\Models\Task;
use App\Models\User;
use Illuminate\Database\Seeder;

class UserSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        // User::factory(10)->create();

        $defaultUser = User::factory()->create([
            'name' => 'Test User',
            'email' => 'test@example.com',
        ]);

        Task::factory()
            ->recycle($defaultUser)
            ->recycle($defaultUser->priorities)
            ->recycle($defaultUser->categories)
            ->count(8)
            ->create();
    }
}
