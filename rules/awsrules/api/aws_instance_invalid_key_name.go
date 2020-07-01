// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsInstanceInvalidKeyNameRule checks whether attribute value actually exists
type AwsInstanceInvalidKeyNameRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsInstanceInvalidKeyNameRule returns new rule with default attributes
func NewAwsInstanceInvalidKeyNameRule() *AwsInstanceInvalidKeyNameRule {
	return &AwsInstanceInvalidKeyNameRule{
		resourceType:  "aws_instance",
		attributeName: "key_name",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsInstanceInvalidKeyNameRule) Name() string {
	return "aws_instance_invalid_key_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsInstanceInvalidKeyNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsInstanceInvalidKeyNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsInstanceInvalidKeyNameRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by DescribeKeyPairs
func (r *AwsInstanceInvalidKeyNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeKeyPairs")
			var err error
			r.data, err = runner.AwsClient.DescribeKeyPairs()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeKeyPairs",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid key name.`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
