// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsDxPublicVirtualInterfaceInvalidAddressFamilyRule checks the pattern is valid
type AwsDxPublicVirtualInterfaceInvalidAddressFamilyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDxPublicVirtualInterfaceInvalidAddressFamilyRule returns new rule with default attributes
func NewAwsDxPublicVirtualInterfaceInvalidAddressFamilyRule() *AwsDxPublicVirtualInterfaceInvalidAddressFamilyRule {
	return &AwsDxPublicVirtualInterfaceInvalidAddressFamilyRule{
		resourceType:  "aws_dx_public_virtual_interface",
		attributeName: "address_family",
		enum: []string{
			"ipv4",
			"ipv6",
		},
	}
}

// Name returns the rule name
func (r *AwsDxPublicVirtualInterfaceInvalidAddressFamilyRule) Name() string {
	return "aws_dx_public_virtual_interface_invalid_address_family"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDxPublicVirtualInterfaceInvalidAddressFamilyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDxPublicVirtualInterfaceInvalidAddressFamilyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDxPublicVirtualInterfaceInvalidAddressFamilyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDxPublicVirtualInterfaceInvalidAddressFamilyRule) Check(runner *tflint.Runner) error {
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
