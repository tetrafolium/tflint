// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsDynamoDBTableItemInvalidHashKeyRule checks the pattern is valid
type AwsDynamoDBTableItemInvalidHashKeyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsDynamoDBTableItemInvalidHashKeyRule returns new rule with default attributes
func NewAwsDynamoDBTableItemInvalidHashKeyRule() *AwsDynamoDBTableItemInvalidHashKeyRule {
	return &AwsDynamoDBTableItemInvalidHashKeyRule{
		resourceType:  "aws_dynamodb_table_item",
		attributeName: "hash_key",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsDynamoDBTableItemInvalidHashKeyRule) Name() string {
	return "aws_dynamodb_table_item_invalid_hash_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDynamoDBTableItemInvalidHashKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDynamoDBTableItemInvalidHashKeyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDynamoDBTableItemInvalidHashKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDynamoDBTableItemInvalidHashKeyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"hash_key must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"hash_key must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
