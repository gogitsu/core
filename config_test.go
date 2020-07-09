package core

import (
	"testing"
)

/*

service:
  group: test-group
  name: test-service
  version: 0.1.0

*/

type Configuration struct {
	Service struct {
		Group   string `yaml:"group" env:"GROUP" env-default:"env-group"`
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"service"`
}

func TestReadConfig(t *testing.T) {
	c := &Configurator{}
	cfg := Configuration{}
	// c.AddPath("config")
	c.SetFileName("config-test.yml")
	if err := c.ReadConfig(&cfg); err != nil {
		t.Error(err.Error())
	}

	t.Logf("Configuration: %+v\n", cfg)
}
