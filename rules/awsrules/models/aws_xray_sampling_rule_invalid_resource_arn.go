// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsXraySamplingRuleInvalidResourceArnRule checks the pattern is valid
type AwsXraySamplingRuleInvalidResourceArnRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsXraySamplingRuleInvalidResourceArnRule returns new rule with default attributes
func NewAwsXraySamplingRuleInvalidResourceArnRule() *AwsXraySamplingRuleInvalidResourceArnRule {
	return &AwsXraySamplingRuleInvalidResourceArnRule{
		resourceType:  "aws_xray_sampling_rule",
		attributeName: "resource_arn",
		max:           500,
	}
}

// Name returns the rule name
func (r *AwsXraySamplingRuleInvalidResourceArnRule) Name() string {
	return "aws_xray_sampling_rule_invalid_resource_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsXraySamplingRuleInvalidResourceArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsXraySamplingRuleInvalidResourceArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsXraySamplingRuleInvalidResourceArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsXraySamplingRuleInvalidResourceArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"resource_arn must be 500 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
