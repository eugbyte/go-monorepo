package config

import (
	"sync"

	"golang.org/x/sync/errgroup"
)

type muConfig struct {
	mu     sync.Mutex
	values []string
}

type FetchVal = func(name string) (string, error)

func FetchAll(fetchVal FetchVal, secretNames ...string) ([]string, error) {
	var conf muConfig = muConfig{}
	conf.values = make([]string, len(secretNames))

	grp := new(errgroup.Group)

	for i := 0; i < len(secretNames); i++ {
		// create a copy, otherwise i will just be the last value due to closure
		// variables declared by the init statement are re-used in each iteration.
		// this means that when the program is running, there's just a single object representing i,
		index := i
		grp.Go(func() error {
			conf.mu.Lock()
			defer conf.mu.Unlock()
			value, err := fetchVal(secretNames[index])
			conf.values[index] = value
			return err
		})
	}

	err := grp.Wait()
	return conf.values, err
}
