// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"


type Config struct {
	Period time.Duration `config:"period"`
	Facter *string
}

var DefaultConfig = Config{
	Period: 600 * time.Second,
	Facter: &FACTER_DEFAULT_PATH,
}
