package ctors

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-siris/siris/core/errors"
	"github.com/spf13/viper"
)

const (
	envVarName    = "PUSHAPI_ENV"
	configVarName = "PUSHAPI_CONFIG"
)

func getEnvVariable() (string, error) {
	env := os.Getenv(envVarName)
	if env == "" {
		return "", errors.New(fmt.Sprintf("you forgot to pass the %s environment variable", envVarName))
	}
	fmt.Println("[config] env:", env)
	return env, nil
}

func setupFromDefaults(config *viper.Viper, env string) {
	config.Set("env", env)

	/*
		server
	*/
	config.SetDefault("api.enable_auth", true)

	/*
		mongodb
	*/
	config.SetDefault("mongodb.database", "push-api")

	/*
		server
	*/
	config.SetDefault("server.port", "8000")
}

func setupFromConfigurationFile(config *viper.Viper, env string) error {
	// try to use custom config file, or falls back to file corresponding to env
	filepath := os.Getenv(configVarName)
	if filepath == "" {
		filepath = fmt.Sprintf("./config/%s.yml", env)
	}

	config.SetConfigFile(filepath)
	if err := config.ReadInConfig(); err != nil {
		return errors.New(fmt.Sprintf("error loading config file: %s", filepath))
	}

	fmt.Println("[config] loaded config from file:", filepath)
	return nil
}

func setupFromEnvironment(config *viper.Viper) {
	replacer := strings.NewReplacer(".", "__")
	config.SetEnvKeyReplacer(replacer)
	config.SetEnvPrefix("pushapi")
	config.AutomaticEnv()
}


func NewViper() (*viper.Viper, error) {
	env, err := getEnvVariable()
	if err != nil {
		return nil, err
	}

	config := viper.New()
	setupFromDefaults(config, env)
	if err := setupFromConfigurationFile(config, env); err != nil {
		return nil, err
	}
	setupFromEnvironment(config)

	return config, nil
}
