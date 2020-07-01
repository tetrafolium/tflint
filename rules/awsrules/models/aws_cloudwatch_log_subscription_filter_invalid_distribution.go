// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsCloudwatchLogSubscriptionFilterInvalidDistributionRule checks the pattern is valid
type AwsCloudwatchLogSubscriptionFilterInvalidDistributionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCloudwatchLogSubscriptionFilterInvalidDistributionRule returns new rule with default attributes
func NewAwsCloudwatchLogSubscriptionFilterInvalidDistributionRule() *AwsCloudwatchLogSubscriptionFilterInvalidDistributionRule {
	return &AwsCloudwatchLogSubscriptionFilterInvalidDistributionRule{
		resourceType:  "aws_cloudwatch_log_subscription_filter",
		attributeName: "distribution",
		enum: []string{
			"Random",
			"ByLogStream",
		},
	}
}

// Name returns the rule name
func (r *AwsCloudwatchLogSubscriptionFilterInvalidDistributionRule) Name() string {
	return "aws_cloudwatch_log_subscription_filter_invalid_distribution"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchLogSubscriptionFilterInvalidDistributionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchLogSubscriptionFilterInvalidDistributionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchLogSubscriptionFilterInvalidDistributionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchLogSubscriptionFilterInvalidDistributionRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as distribution`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
