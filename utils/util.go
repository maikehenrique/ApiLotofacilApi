package util

import "encoding/json"

func StructToJson(v interface{}) *json.RawMessage {
	b, _ := json.Marshal(v)
	json := json.RawMessage(b)
	return &json
}
