## Template for GitHub Apps built with Golang

[![Build, Test and Lint Action](https://github.com/MartinHeinz/go-github-app/workflows/Build,%20Test,%20Lint/badge.svg)](https://github.com/MartinHeinz/go-github-app/workflows/Build,%20Test,%20Lint/badge.svg)
[![Release Action](https://github.com/MartinHeinz/go-github-app/workflows/Release/badge.svg)](https://github.com/https://github.com/MartinHeinz/go-github-app/workflows/Release/badge.svg)
[![Maintainability](https://api.codeclimate.com/v1/badges/05a671e6cc9b25ddd1e5/maintainability)](https://codeclimate.com/github/MartinHeinz/go-github-app/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/05a671e6cc9b25ddd1e5/test_coverage)](https://codeclimate.com/github/MartinHeinz/go-github-app/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/MartinHeinz/go-github-app)](https://goreportcard.com/report/github.com/MartinHeinz/go-github-app)

### Running

```bash
docker run --rm --name go-github-app \
    -v $(pwd)/config:/config \
    -p 8080:8080 \
    ghcr.io/martinheinz/go-github-app/app
```