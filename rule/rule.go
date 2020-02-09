package rule

import (
	"fmt"
	"github.com/RossyWhite/tflint-ruleset-onename/config"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/pkg/errors"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// OneNameRule checks whether the resource comply with the specific naming convention
type OneNameRule struct{}

// NewOneNameRule returns a new rule
func NewOneNameRule() *OneNameRule {
	return &OneNameRule{}
}

// Name returns the rule name
func (r *OneNameRule) Name() string {
	return "one_name"
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
	return "https://github.com/RossyWhite/tflint-ruleset-onenmae"
}

// Check checks whether each attribute satisfy the condition given by config file
func (r *OneNameRule) Check(runner tflint.Runner) error {
	conf := config.NewConfig()
	err := conf.LoadConfig()

	if err != nil {
		return errors.Wrap(err, "loadConfig failed")
	}

	for _, rule := range conf.Rules {
		err = runner.WalkResourceAttributes(rule.Resource, rule.Attribute, func(attribute *hcl.Attribute) error {
			var val string
			err := runner.EvaluateExpr(attribute.Expr, &val)


			return runner.EnsureNoError(err, func() error {
				return runner.EmitIssue(
					r,
					fmt.Sprintf("%v", attribute.Expr),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			})
		})
	}

	return err
}
