package formats

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	colors "github.com/TwinProduction/go-color"
)

func Trace(objs ...interface{}) {
	var strs []string
	for _, obj := range objs {
		strs = append(strs, Stringify(obj))
	}
	fmt.Println(colors.Green, strings.Join(strs, " "), colors.Reset)
}

// Convert objects to string
func Stringify(obj interface{}) string {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		log.Panicln(err)
	}

	return string(objBytes)
}

func FormatURL(baseURL string, pathVariables []string, queryParams map[string]string) string {
	url := baseURL
	if len(pathVariables) > 0 {
		url += "/" + strings.Join(pathVariables, "/")
	}

	if len(queryParams) > 0 {
		var strs []string
		for key, val := range queryParams {
			strs = append(strs, fmt.Sprintf("%s=%s", key, val))
		}
		url += "?" + strings.Join(strs, "&")
	}

	return url
}
