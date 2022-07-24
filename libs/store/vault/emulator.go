package vault

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

/*
	Fake token used for bypassing the fake authentication of Lowkey Vault
*/
type FakeCredential struct{}

//goland:noinspection GoUnusedParameter
func (f *FakeCredential) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{Token: "faketoken", ExpiresOn: time.Now().Add(time.Hour).UTC()}, nil
}

/*
	Ignore SSL error caused by the self-signed certificate.
*/
func InsecureClient() http.Client {
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return http.Client{Transport: customTransport}
}

func NewMockVaultService(vaultURI string) VaultServicer {
	vs := vaultService{}
	stage := config.Stage()
	formats.Trace(stage)

	formats.Trace("Creating emulated vault...")
	httpClient := InsecureClient()
	vs.client = azsecrets.NewClient("https://localhost:8443",
		&FakeCredential{},
		&policy.ClientOptions{Transport: &httpClient})
	return &vs
}
