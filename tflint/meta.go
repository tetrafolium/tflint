package tflint

import "fmt"

// Version is application version
const Version string = "0.17.0"

// ReferenceLink returns the rule reference link
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/tetrafolium/tflint/blob/v%s/docs/rules/%s.md", Version, name)
}
