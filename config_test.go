package core

import (
	"testing"
)

// Cfg .
type Cfg struct {
	Service struct {
		Group   string `yaml:"group" env:"GROUP" env-default:"env-group"`
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"service"`
}

func TestReadConfig(t *testing.T) {
	c := &Configurator{}
	cfg := Cfg{}
	c.SetFileName("config-test.yml")
	if err := c.ReadConfig(&cfg); err != nil {
		t.Error(err.Error())
	}

	t.Logf("Configuration: %+v\n", cfg)
}
