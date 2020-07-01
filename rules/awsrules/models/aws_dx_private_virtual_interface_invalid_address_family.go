// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule checks the pattern is valid
type AwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule returns new rule with default attributes
func NewAwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule() *AwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule {
	return &AwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule{
		resourceType:  "aws_dx_private_virtual_interface",
		attributeName: "address_family",
		enum: []string{
			"ipv4",
			"ipv6",
		},
	}
}

// Name returns the rule name
func (r *AwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule) Name() string {
	return "aws_dx_private_virtual_interface_invalid_address_family"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as address_family`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
