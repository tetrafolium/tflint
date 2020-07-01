// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsOrganizationsPolicyInvalidNameRule checks the pattern is valid
type AwsOrganizationsPolicyInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsPolicyInvalidNameRule returns new rule with default attributes
func NewAwsOrganizationsPolicyInvalidNameRule() *AwsOrganizationsPolicyInvalidNameRule {
	return &AwsOrganizationsPolicyInvalidNameRule{
		resourceType:  "aws_organizations_policy",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\s\S]*$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsPolicyInvalidNameRule) Name() string {
	return "aws_organizations_policy_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsPolicyInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOrganizationsPolicyInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsPolicyInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsPolicyInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 128 characters or less",
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\s\S]*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
