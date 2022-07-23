package vault

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

type VaultServicer interface {
	GetSecret(secretName string) (string, error)
	SetSecret(secretName string, secretValue string) error
}

type VaultService struct {
	client *azsecrets.Client
}

func NewVaultService(vaultURI string) VaultServicer {
	vs := VaultService{}
	stage := config.Stage()
	formats.Trace(stage)

	if stage == config.DEV {
		httpClient := InsecureClient()
		vs.client = azsecrets.NewClient("https://localhost:8443",
			&FakeCredential{},
			&policy.ClientOptions{Transport: &httpClient})
		return &vs
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to initialise vault service. %v", err)
	}

	// Establish a connection to the Key Vault client
	vs.client = azsecrets.NewClient(vaultURI, credential, nil)
	return &vs
}

func (vs *VaultService) GetSecret(secretName string) (string, error) {
	// Get a secret. An empty string version gets the latest version of the secret.
	version := ""
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	respn, err := vs.client.GetSecret(ctx, secretName, version, nil)
	formats.Trace(err)
	if err != nil {
		return "", err
	}
	formats.Trace(respn)
	return *respn.Value, err
}

func (vs *VaultService) SetSecret(secretName string, secretValue string) error {
	params := azsecrets.SetSecretParameters{Value: &secretValue}
	_, err := vs.client.SetSecret(context.TODO(), secretName, params, nil)
	if err != nil {
		log.Fatalf("failed to create a secret: %v", err)
	}
	return nil
}
