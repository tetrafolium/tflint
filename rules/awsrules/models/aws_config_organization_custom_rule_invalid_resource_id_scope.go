// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule checks the pattern is valid
type AwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule returns new rule with default attributes
func NewAwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule() *AwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule {
	return &AwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule{
		resourceType:  "aws_config_organization_custom_rule",
		attributeName: "resource_id_scope",
		max:           768,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule) Name() string {
	return "aws_config_organization_custom_rule_invalid_resource_id_scope"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"resource_id_scope must be 768 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"resource_id_scope must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
