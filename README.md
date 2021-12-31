## Template for GitHub Apps built with Golang

[![Build, Test and Lint Action](https://github.com/MartinHeinz/go-github-app/workflows/Build,%20Test,%20Lint/badge.svg)](https://github.com/MartinHeinz/go-github-app/workflows/Build,%20Test,%20Lint/badge.svg)
[![Release Action](https://github.com/MartinHeinz/go-github-app/workflows/Release/badge.svg)](https://github.com/https://github.com/MartinHeinz/go-github-app/workflows/Release/badge.svg)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=MartinHeinz_go-github-app&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=MartinHeinz_go-github-app)
[![Test Coverage](https://api.codeclimate.com/v1/badges/ec7ebefe63609984cb5c/test_coverage)](https://codeclimate.com/github/MartinHeinz/go-github-app/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/MartinHeinz/go-github-app)](https://goreportcard.com/report/github.com/MartinHeinz/go-github-app)

### Running

```bash
docker run --rm --name go-github-app \
    -v $(pwd)/config:/config \
    -p 8080:8080 \
    ghcr.io/martinheinz/go-github-app/app
```