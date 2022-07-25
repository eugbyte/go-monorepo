package vault

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

type VaultServicer interface {
	GetSecret(secretName string) (string, error)
	SetSecret(secretName string, secretValue string) error
}

type vaultService struct {
	client *azsecrets.Client
}

func NewVaultService(vaultURI string) VaultServicer {
	vs := vaultService{}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to initialise vault service. %v", err)
	}

	// Establish a connection to the Key Vault client
	vs.client = azsecrets.NewClient(vaultURI, credential, nil)
	return &vs
}

func (vs *vaultService) GetSecret(secretName string) (string, error) {
	// Get a secret. An empty string version gets the latest version of the secret.
	version := ""
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	respn, err := vs.client.GetSecret(ctx, secretName, version, nil)
	if err != nil {
		return "", err
	}
	return *respn.Value, err
}

func (vs *vaultService) SetSecret(secretName string, secretValue string) error {
	if !formats.ValidateAzureParamString(secretName) {
		// https://docs.microsoft.com/en-us/azure/key-vault/secrets/quick-create-portal#add-a-secret-to-key-vault
		return errors.New("Secret name must only contain only 0-9, a-z, A-Z, and -. ")
	}

	params := azsecrets.SetSecretParameters{Value: &secretValue}
	_, err := vs.client.SetSecret(context.TODO(), secretName, params, nil)
	if err != nil {
		log.Printf("failed to create a secret: %v", err)
		return err
	}
	return nil
}
