// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/tetrafolium/tflint/tflint"
)

func Test_AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudwatch_metric_alarm" "foo" {
	extended_statistic = "p101"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudwatchMetricAlarmInvalidExtendedStatisticRule(),
					Message: `"p101" does not match valid pattern ^p(\d{1,2}(\.\d{0,2})?|100)$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudwatch_metric_alarm" "foo" {
	extended_statistic = "p100"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudwatchMetricAlarmInvalidExtendedStatisticRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
