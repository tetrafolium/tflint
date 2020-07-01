// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsBatchJobQueueInvalidStateRule checks the pattern is valid
type AwsBatchJobQueueInvalidStateRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsBatchJobQueueInvalidStateRule returns new rule with default attributes
func NewAwsBatchJobQueueInvalidStateRule() *AwsBatchJobQueueInvalidStateRule {
	return &AwsBatchJobQueueInvalidStateRule{
		resourceType:  "aws_batch_job_queue",
		attributeName: "state",
		enum: []string{
			"ENABLED",
			"DISABLED",
		},
	}
}

// Name returns the rule name
func (r *AwsBatchJobQueueInvalidStateRule) Name() string {
	return "aws_batch_job_queue_invalid_state"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsBatchJobQueueInvalidStateRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsBatchJobQueueInvalidStateRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsBatchJobQueueInvalidStateRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsBatchJobQueueInvalidStateRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as state`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
