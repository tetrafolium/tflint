// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/tetrafolium/tflint/tflint"
)

func Test_AwsBudgetsBudgetInvalidBudgetTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_budgets_budget" "foo" {
	budget_type = "MONEY"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsBudgetsBudgetInvalidBudgetTypeRule(),
					Message: `"MONEY" is an invalid value as budget_type`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_budgets_budget" "foo" {
	budget_type = "USAGE"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsBudgetsBudgetInvalidBudgetTypeRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
