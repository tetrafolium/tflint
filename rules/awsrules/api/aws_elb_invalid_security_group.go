// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsELBInvalidSecurityGroupRule checks whether attribute value actually exists
type AwsELBInvalidSecurityGroupRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsELBInvalidSecurityGroupRule returns new rule with default attributes
func NewAwsELBInvalidSecurityGroupRule() *AwsELBInvalidSecurityGroupRule {
	return &AwsELBInvalidSecurityGroupRule{
		resourceType:  "aws_elb",
		attributeName: "security_groups",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsELBInvalidSecurityGroupRule) Name() string {
	return "aws_elb_invalid_security_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsELBInvalidSecurityGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsELBInvalidSecurityGroupRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsELBInvalidSecurityGroupRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by DescribeSecurityGroups
func (r *AwsELBInvalidSecurityGroupRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeSecurityGroups")
			var err error
			r.data, err = runner.AwsClient.DescribeSecurityGroups()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeSecurityGroups",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		return runner.EachStringSliceExprs(attribute.Expr, func(val string, expr hcl.Expression) {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid security group.`, val),
					expr.Range(),
				)
			}
		})
	})
}
