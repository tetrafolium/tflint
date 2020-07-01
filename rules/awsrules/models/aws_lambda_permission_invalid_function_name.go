// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsLambdaPermissionInvalidFunctionNameRule checks the pattern is valid
type AwsLambdaPermissionInvalidFunctionNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsLambdaPermissionInvalidFunctionNameRule returns new rule with default attributes
func NewAwsLambdaPermissionInvalidFunctionNameRule() *AwsLambdaPermissionInvalidFunctionNameRule {
	return &AwsLambdaPermissionInvalidFunctionNameRule{
		resourceType:  "aws_lambda_permission",
		attributeName: "function_name",
		max:           140,
		min:           1,
		pattern:       regexp.MustCompile(`^(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}(-gov)?-[a-z]+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\$LATEST|[a-zA-Z0-9-_]+))?$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaPermissionInvalidFunctionNameRule) Name() string {
	return "aws_lambda_permission_invalid_function_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaPermissionInvalidFunctionNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaPermissionInvalidFunctionNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaPermissionInvalidFunctionNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaPermissionInvalidFunctionNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"function_name must be 140 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"function_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}(-gov)?-[a-z]+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\$LATEST|[a-zA-Z0-9-_]+))?$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
