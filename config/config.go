package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

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
func ReConfig(cfg interface{}) {
	env = os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	loadConfig(&cfg)
	instance = cfg
}

func loadConfig(cfg interface{}) {
	// viper.AddConfigPath(".")
	// viper.AddConfigPath("./config")
	// viper.SetConfigName("config-" + env)
	// viper.SetConfigType("yml")
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// viper.AutomaticEnv()
	// viper.BindEnv("SERVICE_GROUP")

	// if err := viper.ReadInConfig(); err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %s", err))
	// }

	ReadConfig()

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error decoding configuration into struct: %v", err))
	}
}

// ReadConfig read configuration from config file.
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
