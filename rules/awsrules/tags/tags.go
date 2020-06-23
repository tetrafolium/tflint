package tags

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform/configs"
)

//go:generate go run ./generator/main.go

// Has returns whether the resource has tags
func Has(resource *configs.Resource) bool {
	_, ok := resources[resource.Type]
	return ok
}

func hasTagBlocks(resource *configs.Resource) bool {
	return resource.Type == "aws_autoscaling_group"
}

// Tag represents a key-value AWS tag.
// Most tags are defined in a single map expression, but for compatibility
// with resources that use tag {} blocks an hcl.Range is included with each tag.
type Tag struct {
	Key   string
	Value string
	Range hcl.Range
}

// Tags holds a map of Tag objects, where each key is equal to the Key value of the tag
type Tags map[string]Tag
