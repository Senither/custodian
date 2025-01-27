# PHP + Laravel + Livewire

This is a Laravel and Livewire version of Custodian running on PHP 8.2+

## Table of Contents

- [The Stack](#the-stack)
- [Features](#features)
- [Installation](#installation)
    - [Docker](#docker)
    - [Local Environment](#local-environment)
- [Testing](#testing)

## The Stack

The backend is built with [Laravel 11](https://laravel.com/) and [Livewire 3](https://livewire.laravel.com/), with test coverage using [Pest 3](https://pestphp.com/), and [Pint](https://laravel.com/docs/11.x/pint#main-content) to ensure formatting consistency.

For the frontend we're using [TailwindCSS 3](https://tailwindcss.com/) along with [DaisyUI](https://daisyui.com/) from the [template files](/template/), as well as [AlpineJS 3](https://alpinejs.dev/) that comes with Livewire to make modals and other in-browser features feel more snappy, all of this is built using [Vite 6](https://vite.dev/) and [NodeJS 18+](https://nodejs.org/).

The project is intended to be run within a Docker environment in production, and takes up around 150 MB of disk space when built, and around ~75 MB of RAM when running.

## Features

This is a list of features for this version of Custodian that are outside the scope of the core [Custodian requirements](/README.md#requirements).

- Event dispatching when creating new users
    - Event listeners hooks into this and create our task categories and priorities models for us
- Query scope limiter
    - This helps prevent touching any data that doesn't belong to the authenticated user
- Model observers
    - Automatically clears the cache when a model is created, updated, or deleted
- Model policies
    - Ensures that the authenticated user can only see the actions they're allowed to perform (create, update, and delete buttons are hidden when the user doesn't have the permission)

## Installation

Following the Custodian requirements, you can choose to install the project using Docker or by setting up a local development environment.

### Docker

Start by cloning the repository and going into this directory, all the following commands will be run from this directory.

```bash
git clone https://github.com/Senither/custodian.git
cd custodian/php-laravel-livewire
```

From here we can use the `docker compose` command to build and start the Docker container.

```bash
docker compose up -d
```

You can now open the application in your browser by visiting [http://localhost:8000](http://localhost:8000).

### Local Environment

Start by cloning the repository and going into this directory, all the following commands will be run from this directory.

```bash
git clone https://github.com/Senither/custodian.git
cd custodian/php-laravel-livewire
```

Next we can install all the dependencies using Composer and npm.

```bash
composer install
npm install
```

When the node modules have been installed we can compile the assets.

```bash
# Starts a watcher that will recompile the assets when they change
npm run dev
# Or to compile the assets once for production
npm run build
```

After the dependencies have been installed and our assets are ready we can copy the `.env.example` file to `.env` and generate a new application key.

```bash
cp .env.example .env
php artisan key:generate
```

Now we can run the migrations to create the database tables and seed the database with some test data.

```bash
php artisan migrate --seed
```

Finally we can start the development server and open the application in our browser.

```bash
php artisan serve
```

You can now open the application in your browser by visiting [http://localhost:8000](http://localhost:8000).

## Testing

To run the test suite first ensure that you have the project [installed and it's up and running](#installation), then you can run the following command to run the tests.

```bash
# Run the test suite using Artisan
php artisan test
# Or to run the test suite using Pest directly
vendor/bin/pest
```
