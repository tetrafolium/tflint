// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule checks the pattern is valid
type AwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule returns new rule with default attributes
func NewAwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule() *AwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule {
	return &AwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule{
		resourceType:  "aws_storagegateway_smb_file_share",
		attributeName: "default_storage_class",
		max:           50,
		min:           5,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule) Name() string {
	return "aws_storagegateway_smb_file_share_invalid_default_storage_class"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"default_storage_class must be 50 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"default_storage_class must be 5 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
