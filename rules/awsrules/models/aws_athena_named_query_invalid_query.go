// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsAthenaNamedQueryInvalidQueryRule checks the pattern is valid
type AwsAthenaNamedQueryInvalidQueryRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAthenaNamedQueryInvalidQueryRule returns new rule with default attributes
func NewAwsAthenaNamedQueryInvalidQueryRule() *AwsAthenaNamedQueryInvalidQueryRule {
	return &AwsAthenaNamedQueryInvalidQueryRule{
		resourceType:  "aws_athena_named_query",
		attributeName: "query",
		max:           262144,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAthenaNamedQueryInvalidQueryRule) Name() string {
	return "aws_athena_named_query_invalid_query"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAthenaNamedQueryInvalidQueryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAthenaNamedQueryInvalidQueryRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAthenaNamedQueryInvalidQueryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAthenaNamedQueryInvalidQueryRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"query must be 262144 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"query must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
