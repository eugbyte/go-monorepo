package vault

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type VaultServiceImpl interface {
	Init(vaultURI string)
	GetSecret(secretName string) (string, error)
	SetSecret(secretName string, secretValue string) error
}

type VaultService struct {
	client *azsecrets.Client
}

func (vs VaultService) Init(vaultURI string) {
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to initialise vault service. %v", err)
	}

	// Establish a connection to the Key Vault client
	vs.client = azsecrets.NewClient(vaultURI, credential, nil)
}

func (vs VaultService) GetSecret(secretName string) (string, error) {
	// Get a secret. An empty string version gets the latest version of the secret.
	version := ""
	respn, err := vs.client.GetSecret(context.TODO(), secretName, version, nil)
	if err != nil {
		return "", err
	}
	return *respn.Value, err
}

func (vs VaultService) SetSecret(secretName string, secretValue string) error {
	params := azsecrets.SetSecretParameters{Value: &secretValue}
	_, err := vs.client.SetSecret(context.TODO(), secretName, params, nil)
	if err != nil {
		log.Fatalf("failed to create a secret: %v", err)
	}
	return nil
}
