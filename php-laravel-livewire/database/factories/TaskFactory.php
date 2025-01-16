<?php

namespace Database\Factories;

use App\Models\Category;
use App\Models\Priority;
use App\Models\User;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Task>
 */
class TaskFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'user_id' => User::factory(),
            'status' => $this->faker->boolean,
            'message' => $this->faker->sentence,
            'priority_id' => Priority::factory(),
            'category_id' => Category::factory(),
        ];
    }
}
