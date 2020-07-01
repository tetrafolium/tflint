// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsGlueCrawlerInvalidSecurityConfigurationRule checks the pattern is valid
type AwsGlueCrawlerInvalidSecurityConfigurationRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsGlueCrawlerInvalidSecurityConfigurationRule returns new rule with default attributes
func NewAwsGlueCrawlerInvalidSecurityConfigurationRule() *AwsGlueCrawlerInvalidSecurityConfigurationRule {
	return &AwsGlueCrawlerInvalidSecurityConfigurationRule{
		resourceType:  "aws_glue_crawler",
		attributeName: "security_configuration",
		max:           128,
	}
}

// Name returns the rule name
func (r *AwsGlueCrawlerInvalidSecurityConfigurationRule) Name() string {
	return "aws_glue_crawler_invalid_security_configuration"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlueCrawlerInvalidSecurityConfigurationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlueCrawlerInvalidSecurityConfigurationRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlueCrawlerInvalidSecurityConfigurationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlueCrawlerInvalidSecurityConfigurationRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"security_configuration must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
