// SPDX-FileCopyrightText: 2022 Coop Norge SA
//
// SPDX-License-Identifier: MIT

package example

import (
	"encoding/json"
	"testing"

	"github.com/coopnorge/go-masker-lib"
	"github.com/stretchr/testify/assert"
)

func TestCensoredStringExample(t *testing.T) {
	secretValue := "secretvalue"

	// masker.CensoredString can be used to protect sensitive or secret values.
	protectedValue := masker.CensoredString(secretValue)

	// The protected value will not appear in formatted output.
	assert.NotContains(t, protectedValue.String(), secretValue)

	// The protected value will not appear in marshalled output.
	m, err := json.Marshal(protectedValue)
	assert.NoError(t, err)
	assert.NotContains(t, m, secretValue)

	// The underlying secret value can be revealed.
	assert.Equal(t, protectedValue.UnmaskString(), secretValue)
}
