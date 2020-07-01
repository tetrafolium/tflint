// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsAppmeshVirtualNodeInvalidMeshNameRule checks the pattern is valid
type AwsAppmeshVirtualNodeInvalidMeshNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppmeshVirtualNodeInvalidMeshNameRule returns new rule with default attributes
func NewAwsAppmeshVirtualNodeInvalidMeshNameRule() *AwsAppmeshVirtualNodeInvalidMeshNameRule {
	return &AwsAppmeshVirtualNodeInvalidMeshNameRule{
		resourceType:  "aws_appmesh_virtual_node",
		attributeName: "mesh_name",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAppmeshVirtualNodeInvalidMeshNameRule) Name() string {
	return "aws_appmesh_virtual_node_invalid_mesh_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppmeshVirtualNodeInvalidMeshNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppmeshVirtualNodeInvalidMeshNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppmeshVirtualNodeInvalidMeshNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppmeshVirtualNodeInvalidMeshNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"mesh_name must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"mesh_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
