package tics

import (
	"encoding/json"
)

// Quick reusable json boiler, intended to be used on byte slices.
func ToJson(i interface{}) map[string]interface{} {

	var j map[string]interface{}

	b, ok := i.([]byte)
	if ok {
		err := json.Unmarshal(b, &j)
		if err != nil {
			ThrowSys(ToJson, err)
		}
	}

	return j
}
