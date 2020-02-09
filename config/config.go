package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

const (
	localConfigPath   = "./.tflint.d/configs/onename.json"
	defaultConfigPath = "~/.tflint.d/configs/onename.json"
)

// Config is plugin configuration
type Config struct {
	Rules []*Rule `validate:"required,dive,required",json:"rules"`
}

// Rule is a configuration of each rule
type Rule struct {
	Resource  string `validate:"required",json:"resource"`
	Attribute string `validate:"required",json:"attribute"`
	Regex     string `validate:"required",json:"regex"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfig() error {
	path, err := configPath()
	if err != nil {
		return errors.Wrap(err, "configPath failed")
	}

	f, err := os.Open(path)
	defer func() { _ = f.Close() }()

	if err != nil {
		return errors.New("os.Open failed")
	}

	if err := json.NewDecoder(f).Decode(c); err != nil {
		return errors.Wrap(err, "Decode failed")
	}

	if err := c.validate(); err != nil {
		return errors.Wrap(err, "Validate failed")
	}

	return nil
}

func (c *Config) validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

func configPath() (string, error) {
	if fileExists(localConfigPath) {
		return localConfigPath, nil
	}

	if fileExists(defaultConfigPath) {
		return defaultConfigPath, nil
	}

	return "", errors.New(fmt.Sprintf("%s or %s doesn't exsit", defaultConfigPath, localConfigPath))
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
