// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/tetrafolium/tflint/tflint"
)

func Test_AwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_instance" "foo" {
	instance_initiated_shutdown_behavior = "restart"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule(),
					Message: `"restart" is an invalid value as instance_initiated_shutdown_behavior`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_instance" "foo" {
	instance_initiated_shutdown_behavior = "stop"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
