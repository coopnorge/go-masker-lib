# Masker library for Go


![branch status](https://github.com/coopnorge/go-masker-lib/actions/workflows/validate.yml/badge.svg?branch=main)


[![Go Report Card](https://goreportcard.com/badge/github.com/coopnorge/go-masker-lib)](https://goreportcard.com/report/github.com/coopnorge/go-masker-lib)
[![codecov](https://codecov.io/gh/coopnorge/go-masker-lib/branch/master/graph/badge.svg)](https://codecov.io/gh/coopnorge/go-masker-lib)
[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/coopnorge/go-masker-lib)
[![license](https://img.shields.io/badge/license-MIT-green)](./LICENSE)

This library provides Go types which makes it easy to protect sensitive data from exposure by masking the data in situations that would result in exposure such as string formatting, logging and marshalling to JSON or YAML.

|Type|Protection Strategy|
|--|--|
|CensoredString|Protects the a value by masking it with a constant value.|

## Using CensoredString

The code below provides an example of how to use `CensoredString`:

```golang
package example

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/coopnorge/go-masker-lib"
	"github.com/stretchr/testify/assert"
)

func TestCensoredStringExample(t *testing.T) {
	secretValue := "secretvalue"

	// masker.CensoredString can be used to protect sensitive or secret values.
	protectedValue := masker.CensoredString(secretValue)

	// The protected value will not appear in formatted output.
	assert.NotContains(t, fmt.Sprintf("%s", protectedValue), secretValue)

	// The protected value will not appear in marshalled output.
	m, err := json.Marshal(protectedValue)
	assert.NoError(t, err)
	assert.NotContains(t, m, secretValue)

	// The underlying secret value can be revealed.
	assert.Equal(t, fmt.Sprintf("%s", protectedValue.UnmaskString()), secretValue)
}
```

## Developing


```bash
# build images
docker-compose build
# see available targets
docker-compose run --rm devtools make help
# validate
docker-compose run --rm devtools make validate VERBOSE=all
# run in watch mode
docker-compose run --rm devtools make watch
```
