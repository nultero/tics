package tics

import (
	"encoding/json"
)

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
