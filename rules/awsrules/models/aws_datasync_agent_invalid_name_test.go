// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/tetrafolium/tflint/tflint"
)

func Test_AwsDatasyncAgentInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_datasync_agent" "foo" {
	name = "example^example"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsDatasyncAgentInvalidNameRule(),
					Message: `"example^example" does not match valid pattern ^[a-zA-Z0-9\s+=._:@/-]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_datasync_agent" "foo" {
	name = "example"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsDatasyncAgentInvalidNameRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
