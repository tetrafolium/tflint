// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsFsxWindowsFileSystemInvalidKmsKeyIDRule checks the pattern is valid
type AwsFsxWindowsFileSystemInvalidKmsKeyIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFsxWindowsFileSystemInvalidKmsKeyIDRule returns new rule with default attributes
func NewAwsFsxWindowsFileSystemInvalidKmsKeyIDRule() *AwsFsxWindowsFileSystemInvalidKmsKeyIDRule {
	return &AwsFsxWindowsFileSystemInvalidKmsKeyIDRule{
		resourceType:  "aws_fsx_windows_file_system",
		attributeName: "kms_key_id",
		max:           2048,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}|arn:aws[a-z-]{0,7}:kms:[a-z]{2}-[a-z-]{4,}-\d+:\d{12}:(key|alias)\/([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}|[a-zA-Z0-9:\/_-]+)|alias\/[a-zA-Z0-9:\/_-]+$`),
	}
}

// Name returns the rule name
func (r *AwsFsxWindowsFileSystemInvalidKmsKeyIDRule) Name() string {
	return "aws_fsx_windows_file_system_invalid_kms_key_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxWindowsFileSystemInvalidKmsKeyIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxWindowsFileSystemInvalidKmsKeyIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxWindowsFileSystemInvalidKmsKeyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxWindowsFileSystemInvalidKmsKeyIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"kms_key_id must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"kms_key_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}|arn:aws[a-z-]{0,7}:kms:[a-z]{2}-[a-z-]{4,}-\d+:\d{12}:(key|alias)\/([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}|[a-zA-Z0-9:\/_-]+)|alias\/[a-zA-Z0-9:\/_-]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}