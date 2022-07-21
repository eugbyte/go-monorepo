package app_config

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

type AppConfigService struct {
	baseUrl string
}

func (ac *AppConfigService) Init(subscriptionId string, resourceGroupName string, name string) {
	ac.baseUrl = fmt.Sprintf(
		`https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/config/web`,
		subscriptionId,
		resourceGroupName,
		name)
}

// key - key name
// label and api version are optional
func (ac *AppConfigService) GetValue(key string, label *string, apiVersion *string) (string, error) {
	pathVariables := []string{"kv", key}
	queryParams := map[string]string{}

	if label != nil {
		queryParams["label"] = *label
	}
	if apiVersion != nil {
		queryParams["apiVersion"] = *apiVersion
	}

	url := formats.FormatURL(ac.baseUrl, pathVariables, queryParams)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	formats.Trace(result)

	return result["value"].(string), nil
}
