package models

import "encoding/json"

type RequestBody struct {
	Data     map[string]json.RawMessage
	Metadata map[string]interface{}
}

type ResponseBody struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue string
}
