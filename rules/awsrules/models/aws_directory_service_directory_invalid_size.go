// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsDirectoryServiceDirectoryInvalidSizeRule checks the pattern is valid
type AwsDirectoryServiceDirectoryInvalidSizeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDirectoryServiceDirectoryInvalidSizeRule returns new rule with default attributes
func NewAwsDirectoryServiceDirectoryInvalidSizeRule() *AwsDirectoryServiceDirectoryInvalidSizeRule {
	return &AwsDirectoryServiceDirectoryInvalidSizeRule{
		resourceType:  "aws_directory_service_directory",
		attributeName: "size",
		enum: []string{
			"Small",
			"Large",
		},
	}
}

// Name returns the rule name
func (r *AwsDirectoryServiceDirectoryInvalidSizeRule) Name() string {
	return "aws_directory_service_directory_invalid_size"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDirectoryServiceDirectoryInvalidSizeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDirectoryServiceDirectoryInvalidSizeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDirectoryServiceDirectoryInvalidSizeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDirectoryServiceDirectoryInvalidSizeRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as size`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
