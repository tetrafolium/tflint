// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsIAMRoleInvalidPermissionsBoundaryRule checks the pattern is valid
type AwsIAMRoleInvalidPermissionsBoundaryRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsIAMRoleInvalidPermissionsBoundaryRule returns new rule with default attributes
func NewAwsIAMRoleInvalidPermissionsBoundaryRule() *AwsIAMRoleInvalidPermissionsBoundaryRule {
	return &AwsIAMRoleInvalidPermissionsBoundaryRule{
		resourceType:  "aws_iam_role",
		attributeName: "permissions_boundary",
		max:           2048,
		min:           20,
	}
}

// Name returns the rule name
func (r *AwsIAMRoleInvalidPermissionsBoundaryRule) Name() string {
	return "aws_iam_role_invalid_permissions_boundary"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMRoleInvalidPermissionsBoundaryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMRoleInvalidPermissionsBoundaryRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMRoleInvalidPermissionsBoundaryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMRoleInvalidPermissionsBoundaryRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"permissions_boundary must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"permissions_boundary must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
