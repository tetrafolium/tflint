// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsIAMUserLoginProfileInvalidUserRule checks the pattern is valid
type AwsIAMUserLoginProfileInvalidUserRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMUserLoginProfileInvalidUserRule returns new rule with default attributes
func NewAwsIAMUserLoginProfileInvalidUserRule() *AwsIAMUserLoginProfileInvalidUserRule {
	return &AwsIAMUserLoginProfileInvalidUserRule{
		resourceType:  "aws_iam_user_login_profile",
		attributeName: "user",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMUserLoginProfileInvalidUserRule) Name() string {
	return "aws_iam_user_login_profile_invalid_user"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMUserLoginProfileInvalidUserRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMUserLoginProfileInvalidUserRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMUserLoginProfileInvalidUserRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMUserLoginProfileInvalidUserRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"user must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"user must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w+=,.@-]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
