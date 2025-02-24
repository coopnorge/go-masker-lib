package masker_test

import (
	"encoding/json"
	"fmt"

	masker "github.com/coopnorge/go-masker-lib"
)

type valueContainer struct {
	Value masker.CensoredString
}

func ExampleCensoredString() {
	protectedValue := masker.CensoredString("secretvalue")
	fmt.Printf("%s", protectedValue)
	// Output: ###CENSORED###
}

func ExampleCensoredString_UnmaskString() {
	protectedValue := masker.CensoredString("secretvalue")
	fmt.Println(protectedValue.UnmaskString())
	// Output: secretvalue
}

func ExampleCensoredString_json_unmarshal() {
	data := valueContainer{}
	err := json.Unmarshal([]byte("{\"Value\":\"secretvalue\"}"), &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", data)
	// Output: {Value:###CENSORED###}
}

func ExampleCensoredString_json_marshal() {
	data := &valueContainer{
		Value: "secretvalue",
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bytes)
	// Output: {"Value":"###CENSORED###"}
}
