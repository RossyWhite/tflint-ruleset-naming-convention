package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

const (
	LocalConfigPath   = "./.tflint.d/configs/onename.json"
	DefaultConfigPath = "~/.tflint.d/configs/onename.json"
)

// Config is plugin configuration
type Config struct {
	Rules []*Rule `validate:"required,dive,required",json:"rules"`
	path  string
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

func (c *Config) LoadConfig(path string) error {
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


func Path() (string, error) {
	if fileExists(LocalConfigPath) {
		return LocalConfigPath, nil
	}

	if fileExists(DefaultConfigPath) {
		return DefaultConfigPath, nil
	}

	return "", errors.New(fmt.Sprintf("%s or %s doesn't exsit",
		DefaultConfigPath, LocalConfigPath))
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func (c *Config) validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
