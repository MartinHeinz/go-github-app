# Template for GitHub Apps built with Golang

[![Build, Test and Lint Action](https://github.com/MartinHeinz/go-github-app/workflows/Build,%20Test,%20Lint/badge.svg)](https://github.com/MartinHeinz/go-github-app/workflows/Build,%20Test,%20Lint/badge.svg)
[![Release Action](https://github.com/MartinHeinz/go-github-app/workflows/Release/badge.svg)](https://github.com/https://github.com/MartinHeinz/go-github-app/workflows/Release/badge.svg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=MartinHeinz_go-github-app&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=MartinHeinz_go-github-app)
[![Maintainability](https://api.codeclimate.com/v1/badges/05a671e6cc9b25ddd1e5/maintainability)](https://codeclimate.com/github/MartinHeinz/go-github-app/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/05a671e6cc9b25ddd1e5/test_coverage)](https://codeclimate.com/github/MartinHeinz/go-github-app/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/MartinHeinz/go-github-app)](https://goreportcard.com/report/github.com/MartinHeinz/go-github-app)

## Blog Posts - More Information About This Repo

You can find more information about this project/repository and how to use it in following blog post:

[Building GitHub Apps with Golang](https://martinheinz.dev/blog/65)

## Quick Start

To use this repository as starter for your project you can run configure_project.sh script, which sets up all variables and file names. This way you can avoid configuring and renaming things yourself:

```bash
./configure_project.sh \
    APP_ID="54321" \
    INSTALLATION_ID="987654321" \
    WEBHOOK_SECRET="verysecret" \
    KEY_PATH="./github_key.pem" \
    REGISTRY="ghcr.io/<GITHUB_USERNAME>/go-github-app"
```

## Running

```bash
make container  # Builds containerized application
make run        # Runs container at localhost

# From another terminal:
curl http://localhost:8080/api/v1/github/pullrequests/octocat/hello-world
```

## Testing

Test are run inside container image, equivalent to the container in which the application runs. To run tests:

```bash
make test

Running tests:
?   	github.com/MartinHeinz/go-github-app/cmd/app	[no test files]
ok  	github.com/MartinHeinz/go-github-app/cmd/app/apis	0.010s
?   	github.com/MartinHeinz/go-github-app/cmd/app/config	[no test files]
?   	github.com/MartinHeinz/go-github-app/cmd/app/httputil	[no test files]
?   	github.com/MartinHeinz/go-github-app/cmd/app/test_data	[no test files]
?   	github.com/MartinHeinz/go-github-app/cmd/app/utils	[no test files]
ok  	github.com/MartinHeinz/go-github-app/cmd/app/webhooks	0.006s
?   	github.com/MartinHeinz/go-github-app/pkg	[no test files]

Checking gofmt: PASS

Checking go vet: PASS
```

## CI/CD

Predefined CI/CD uses GitHub Actions:

- _Build, Test, Lint_ Workflow (`build.yaml`):
    - Builds binary and container image
    - Runs tests and generates code coverage report
    - Performs SonarCloud code analysis
    - Sends coverage starts to CodeClimate

- _Release_ Workflow (`release.yaml`, triggered on _tag_ creation):
    - Builds container image
    - Pushes the image to GitHub container registry
