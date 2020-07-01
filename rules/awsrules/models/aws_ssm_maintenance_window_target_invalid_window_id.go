// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsSsmMaintenanceWindowTargetInvalidWindowIDRule checks the pattern is valid
type AwsSsmMaintenanceWindowTargetInvalidWindowIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsmMaintenanceWindowTargetInvalidWindowIDRule returns new rule with default attributes
func NewAwsSsmMaintenanceWindowTargetInvalidWindowIDRule() *AwsSsmMaintenanceWindowTargetInvalidWindowIDRule {
	return &AwsSsmMaintenanceWindowTargetInvalidWindowIDRule{
		resourceType:  "aws_ssm_maintenance_window_target",
		attributeName: "window_id",
		max:           20,
		min:           20,
		pattern:       regexp.MustCompile(`^mw-[0-9a-f]{17}$`),
	}
}

// Name returns the rule name
func (r *AwsSsmMaintenanceWindowTargetInvalidWindowIDRule) Name() string {
	return "aws_ssm_maintenance_window_target_invalid_window_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmMaintenanceWindowTargetInvalidWindowIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmMaintenanceWindowTargetInvalidWindowIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmMaintenanceWindowTargetInvalidWindowIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmMaintenanceWindowTargetInvalidWindowIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"window_id must be 20 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"window_id must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^mw-[0-9a-f]{17}$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
