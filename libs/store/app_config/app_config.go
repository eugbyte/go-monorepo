package appconfig

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	appconfig "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
	"github.com/pkg/errors"
)

type AppConfigServicer interface {
	GetConfig(ctx context.Context, keyName string, _label *string) (string, error)
}

type AppConfigService struct {
	client            *appconfig.KeyValuesClient
	resourceGroupName string
	configStoreName   string
}

// get the subId, resourceGroupName, configStoreName from the properties tab
// for the subId, ignore the prefix '/subscriptions/'
func NewAppConfig(subId string, resourceGroupName string, configStoreName string) AppConfigServicer {
	var appConfig AppConfigService = AppConfigService{
		resourceGroupName: resourceGroupName,
		configStoreName:   configStoreName,
	}
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create credentials"))
	}
	client, err := appconfig.NewKeyValuesClient(subId, cred, nil)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create client"))
	}
	appConfig.client = client
	return appConfig
}

func (ac AppConfigService) GetConfig(ctx context.Context, keyName string, _label *string) (string, error) {
	label := ""
	if _label != nil {
		label = *_label
	}

	keyLabel := keyName
	if len(label) >= 1 {
		// Key and label are joined by $ character. Label is optional
		// https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration#example-KeyValuesClient.Get
		keyLabel = fmt.Sprintf("%s$%s", keyName, label)
	}

	res, err := ac.client.Get(ctx,
		ac.resourceGroupName,
		ac.configStoreName,
		keyLabel,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}

	val := res.Properties.Value
	if val == nil {
		return "", errors.New("No such key")
	}
	return *val, nil
}
