// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule checks the pattern is valid
type AwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule returns new rule with default attributes
func NewAwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule() *AwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule {
	return &AwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule{
		resourceType:  "aws_service_discovery_public_dns_namespace",
		attributeName: "description",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule) Name() string {
	return "aws_service_discovery_public_dns_namespace_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
