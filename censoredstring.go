// SPDX-FileCopyrightText: 2022 Coop Norge SA
//
// SPDX-License-Identifier: MIT

package masker

import (
	"encoding"
	"fmt"
)

// CensoredText is the text displayed instead of protected values when they are masked.
const CensoredText = "###CENSORED###"

// CensoredString masks a string value by censoring it when it printed or marshalled to JSON, YAML or Text. The protected value can be revealed by calling the UnmaskString() method on the CensoredString value.
type CensoredString string

// MarshalText marshals the CensoredString value into a textual form, which will invariably be the value of the CensoredText constant.
func (s CensoredString) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

// String returns the native format for a CensoredString value, which will invariably be the CensoredString constant.
func (s CensoredString) String() string {
	return CensoredText
}

// GoString returns the Go syntax for a CensoredString value, which will invariably be the Go syntax for the CensoredString constant.
func (s CensoredString) GoString() string {
	return fmt.Sprintf("%#v", s.String())
}

// UnmaskString returns the underlying protected string of a CensoredString value.
func (s CensoredString) UnmaskString() string {
	return string(s)
}

var (
	_ encoding.TextMarshaler = (*CensoredString)(nil)
	_ fmt.Stringer           = (*CensoredString)(nil)
	_ fmt.GoStringer         = (*CensoredString)(nil)
	_ StringUnmasker         = (*CensoredString)(nil)
)
