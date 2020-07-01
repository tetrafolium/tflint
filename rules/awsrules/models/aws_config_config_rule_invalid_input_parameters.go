// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsConfigConfigRuleInvalidInputParametersRule checks the pattern is valid
type AwsConfigConfigRuleInvalidInputParametersRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigConfigRuleInvalidInputParametersRule returns new rule with default attributes
func NewAwsConfigConfigRuleInvalidInputParametersRule() *AwsConfigConfigRuleInvalidInputParametersRule {
	return &AwsConfigConfigRuleInvalidInputParametersRule{
		resourceType:  "aws_config_config_rule",
		attributeName: "input_parameters",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigConfigRuleInvalidInputParametersRule) Name() string {
	return "aws_config_config_rule_invalid_input_parameters"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigConfigRuleInvalidInputParametersRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigConfigRuleInvalidInputParametersRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigConfigRuleInvalidInputParametersRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigConfigRuleInvalidInputParametersRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"input_parameters must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"input_parameters must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
