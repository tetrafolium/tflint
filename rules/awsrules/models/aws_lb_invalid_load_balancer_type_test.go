// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/tetrafolium/tflint/tflint"
)

func Test_AwsLbInvalidLoadBalancerTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_lb" "foo" {
	load_balancer_type = "classic"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsLbInvalidLoadBalancerTypeRule(),
					Message: `"classic" is an invalid value as load_balancer_type`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_lb" "foo" {
	load_balancer_type = "application"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsLbInvalidLoadBalancerTypeRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
