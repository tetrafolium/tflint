package terraformrules

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
	"log"
)

// TerraformRequiredVersionRule checks whether a terraform version has required_version attribute
type TerraformRequiredVersionRule struct{}

// NewTerraformRequiredVersionRule returns new rule with default attributes
func NewTerraformRequiredVersionRule() *TerraformRequiredVersionRule {
	return &TerraformRequiredVersionRule{}
}

// Name returns the rule name
func (r *TerraformRequiredVersionRule) Name() string {
	return "terraform_required_version"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformRequiredVersionRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (r *TerraformRequiredVersionRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *TerraformRequiredVersionRule) Link() string {
	return tflint.ReferenceLink(r.Name())
}

// Check checks whether variables have descriptions
func (r *TerraformRequiredVersionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	module := runner.TFConfig.Module
	versionConstraints := module.CoreVersionConstraints
	if len(versionConstraints) == 0 {
		runner.EmitIssue(
			r,
			fmt.Sprintf("terraform \"required_version\" attribute is required"),
			hcl.Range{},
		)
		return nil
	}

	return nil
}
