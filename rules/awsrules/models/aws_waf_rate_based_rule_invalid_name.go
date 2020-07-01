// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsWafRateBasedRuleInvalidNameRule checks the pattern is valid
type AwsWafRateBasedRuleInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsWafRateBasedRuleInvalidNameRule returns new rule with default attributes
func NewAwsWafRateBasedRuleInvalidNameRule() *AwsWafRateBasedRuleInvalidNameRule {
	return &AwsWafRateBasedRuleInvalidNameRule{
		resourceType:  "aws_waf_rate_based_rule",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsWafRateBasedRuleInvalidNameRule) Name() string {
	return "aws_waf_rate_based_rule_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafRateBasedRuleInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafRateBasedRuleInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafRateBasedRuleInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafRateBasedRuleInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
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
