# Node + Nuxt + Vue

This is a Nuxt 3 and Vue version of Custodian running on Node 18+ (Preferably node v22)

## Table of Contents

- [The Stack](#the-stack)
- [Features](#features)
- [Installation](#installation)
    - [Docker](#docker)
    - [Local Environment](#local-environment)
- [Testing](#testing)

## The Stack

The backend is built with [Nuxt 3](https://nuxt.com/) and [Vue 3](https://vuejs.org/), with [Primsa](https://www.prisma.io/) for connection and managing the database, and [Zod](https://zod.dev/) for validating requests, and [ESLint](https://eslint.org/) to ensure formatting consistency.

For the frontend we're using [TailwindCSS 3](https://tailwindcss.com/) along with [DaisyUI](https://daisyui.com/) from the [template files](/template/), as well as [Vue Server Components](https://vuejs.org/guide/scaling-up/ssr.html) that comes with Nuxt and Vue, all of this is built using the [Node Server Nuxt preset](https://nuxt.com/docs/getting-started/deployment#presets) and [NodeJS v22](https://nodejs.org/).

The project is intended to be run within a Docker environment in production, and takes up around 177 MB of disk space when built, and around ~40 MB of RAM when running.

## Features

This is a list of features for this version of Custodian that are outside the scope of the core [Custodian requirements](/README.md#requirements).

> [!NOTE]
> _There are currently no extra features added to this version of Custodian_

## Installation

Following the Custodian requirements, you can choose to install the project using Docker or by setting up a local development environment.

### Docker

Start by cloning the repository and going into this directory, all the following commands will be run from this directory.

```bash
git clone https://github.com/Senither/custodian.git
cd custodian/node-nuxt-vue
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
cd custodian/node-nuxt-vue
```

Next we can install all the dependencies using npm.

```bash
npm install
```

While the node modules are being installed we can setup the environment variables by copying the `.env.example` file to `.env`.

```bash
mv .env.example .env
```

When the node modules have been installed we're now ready to migrate the database.

```bash
npm run db:migrate:deploy
```

We're now ready to start the application.

```bash
npm run dev
```

You can now open the application in your browser by visiting [http://localhost:3000](http://localhost:3000).

## Testing

There are currently no tests available for this version of Custodian.
