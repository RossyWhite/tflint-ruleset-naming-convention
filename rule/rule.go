package rule

import (
	"fmt"
	"github.com/RossyWhite/tflint-ruleset-onename/config"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/pkg/errors"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"regexp"
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

				return runner.EmitIssue(
					r,
					fmt.Sprintf("%s.%s does match pattern `%s`",
						rule.Resource, rule.Attribute, rule.Regex),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			})
		})

		//if err != nil {
		//	break
		//}
	}

	return err
}

func checkRegexp(reg, str string) error {
	r, err := regexp.Compile(reg)

	if err != nil {
		return err
	}

	if ok := r.MatchString(str); !ok {
		return errors.New("Regex Not Matched")
	}

	return nil
}
