package rule

import (
	"github.com/RossyWhite/tflint-ruleset-onename/config"
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

type Case struct {
	Name      string
	TfContent string
	Config    config.Config
	Expected  helper.Issues
}

func Test_Rule(t *testing.T) {
	cases := createTestCase()

	for _, tc := range cases {
		rule := NewOneNameRule()
		rule.conf = tc.Config

		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.TfContent})
		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}

func createTestCase() []Case {
	return []Case{
		{
			Name: "issue found",
			TfContent: `
resource "aws_instance" "web" {
    instance_type = "t2.micro"
}`,
			Config: &config.RuleConfig{
				Rules: []*config.Rule{
					{
						Resource: "aws_instance",
						Attribute: "instance_type",
						Regex: `[A-Z].*`,
					},
				},
			},
			Expected: helper.Issues{
				{
					Rule:    NewOneNameRule(),
					Message: "aws_instance.instance_type does not match pattern `[A-Z].*`",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 21},
						End:      hcl.Pos{Line: 3, Column: 31},
					},
				},
			},
		},

		{
			Name: "issue not found",
			TfContent: `
resource "aws_instance" "web" {
    instance_type = "t2.micro"
}`,
			Config: &config.RuleConfig{
				Rules: []*config.Rule{
					{
						Resource: "aws_instance",
						Attribute: "instance_type",
						Regex: `^t[1-9]\.[a-z].+$`,
					},
				},
			},
			Expected: helper.Issues(nil),
		},

		{
			Name: "only one issue",
			TfContent: `
resource "aws_instance" "web" {
    instance_type = "t2.micro"
}`,
			Config: &config.RuleConfig{
				Rules: []*config.Rule{
					{
						Resource: "aws_instance",
						Attribute: "instance_type",
						Regex: `[A-Z].*`,
					},
					{
						Resource: "aws_instance",
						Attribute: "instance_type",
						Regex: `^t[1-9]\.[a-z].+$`,
					},
				},
			},
			Expected: helper.Issues{
				{
					Rule:    NewOneNameRule(),
					Message: "aws_instance.instance_type does not match pattern `[A-Z].*`",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 21},
						End:      hcl.Pos{Line: 3, Column: 31},
					},
				},
			},
		},

		// TODO: consider to output error
		{
			Name: "not exist attribute",
			TfContent: `
resource "aws_instance" "web" {
    instance_type = "t2.micro"
}`,
			Config: &config.RuleConfig{
				Rules: []*config.Rule{
					{
						Resource: "aws_instance",
						Attribute: "hoge",
						Regex: `[A-Z].*`,
					},
				},
			},
			Expected: helper.Issues(nil),
		},
	}
}
