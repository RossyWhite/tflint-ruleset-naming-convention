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

type Config interface {
	Load(path string) error
	GetRules() []*Rule
}

// Config is plugin configuration
type RuleConfig struct {
	Rules []*Rule `validate:"required,dive,required",json:"rules"`
}

// Rule is a configuration of each rule
type Rule struct {
	Resource  string `validate:"required",json:"resource"`
	Attribute string `validate:"required",json:"attribute"`
	Regex     string `validate:"required",json:"regex"`
}

func NewRuleConfig() *RuleConfig {
	return &RuleConfig{}
}

func (c *RuleConfig) Load(path string) error {
	var fp string

	if path != "" {
		fp = path
	}

	fp, err := configPath()
	if err != nil {
		return err
	}

	f, err := os.Open(fp)
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

func (c *RuleConfig) GetRules() []*Rule {
	return c.Rules
}

func (c *RuleConfig) validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

func configPath() (string, error) {
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
