package rule

import (
	"fmt"
	"github.com/RossyWhite/tflint-ruleset-naming-convention/config"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/pkg/errors"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"regexp"
)


// OneNameRule checks whether the resource comply with the specific naming convention
type OneNameRule struct{
	conf config.Config
}

// NewOneNameRule returns a new rule
func NewOneNameRule() *OneNameRule {
	return &OneNameRule{}
}

// Name returns the rule name
func (r *OneNameRule) Name() string {
	return "naming_convention"
}

// Enabled returns whether the rule is enabled by default
func (r *OneNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *OneNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *OneNameRule) Link() string {
	return "https://github.com/RossyWhite/tflint-ruleset-naming-convention"
}

// Check checks whether each attribute satisfy the condition given by config file
func (r *OneNameRule) Check(runner tflint.Runner) error {
	conf, err := r.loadConfig()
	r.conf = conf

	if err != nil {
		return errors.Wrap(err, "loadConfig failed")
	}

	for _, rule := range r.conf.GetRules() {
		err = runner.WalkResourceAttributes(rule.Resource, rule.Attribute, func(attribute *hcl.Attribute) error {
			var val string
			if err := runner.EvaluateExpr(attribute.Expr, &val); err != nil {
				return err
			}

			matched, err := regexp.MatchString(rule.Regex, val)
			return runner.EnsureNoError(err, func() error {
				if !matched {
					return runner.EmitIssue(
						r,
						fmt.Sprintf("%s.%s does not match pattern `%s`",
							rule.Resource, rule.Attribute, rule.Regex),
						attribute.Expr.Range(),
						tflint.Metadata{Expr: attribute.Expr},
					)
				}
				return nil
			})
		})

		if err != nil {
			return err
		}
	}

	return nil
}

// loadConfig loads the configuration from config file
func (r *OneNameRule) loadConfig() (config.Config, error) {
	if r.conf != nil {
		return r.conf, nil
	}
	conf := config.NewRuleConfig()
	err := conf.Load("")
	if err != nil {
		return nil, err
	}

	return conf, nil
}

