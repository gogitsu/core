package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// https://github.com/jmartin82/mconfig

var env string
var onceConfig sync.Once
var onceRead sync.Once
var instance interface{}

// Config loads the configuration structure singleton instance.
// It loads the configuration from file just one time at the first call.
// All next calls will return the cached configuration struct.
func Config(cfg interface{}) {
	onceConfig.Do(func() {
		ReConfig(&cfg)
	})

	cfg = instance
}

// ReConfig loads configuration from file each time is called.
// To be used in reload configuration case.
// Please, try and use only Config().
func ReConfig(cfg interface{}) {
	env = os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	loadConfig(&cfg)
	instance = cfg
}

//
// TODO: remove the panic() call and return an error... to be able to use default configurations.
//
func loadConfig(cfg interface{}) {
	ReadConfig()

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error decoding configuration into struct: %v", err))
	}
}

// ReadConfig read configuration from config file.
// Used by ReConfig() but it is also public in the case we want
// only load configuration into Viper and use Getxxxx() methods.
func ReadConfig() {
	onceRead.Do(func() {
		ReReadConfig()
	})
}

// ReReadConfig reload configuration from file.
func ReReadConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config-" + env)
	viper.SetConfigType("yml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.BindEnv("SERVICE_GROUP")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

// Get returns a configuration map by key.
func Get(key string) interface{} {
	return viper.Get(key)
}

// GetString returns a string by key.
func GetString(key string) string {
	return viper.GetString(key)
}

// GetStruct get configuration by key and decode it into the passed struct.
func GetStruct(cstruct interface{}) error {
	if err := viper.Unmarshal(&cstruct); err != nil {
		return err
	}
	return nil
}
