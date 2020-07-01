// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsKmsKeyInvalidDescriptionRule checks the pattern is valid
type AwsKmsKeyInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsKmsKeyInvalidDescriptionRule returns new rule with default attributes
func NewAwsKmsKeyInvalidDescriptionRule() *AwsKmsKeyInvalidDescriptionRule {
	return &AwsKmsKeyInvalidDescriptionRule{
		resourceType:  "aws_kms_key",
		attributeName: "description",
		max:           8192,
	}
}

// Name returns the rule name
func (r *AwsKmsKeyInvalidDescriptionRule) Name() string {
	return "aws_kms_key_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKmsKeyInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKmsKeyInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKmsKeyInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKmsKeyInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 8192 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
