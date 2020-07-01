// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsCloudwatchEventRuleInvalidRoleArnRule checks the pattern is valid
type AwsCloudwatchEventRuleInvalidRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudwatchEventRuleInvalidRoleArnRule returns new rule with default attributes
func NewAwsCloudwatchEventRuleInvalidRoleArnRule() *AwsCloudwatchEventRuleInvalidRoleArnRule {
	return &AwsCloudwatchEventRuleInvalidRoleArnRule{
		resourceType:  "aws_cloudwatch_event_rule",
		attributeName: "role_arn",
		max:           1600,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventRuleInvalidRoleArnRule) Name() string {
	return "aws_cloudwatch_event_rule_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventRuleInvalidRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventRuleInvalidRoleArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventRuleInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventRuleInvalidRoleArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"role_arn must be 1600 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role_arn must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
