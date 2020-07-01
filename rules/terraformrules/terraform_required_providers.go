package terraformrules

import (
	"fmt"
	"log"

	"github.com/tetrafolium/tflint/tflint"
)

// TerraformRequiredProvidersRule checks whether Terraform sets version constraints for all configured providers
type TerraformRequiredProvidersRule struct{}

// NewTerraformRequiredProvidersRule returns new rule with default attributes
func NewTerraformRequiredProvidersRule() *TerraformRequiredProvidersRule {
	return &TerraformRequiredProvidersRule{}
}

// Name returns the rule name
func (r *TerraformRequiredProvidersRule) Name() string {
	return "terraform_required_providers"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformRequiredProvidersRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (r *TerraformRequiredProvidersRule) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *TerraformRequiredProvidersRule) Link() string {
	return tflint.ReferenceLink(r.Name())
}

// Check checks whether variables have descriptions
func (r *TerraformRequiredProvidersRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	module := runner.TFConfig.Module
	for _, provider := range module.ProviderConfigs {
		if _, ok := module.ProviderRequirements[provider.Name]; !ok && provider.Version.Required == nil {
			message := fmt.Sprintf(`Provider "%s" should have a version constraint in required_providers`, provider.Name)
			if provider.Alias != "" {
				message += fmt.Sprintf(" (%s.%s)", provider.Name, provider.Alias)
			}
			runner.EmitIssue(r, message, provider.DeclRange)
		}
	}

	return nil
}
