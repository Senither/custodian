# Custodian

Custodian is a simple TODO app that has been written in multiple different languages and stacks. The goal of the project is to try and challenge myself to build the same app in different languages and stacks as I learn them, and to see how the different languages and stacks compare to each other.

## Table of Contents

- [Requirements](#requirements)
- [Versions](#versions)
- [License](#license)

## Requirements

All the different versions of Custodian should be built with the same basic functionality in mind. The app **MUST** be able to:

- Authenticate users
    - This means sign ups, signin, and password resets
- Allow users to create and manage their tasks (create, update and delete)
- Filter tasks by the following:
    - completion status
    - priority
    - category
    - search query
- Allow users to edit their profile (Change name, email, and password)

Additionally, regardless of the stack the app is built on the app itself should support a [SQLite database](https://www.sqlite.org/) and be able to run within a Docker container.

When possible the app should also be built using the [template files](template) provided in this repository. These files are meant to provide a basic structure for the app and to help ensure that the different versions of Custodian are as similar as possible.

## Versions

Each version of Custodian can be found in its respective directory. The following is a list of the different versions of Custodian that have been built:

- [Nuxt + Vue](node-nuxt-vue/)
    - **Language**: Node v22 (Can work with v18+)
    - **Backend**: Nuxt, Nuxt Auth Utils, Zod, Prisma, ESLint
    - **Frontend**: TailwindCSS, DaisyUI, Vue, Vue Router
    - **Docker Image Size**: ~177 MB

- [Laravel + Livewire](php-laravel-livewire/)
    - **Language**: PHP v8.2+
    - **Backend**: Laravel, Livewire, Pest PHP, Pint
    - **Frontend**: TailwindCSS, DaisyUI, AlpineJS, Vite
    - **Docker Image Size**: ~150 MB

## License

This project is open-sourced software licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.
