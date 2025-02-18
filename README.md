# Masker library for Go

![branch status](https://github.com/coopnorge/go-masker-lib/actions/workflows/ci.yaml/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/coopnorge/go-masker-lib)](https://goreportcard.com/report/github.com/coopnorge/go-masker-lib)
[![codecov](https://codecov.io/gh/coopnorge/go-masker-lib/branch/main/graph/badge.svg)](https://codecov.io/gh/coopnorge/go-masker-lib)
[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/coopnorge/go-masker-lib)
[![license](https://img.shields.io/badge/license-MIT-green)](./LICENSE)

## Documentation

See [./docs/index.md](./docs/index.md)

## Developing

```bash
# build images
docker-compose build
# see available targets
docker-compose run --rm golang-devtools make help
# validate
docker-compose run --rm golang-devtools make validate VERBOSE=all
# run in watch mode
docker-compose run --rm golang-devtools make watch
```
