// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsIAMSamlProviderInvalidSamlMetadataDocumentRule checks the pattern is valid
type AwsIAMSamlProviderInvalidSamlMetadataDocumentRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsIAMSamlProviderInvalidSamlMetadataDocumentRule returns new rule with default attributes
func NewAwsIAMSamlProviderInvalidSamlMetadataDocumentRule() *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule {
	return &AwsIAMSamlProviderInvalidSamlMetadataDocumentRule{
		resourceType:  "aws_iam_saml_provider",
		attributeName: "saml_metadata_document",
		max:           10000000,
		min:           1000,
	}
}

// Name returns the rule name
func (r *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule) Name() string {
	return "aws_iam_saml_provider_invalid_saml_metadata_document"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"saml_metadata_document must be 10000000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"saml_metadata_document must be 1000 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
