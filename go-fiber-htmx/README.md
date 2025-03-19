# Go + Fiber + HTMX

This is a Fiber v2 and HTMX v2 version of Custodian running on Go 1.24.

## Table of Contents

- [The Stack](#the-stack)
- [Features](#features)
- [Installation](#installation)
    - [Docker](#docker)
    - [Local Environment](#local-environment)
- [Testing](#testing)

## The Stack

The backend is built with [Fiber 2](https://gofiber.io/) with [Jet](https://github.com/CloudyKit/jet) as the templating engine when rendering views, along with [GORM](https://gorm.io/) for connection and managing the database, and [Go Playground Validator](https://github.com/go-playground/validator) for validating requests.

For the frontend we're using [TailwindCSS 3](https://tailwindcss.com/) along with [DaisyUI](https://daisyui.com/) from the [template files](/template/), as well as [HTMX 2](https://htmx.org/) with [boosting enabled](https://htmx.org/docs/#boosting), allowing for a SPA like feel to the site, we're also using some [vanilla JS](http://vanilla-js.com/) to get custom behavior to work. The HTMX and JS is served over CDNs or in the HTML directly, while the CSS is built using [NodeJS v22](https://nodejs.org/).

The project is intended to be run within a Docker environment in production, and takes up around 23 MB of disk space when built, and around ~14 MB of RAM when running.

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
cd custodian/go-fiber-htmx
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
cd custodian/go-fiber-htmx
```

Next we can install all the dependencies using npm.

```bash
npm install
```

While the node modules are being installed we can setup the environment variables by copying the `.env.example` file to `.env`.

```bash
cp .env.example .env
```

When the node modules have been installed we're now ready to build the assets.

```bash
# Builds the assets for production
npm run build:prod
# Starts a watcher and builds the assets for development
npm run build:dev
```

We're now ready to start the application.

```bash
go run main.go
```

Alternatively you can use [Air](https://github.com/air-verse/air) to run the application, this will automatically restart the application when changes are made to the files.

If Air is installed you can also run the `serve` npm script to start `air` and `npm run build:dev` at the same time.

```bash
# Starts the application with Air
air
# Starts the application with Air and builds the assets for development
npm run serve
```

You can now open the application in your browser by visiting [http://localhost:8000](http://localhost:8000).

## Testing

The project uses the built-in testing framework that comes with Go, and all the tests are located in the `*_test.go` files in the project. To run the test suite first ensure that you have the project [installed and it's up and running](#installation), then you can run the following command to run the tests.

```bash
go test ./...
```
