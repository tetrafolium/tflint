// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule checks the pattern is valid
type AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsPolicyAttachmentInvalidPolicyIDRule returns new rule with default attributes
func NewAwsOrganizationsPolicyAttachmentInvalidPolicyIDRule() *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule {
	return &AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule{
		resourceType:  "aws_organizations_policy_attachment",
		attributeName: "policy_id",
		max:           130,
		pattern:       regexp.MustCompile(`^p-[0-9a-zA-Z_]{8,128}$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule) Name() string {
	return "aws_organizations_policy_attachment_invalid_policy_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"policy_id must be 130 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^p-[0-9a-zA-Z_]{8,128}$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
