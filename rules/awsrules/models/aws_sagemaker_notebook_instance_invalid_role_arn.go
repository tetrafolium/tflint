// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsSagemakerNotebookInstanceInvalidRoleArnRule checks the pattern is valid
type AwsSagemakerNotebookInstanceInvalidRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerNotebookInstanceInvalidRoleArnRule returns new rule with default attributes
func NewAwsSagemakerNotebookInstanceInvalidRoleArnRule() *AwsSagemakerNotebookInstanceInvalidRoleArnRule {
	return &AwsSagemakerNotebookInstanceInvalidRoleArnRule{
		resourceType:  "aws_sagemaker_notebook_instance",
		attributeName: "role_arn",
		max:           2048,
		min:           20,
		pattern:       regexp.MustCompile(`^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerNotebookInstanceInvalidRoleArnRule) Name() string {
	return "aws_sagemaker_notebook_instance_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerNotebookInstanceInvalidRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerNotebookInstanceInvalidRoleArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerNotebookInstanceInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerNotebookInstanceInvalidRoleArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"role_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role_arn must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
