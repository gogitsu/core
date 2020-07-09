package core

import (
	"errors"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	defaultPath     = "config"
	defaultFileName = "config.yml"
)

var (
	// ErrNoConfigFile returned if no config file has been found.
	ErrNoConfigFile = errors.New("Unable to find configuraiton file in no one of the specified path")

	onceConfig sync.Once
	instance   *Configurator
)

// Configurator is the main configuration struct.
// It olds configuration paths and file to search for.
type Configurator struct {
	paths    []string
	fileName string
}

// GetConfigurator .
func GetConfigurator() *Configurator {
	onceConfig.Do(func() {
		instance = &Configurator{
			paths:    []string{defaultPath},
			fileName: defaultFileName,
		}
	})
	return instance
}

// AddPath adds a path to the configuration paths array.
func (c *Configurator) AddPath(path string) {
	c.paths = append(c.paths, path)
}

// SetFileName sets the name for the configuration file.
func (c *Configurator) SetFileName(fileName string) {
	if fileName != "" {
		c.fileName = fileName
	}
}

// ReadConfig loads in the configuration from file into the cfg struct pointer.
func (c *Configurator) ReadConfig(cfg interface{}) error {
	if len(c.paths) == 0 {
		c.paths = append(c.paths, defaultPath)
	}
	for _, p := range c.paths {
		if err := cleanenv.ReadConfig(p+"/"+c.fileName, cfg); err == nil {
			return nil
		}
	}
	return ErrNoConfigFile
}
