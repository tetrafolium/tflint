package tags

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform/configs"
	"github.com/terraform-linters/tflint/tflint"
	"github.com/zclconf/go-cty/cty"
)

const (
	attributeName = "tags"
	blockName     = "tag"

	// ErrConflict is returned when conflicting tags (both attributes and blocks) are specified in a resource
	ErrConflict conflictError = "conflicting tags attribute and tag blocks found"
)

type conflictError string

func (e conflictError) Error() string {
	return string(e)
}

func tagSchema(blocks bool) *hcl.BodySchema {
	s := &hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{
				Name: attributeName,
			},
		},
	}

	if blocks {
		s.Blocks = []hcl.BlockHeaderSchema{
			{
				Type: blockName,
			},
		}
	}

	return s
}

// Decode parses the tags for a resource into a map that can be used in tag-related rules
func Decode(runner *tflint.Runner, resource *configs.Resource) (Tags, error) {
	hasBlocks := hasTagBlocks(resource)

	body, _, diags := resource.Config.PartialContent(tagSchema(hasBlocks))
	if diags.HasErrors() {
		return nil, diags
	}

	attrTags, err := decodeAttribute(runner, body.Attributes[attributeName], hasBlocks)
	if err != nil {
		return nil, err
	}

	var blockTags Tags
	if hasBlocks {
		blockTags, err = decodeBlocks(runner, body.Blocks)
		if err != nil {
			return nil, err
		}
	}

	if len(attrTags) > 0 && len(blockTags) > 0 {
		return nil, ErrConflict
	}

	if len(blockTags) > 0 {
		return blockTags, nil
	}

	return attrTags, nil
}

func decodeAttribute(runner *tflint.Runner, attr *hcl.Attribute, list bool) (Tags, error) {
	if attr == nil {
		return Tags{}, nil
	}

	if list {
		return decodeAttributeAsList(runner, attr)
	}

	tags := make(map[string]string)
	err := runner.WithExpressionContext(attr.Expr, func() error {
		return runner.EvaluateExpr(attr.Expr, &tags)
	})
	if err != nil {
		return nil, err
	}

	result := make(Tags, len(tags))
	for k, v := range tags {
		result[k] = Tag{
			Key:   k,
			Value: v,
			Range: attr.Expr.Range(),
		}
	}
	return result, nil
}

func decodeAttributeAsList(runner *tflint.Runner, attr *hcl.Attribute) (Tags, error) {
	if attr == nil {
		return Tags{}, nil
	}

	type tag struct {
		Key   string `cty:"key"`
		Value string `cty:"value"`
	}

	var tags []tag
	err := runner.WithExpressionContext(attr.Expr, func() error {
		return runner.EvaluateExprType(attr.Expr, &tags, cty.List(cty.Object(map[string]cty.Type{
			"key":   cty.String,
			"value": cty.String,
		})))
	})
	if err != nil {
		return nil, err
	}

	result := make(Tags, len(tags))
	for _, tag := range tags {
		result[tag.Key] = Tag{
			Key:   tag.Key,
			Value: tag.Value,
			Range: attr.Expr.Range(),
		}
	}
	return result, nil
}

func decodeBlocks(runner *tflint.Runner, blocks hcl.Blocks) (Tags, error) {
	result := make(Tags, len(blocks))

	for _, block := range blocks {
		body, _, diags := block.Body.PartialContent(&hcl.BodySchema{
			Attributes: []hcl.AttributeSchema{
				{
					Name:     "key",
					Required: true,
				},
				{
					Name:     "value",
					Required: true,
				},
			},
		})
		if diags.HasErrors() {
			return nil, diags
		}

		var key string
		var value string
		err := runner.EvaluateExpr(body.Attributes["key"].Expr, &key)
		if err != nil {
			return nil, err
		}
		err = runner.EvaluateExpr(body.Attributes["value"].Expr, &value)
		if err != nil {
			return nil, err
		}

		result[key] = Tag{
			Key:   key,
			Value: value,
			Range: block.DefRange,
		}
	}

	return result, nil
}
