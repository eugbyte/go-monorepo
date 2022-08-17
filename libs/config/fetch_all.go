package config

import (
	"sync"

	"golang.org/x/sync/errgroup"
)

type secretConfig struct {
	mu      sync.Mutex
	secrets []string
}

type FetchVal = func(name string) (string, error)

func FetchAll(fetchVal FetchVal, secretNames ...string) ([]string, error) {
	var conf secretConfig = secretConfig{}
	conf.secrets = make([]string, 0, len(secretNames))

	grp := new(errgroup.Group)

	for i := 0; i < len(secretNames); i++ {
		// create a copy, otherwise i will just be the last value due to closure
		// variables declared by the init statement are re-used in each iteration.
		// this means that when the program is running, there's just a single object representing i,
		index := i
		grp.Go(func() error {
			conf.mu.Lock()
			defer conf.mu.Unlock()
			key, err := fetchVal(secretNames[index])
			conf.secrets[index] = key
			return err
		})
	}

	err := grp.Wait()
	return conf.secrets, err
}
