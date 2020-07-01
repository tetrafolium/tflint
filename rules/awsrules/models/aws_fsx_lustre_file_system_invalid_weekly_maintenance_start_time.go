// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule checks the pattern is valid
type AwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule returns new rule with default attributes
func NewAwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule() *AwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule {
	return &AwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule{
		resourceType:  "aws_fsx_lustre_file_system",
		attributeName: "weekly_maintenance_start_time",
		max:           7,
		min:           7,
		pattern:       regexp.MustCompile(`^[1-7]:([01]\d|2[0-3]):?([0-5]\d)$`),
	}
}

// Name returns the rule name
func (r *AwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule) Name() string {
	return "aws_fsx_lustre_file_system_invalid_weekly_maintenance_start_time"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"weekly_maintenance_start_time must be 7 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"weekly_maintenance_start_time must be 7 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[1-7]:([01]\d|2[0-3]):?([0-5]\d)$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
