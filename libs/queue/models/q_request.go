package models

import "encoding/json"

type RequestBody struct {
	Data     map[string]json.RawMessage
	Metadata map[string]interface{}
}
