package json

import (
	"encoding/json"
)

func ToJSONString(value interface{}) (string, error) {
	bytes, error := json.Marshal(value)

	return string(bytes), error
}

func FromJSONTo(value string, instance interface{}) error {
	return json.Unmarshal([]byte(value), instance)
}