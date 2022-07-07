package models

type ResponseBody struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue string
}
