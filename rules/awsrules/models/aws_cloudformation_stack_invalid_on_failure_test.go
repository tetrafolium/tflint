// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/tetrafolium/tflint/tflint"
)

func Test_AwsCloudformationStackInvalidOnFailureRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudformation_stack" "foo" {
	on_failure = "DO_ANYTHING"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudformationStackInvalidOnFailureRule(),
					Message: `"DO_ANYTHING" is an invalid value as on_failure`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudformation_stack" "foo" {
	on_failure = "DO_NOTHING"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudformationStackInvalidOnFailureRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
