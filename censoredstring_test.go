// SPDX-FileCopyrightText: 2022 Coop Norge SA
//
// SPDX-License-Identifier: MIT

package masker

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v3"
)

type valueContainer struct {
	Value CensoredString `yaml:"Value"`
}

type CensoredStringTestSuite struct {
	suite.Suite
	clear     string
	clearJSON []byte
	masked    CensoredString
}

func (suite *CensoredStringTestSuite) SetupTest() {
	suite.clear = "V8a1fkz8etD4ntmH"
	suite.masked = CensoredString(suite.clear)

	clearJSON, err := json.Marshal(map[string]string{
		"Value": suite.clear,
	})
	assert.NoError(suite.T(), err)
	suite.T().Logf("m = %s", clearJSON)

	suite.clearJSON = clearJSON
}

// TestJSONUnmarshal tests that JSON can be unmarshalled into a CensoredString.
func (suite *CensoredStringTestSuite) TestUnmarshalJSON() {
	container := &valueContainer{}
	suite.Equal("", container.Value.UnmaskString())
	err := json.Unmarshal(suite.clearJSON, container)
	suite.NoError(err)
	suite.T().Logf("container = %#v", container)
	suite.Equal(suite.clear, container.Value.UnmaskString())
}

// TestMarshalJSON tests that a CensoredString is censored when it is marshalled to JSON.
func (suite *CensoredStringTestSuite) TestMarshalJSON() {
	container := &valueContainer{Value: suite.masked}
	pmj, err := json.Marshal(container)
	suite.NoError(err)
	suite.T().Logf("c = %s", pmj)
	suite.Equal(fmt.Sprintf("{\"Value\":\"%s\"}", CensoredText), string(pmj))
}

// TestStringFormat tests that a CensoredString is censored when formatted as a string.
func (suite *CensoredStringTestSuite) TestStringFormat() {
	suite.Equal(CensoredText, fmt.Sprintf("%s", suite.masked))
	suite.Equal(CensoredText, fmt.Sprintf("%s", CensoredString("")))
}

// TestUnmarshalYAML tests that YAML can be unmarshalled into a CensoredString.
func (suite *CensoredStringTestSuite) TestUnmarshalYAML() {
	container := &valueContainer{}
	suite.Equal("", container.Value.UnmaskString())
	err := yaml.Unmarshal(suite.clearJSON, container)
	suite.NoError(err)
	suite.T().Logf("container = %#v", container)
	suite.Equal(suite.clear, container.Value.UnmaskString())
}

// TestDefaultFormat tests that a CensoredString is censored when formatted with default format.
func (suite *CensoredStringTestSuite) TestDefaultFormat() {
	suite.Equal(CensoredText, fmt.Sprintf("%v", suite.masked))
	suite.Equal(CensoredText, fmt.Sprintf("%v", CensoredString("")))
}

// TestGoRepresentationFormat tests that a CensoredString is censored when formatted as go representation.
func (suite *CensoredStringTestSuite) TestGoRepresentationFormat() {
	suite.Equal(fmt.Sprintf("\"%s\"", CensoredText), fmt.Sprintf("%#v", suite.masked))
	suite.Equal(fmt.Sprintf("\"%s\"", CensoredText), fmt.Sprintf("%#v", CensoredString("")))
}

func (suite *CensoredStringTestSuite) TestStringer() {
	suite.Equal(CensoredText, suite.masked.String())
}

func (suite *CensoredStringTestSuite) TestGoStringer() {
	suite.Equal(fmt.Sprintf("\"%s\"", CensoredText), suite.masked.GoString())
}

func (suite *CensoredStringTestSuite) TestMarshalText() {
	text, err := suite.masked.MarshalText()
	suite.NoError(err)
	suite.Equal([]byte(CensoredText), text)
}

func TestCensoredStringTestSuite(t *testing.T) {
	suite.Run(t, new(CensoredStringTestSuite))
}
