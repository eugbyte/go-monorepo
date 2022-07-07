package models

import "encoding/json"

type RequestBody struct {
	Data     map[string]json.RawMessage
	Metadata map[string]interface{}
}

type Output struct {
	Message string `json:"message"`
	Res     struct {
		Body string `json:"body"`
	} `json:"res"`
}

type ResponseBody struct {
	Outputs     Output
	Logs        []string
	ReturnValue interface{}
}
