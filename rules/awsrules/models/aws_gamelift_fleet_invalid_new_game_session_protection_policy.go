// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule checks the pattern is valid
type AwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule returns new rule with default attributes
func NewAwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule() *AwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule {
	return &AwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule{
		resourceType:  "aws_gamelift_fleet",
		attributeName: "new_game_session_protection_policy",
		enum: []string{
			"NoProtection",
			"FullProtection",
		},
	}
}

// Name returns the rule name
func (r *AwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule) Name() string {
	return "aws_gamelift_fleet_invalid_new_game_session_protection_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as new_game_session_protection_policy`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
