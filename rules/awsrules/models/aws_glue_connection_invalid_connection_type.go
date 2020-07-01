// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsGlueConnectionInvalidConnectionTypeRule checks the pattern is valid
type AwsGlueConnectionInvalidConnectionTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGlueConnectionInvalidConnectionTypeRule returns new rule with default attributes
func NewAwsGlueConnectionInvalidConnectionTypeRule() *AwsGlueConnectionInvalidConnectionTypeRule {
	return &AwsGlueConnectionInvalidConnectionTypeRule{
		resourceType:  "aws_glue_connection",
		attributeName: "connection_type",
		enum: []string{
			"JDBC",
			"SFTP",
			"MONGODB",
			"KAFKA",
		},
	}
}

// Name returns the rule name
func (r *AwsGlueConnectionInvalidConnectionTypeRule) Name() string {
	return "aws_glue_connection_invalid_connection_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlueConnectionInvalidConnectionTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlueConnectionInvalidConnectionTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlueConnectionInvalidConnectionTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlueConnectionInvalidConnectionTypeRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as connection_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
