// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/tetrafolium/tflint/tflint"
)

func Test_AwsDmsReplicationTaskInvalidMigrationTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_dms_replication_task" "foo" {
	migration_type = "partial-load"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsDmsReplicationTaskInvalidMigrationTypeRule(),
					Message: `"partial-load" is an invalid value as migration_type`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_dms_replication_task" "foo" {
	migration_type = "full-load"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsDmsReplicationTaskInvalidMigrationTypeRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
