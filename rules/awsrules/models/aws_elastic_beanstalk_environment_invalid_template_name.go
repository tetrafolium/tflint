// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsElasticBeanstalkEnvironmentInvalidTemplateNameRule checks the pattern is valid
type AwsElasticBeanstalkEnvironmentInvalidTemplateNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsElasticBeanstalkEnvironmentInvalidTemplateNameRule returns new rule with default attributes
func NewAwsElasticBeanstalkEnvironmentInvalidTemplateNameRule() *AwsElasticBeanstalkEnvironmentInvalidTemplateNameRule {
	return &AwsElasticBeanstalkEnvironmentInvalidTemplateNameRule{
		resourceType:  "aws_elastic_beanstalk_environment",
		attributeName: "template_name",
		max:           100,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkEnvironmentInvalidTemplateNameRule) Name() string {
	return "aws_elastic_beanstalk_environment_invalid_template_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkEnvironmentInvalidTemplateNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkEnvironmentInvalidTemplateNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkEnvironmentInvalidTemplateNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkEnvironmentInvalidTemplateNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"template_name must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"template_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
