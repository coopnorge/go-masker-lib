// SPDX-FileCopyrightText: 2022 Coop Norge SA
//
// SPDX-License-Identifier: MIT

package masker

// StringUnmasker is an interface that provides a method for revealing masked strings.
type StringUnmasker interface {
	UnmaskString() string
}
