// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsCodedeployAppInvalidNameRule checks the pattern is valid
type AwsCodedeployAppInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCodedeployAppInvalidNameRule returns new rule with default attributes
func NewAwsCodedeployAppInvalidNameRule() *AwsCodedeployAppInvalidNameRule {
	return &AwsCodedeployAppInvalidNameRule{
		resourceType:  "aws_codedeploy_app",
		attributeName: "name",
		max:           100,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCodedeployAppInvalidNameRule) Name() string {
	return "aws_codedeploy_app_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodedeployAppInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodedeployAppInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodedeployAppInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodedeployAppInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
