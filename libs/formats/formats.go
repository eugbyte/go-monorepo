package formats

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	b64 "encoding/base64"

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

// Azure string must only contain only "0-9", "a-z", "A-Z", and "-", "."
// https://docs.microsoft.com/en-us/azure/key-vault/secrets/quick-create-portal#add-a-secret-to-key-vault
func ValidateAzureParamString(secretName string) bool {
	re, err := regexp.Compile(`^(?i)([a-z0-9\\-\\.])*$`)
	if err != nil {
		log.Fatalf("Could not compile regex expression: %v", err)
	}
	return re.MatchString(secretName)
}

func EncodeToBase64(str string) string {
	return b64.StdEncoding.EncodeToString([]byte(str))
}
