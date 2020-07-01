// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/tetrafolium/tflint/tflint"
)

// AwsAppautoscalingPolicyInvalidScalableDimensionRule checks the pattern is valid
type AwsAppautoscalingPolicyInvalidScalableDimensionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAppautoscalingPolicyInvalidScalableDimensionRule returns new rule with default attributes
func NewAwsAppautoscalingPolicyInvalidScalableDimensionRule() *AwsAppautoscalingPolicyInvalidScalableDimensionRule {
	return &AwsAppautoscalingPolicyInvalidScalableDimensionRule{
		resourceType:  "aws_appautoscaling_policy",
		attributeName: "scalable_dimension",
		enum: []string{
			"ecs:service:DesiredCount",
			"ec2:spot-fleet-request:TargetCapacity",
			"elasticmapreduce:instancegroup:InstanceCount",
			"appstream:fleet:DesiredCapacity",
			"dynamodb:table:ReadCapacityUnits",
			"dynamodb:table:WriteCapacityUnits",
			"dynamodb:index:ReadCapacityUnits",
			"dynamodb:index:WriteCapacityUnits",
			"rds:cluster:ReadReplicaCount",
			"sagemaker:variant:DesiredInstanceCount",
			"custom-resource:ResourceType:Property",
			"comprehend:document-classifier-endpoint:DesiredInferenceUnits",
			"lambda:function:ProvisionedConcurrency",
			"cassandra:table:ReadCapacityUnits",
			"cassandra:table:WriteCapacityUnits",
		},
	}
}

// Name returns the rule name
func (r *AwsAppautoscalingPolicyInvalidScalableDimensionRule) Name() string {
	return "aws_appautoscaling_policy_invalid_scalable_dimension"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppautoscalingPolicyInvalidScalableDimensionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppautoscalingPolicyInvalidScalableDimensionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppautoscalingPolicyInvalidScalableDimensionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppautoscalingPolicyInvalidScalableDimensionRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as scalable_dimension`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
