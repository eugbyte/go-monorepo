package lib

import (
	"context"
	"sync"

	appConfig "github.com/web-notify/api/monorepo/libs/store/app_config"
	"github.com/web-notify/api/monorepo/libs/store/vault"
	"golang.org/x/sync/errgroup"
)

type vapidConfig struct {
	mu         sync.Mutex
	PrivateKey string
	PublicKey  string
	Email      string
}

func FetchVapidConfig(vaultService vault.VaultServicer, appConfigService appConfig.AppConfigServicer) (*vapidConfig, error) {
	var conf vapidConfig = vapidConfig{}

	grp := new(errgroup.Group)

	grp.Go(func() error {
		conf.mu.Lock()
		defer conf.mu.Unlock()
		key, err := vaultService.GetSecret("VAPID-PRIVATE-KEY")
		conf.PrivateKey = key
		return err
	})

	grp.Go(func() error {
		conf.mu.Lock()
		defer conf.mu.Unlock()
		key, err := appConfigService.GetConfig(context.TODO(), "VAPID-PUBLIC-KEY", nil)
		conf.PublicKey = key
		return err
	})

	grp.Go(func() error {
		conf.mu.Lock()
		defer conf.mu.Unlock()
		email, err := appConfigService.GetConfig(context.TODO(), "VAPID-SENDER-EMAIL", nil)
		conf.Email = email
		return err
	})

	err := grp.Wait()
	return &conf, err
}
