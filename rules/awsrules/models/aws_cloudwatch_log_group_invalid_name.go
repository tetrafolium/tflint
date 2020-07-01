// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsCloudwatchLogGroupInvalidNameRule checks the pattern is valid
type AwsCloudwatchLogGroupInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchLogGroupInvalidNameRule returns new rule with default attributes
func NewAwsCloudwatchLogGroupInvalidNameRule() *AwsCloudwatchLogGroupInvalidNameRule {
	return &AwsCloudwatchLogGroupInvalidNameRule{
		resourceType:  "aws_cloudwatch_log_group",
		attributeName: "name",
		max:           512,
		min:           1,
		pattern:       regexp.MustCompile(`^[\.\-_/#A-Za-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchLogGroupInvalidNameRule) Name() string {
	return "aws_cloudwatch_log_group_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchLogGroupInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchLogGroupInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchLogGroupInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchLogGroupInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 512 characters or less",
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\.\-_/#A-Za-z0-9]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
