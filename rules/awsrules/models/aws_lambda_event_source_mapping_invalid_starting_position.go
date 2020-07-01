// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsLambdaEventSourceMappingInvalidStartingPositionRule checks the pattern is valid
type AwsLambdaEventSourceMappingInvalidStartingPositionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsLambdaEventSourceMappingInvalidStartingPositionRule returns new rule with default attributes
func NewAwsLambdaEventSourceMappingInvalidStartingPositionRule() *AwsLambdaEventSourceMappingInvalidStartingPositionRule {
	return &AwsLambdaEventSourceMappingInvalidStartingPositionRule{
		resourceType:  "aws_lambda_event_source_mapping",
		attributeName: "starting_position",
		enum: []string{
			"TRIM_HORIZON",
			"LATEST",
			"AT_TIMESTAMP",
		},
	}
}

// Name returns the rule name
func (r *AwsLambdaEventSourceMappingInvalidStartingPositionRule) Name() string {
	return "aws_lambda_event_source_mapping_invalid_starting_position"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaEventSourceMappingInvalidStartingPositionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaEventSourceMappingInvalidStartingPositionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaEventSourceMappingInvalidStartingPositionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaEventSourceMappingInvalidStartingPositionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as starting_position`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
