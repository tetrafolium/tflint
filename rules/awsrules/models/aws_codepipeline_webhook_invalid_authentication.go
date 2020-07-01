// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsCodepipelineWebhookInvalidAuthenticationRule checks the pattern is valid
type AwsCodepipelineWebhookInvalidAuthenticationRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCodepipelineWebhookInvalidAuthenticationRule returns new rule with default attributes
func NewAwsCodepipelineWebhookInvalidAuthenticationRule() *AwsCodepipelineWebhookInvalidAuthenticationRule {
	return &AwsCodepipelineWebhookInvalidAuthenticationRule{
		resourceType:  "aws_codepipeline_webhook",
		attributeName: "authentication",
		enum: []string{
			"GITHUB_HMAC",
			"IP",
			"UNAUTHENTICATED",
		},
	}
}

// Name returns the rule name
func (r *AwsCodepipelineWebhookInvalidAuthenticationRule) Name() string {
	return "aws_codepipeline_webhook_invalid_authentication"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodepipelineWebhookInvalidAuthenticationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodepipelineWebhookInvalidAuthenticationRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodepipelineWebhookInvalidAuthenticationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodepipelineWebhookInvalidAuthenticationRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as authentication`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
