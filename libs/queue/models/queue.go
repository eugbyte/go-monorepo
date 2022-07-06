package models

import "encoding/json"

type RequestBody struct {
	Data     map[string]json.RawMessage
	Metadata map[string]interface{}
}

type QueueData struct {
	Url     string
	Method  string
	Query   map[string]string
	Headers struct {
		ContentType []string `json:"Content-Type"`
	}
	Params map[string]string
	Body   interface{}
}

type Output struct {
	Message  string `json:"message"`
	Response struct {
		Body string `json:"body"`
	} `json:"response"`
}

type ResponseBody struct {
	Outputs     Output
	Logs        []string
	ReturnValue interface{}
}
