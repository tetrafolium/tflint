// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsWafregionalWebACLAssociationInvalidWebACLIDRule checks the pattern is valid
type AwsWafregionalWebACLAssociationInvalidWebACLIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsWafregionalWebACLAssociationInvalidWebACLIDRule returns new rule with default attributes
func NewAwsWafregionalWebACLAssociationInvalidWebACLIDRule() *AwsWafregionalWebACLAssociationInvalidWebACLIDRule {
	return &AwsWafregionalWebACLAssociationInvalidWebACLIDRule{
		resourceType:  "aws_wafregional_web_acl_association",
		attributeName: "web_acl_id",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsWafregionalWebACLAssociationInvalidWebACLIDRule) Name() string {
	return "aws_wafregional_web_acl_association_invalid_web_acl_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafregionalWebACLAssociationInvalidWebACLIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafregionalWebACLAssociationInvalidWebACLIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafregionalWebACLAssociationInvalidWebACLIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafregionalWebACLAssociationInvalidWebACLIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"web_acl_id must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"web_acl_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*\S.*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
