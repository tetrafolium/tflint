// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsAppsyncFunctionInvalidDataSourceRule checks the pattern is valid
type AwsAppsyncFunctionInvalidDataSourceRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAppsyncFunctionInvalidDataSourceRule returns new rule with default attributes
func NewAwsAppsyncFunctionInvalidDataSourceRule() *AwsAppsyncFunctionInvalidDataSourceRule {
	return &AwsAppsyncFunctionInvalidDataSourceRule{
		resourceType:  "aws_appsync_function",
		attributeName: "data_source",
		max:           65536,
		min:           1,
		pattern:       regexp.MustCompile(`^[_A-Za-z][_0-9A-Za-z]*$`),
	}
}

// Name returns the rule name
func (r *AwsAppsyncFunctionInvalidDataSourceRule) Name() string {
	return "aws_appsync_function_invalid_data_source"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppsyncFunctionInvalidDataSourceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppsyncFunctionInvalidDataSourceRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppsyncFunctionInvalidDataSourceRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppsyncFunctionInvalidDataSourceRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"data_source must be 65536 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"data_source must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[_A-Za-z][_0-9A-Za-z]*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
