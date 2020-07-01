// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsIAMPolicyInvalidDescriptionRule checks the pattern is valid
type AwsIAMPolicyInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsIAMPolicyInvalidDescriptionRule returns new rule with default attributes
func NewAwsIAMPolicyInvalidDescriptionRule() *AwsIAMPolicyInvalidDescriptionRule {
	return &AwsIAMPolicyInvalidDescriptionRule{
		resourceType:  "aws_iam_policy",
		attributeName: "description",
		max:           1000,
	}
}

// Name returns the rule name
func (r *AwsIAMPolicyInvalidDescriptionRule) Name() string {
	return "aws_iam_policy_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMPolicyInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMPolicyInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMPolicyInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMPolicyInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 1000 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
