// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsLambdaPermissionInvalidSourceArnRule checks the pattern is valid
type AwsLambdaPermissionInvalidSourceArnRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLambdaPermissionInvalidSourceArnRule returns new rule with default attributes
func NewAwsLambdaPermissionInvalidSourceArnRule() *AwsLambdaPermissionInvalidSourceArnRule {
	return &AwsLambdaPermissionInvalidSourceArnRule{
		resourceType:  "aws_lambda_permission",
		attributeName: "source_arn",
		pattern:       regexp.MustCompile(`^arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-])+:([a-z]{2}(-gov)?-[a-z]+-\d{1})?:(\d{12})?:(.*)$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaPermissionInvalidSourceArnRule) Name() string {
	return "aws_lambda_permission_invalid_source_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaPermissionInvalidSourceArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaPermissionInvalidSourceArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaPermissionInvalidSourceArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaPermissionInvalidSourceArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-])+:([a-z]{2}(-gov)?-[a-z]+-\d{1})?:(\d{12})?:(.*)$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
