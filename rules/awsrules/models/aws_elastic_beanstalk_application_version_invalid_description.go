// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsElasticBeanstalkApplicationVersionInvalidDescriptionRule checks the pattern is valid
type AwsElasticBeanstalkApplicationVersionInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsElasticBeanstalkApplicationVersionInvalidDescriptionRule returns new rule with default attributes
func NewAwsElasticBeanstalkApplicationVersionInvalidDescriptionRule() *AwsElasticBeanstalkApplicationVersionInvalidDescriptionRule {
	return &AwsElasticBeanstalkApplicationVersionInvalidDescriptionRule{
		resourceType:  "aws_elastic_beanstalk_application_version",
		attributeName: "description",
		max:           200,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkApplicationVersionInvalidDescriptionRule) Name() string {
	return "aws_elastic_beanstalk_application_version_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkApplicationVersionInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkApplicationVersionInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkApplicationVersionInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkApplicationVersionInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 200 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
