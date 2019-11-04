package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/arxiv-graphql/config"
)

// FieldsResolver : Resolver for query { fields }
var FieldsResolver = func(p graphql.ResolveParams) (interface{}, error) {
	tag, _ := p.Args["tag"].(string)

	fields := config.Config.Fields
	if tag != "" {
		for _, field := range fields {
			if field.Tag == tag {
				return []config.Field{
					field,
				}, nil
			}
		}
	}

	return fields, nil
}

// CategoriesResolver : Resolver for query { categories }
var CategoriesResolver = func(p graphql.ResolveParams) (interface{}, error) {
	fieldTag, _ := p.Args["field"].(string)

	fields := config.Config.Fields

	if fieldTag != "" {
		for _, field := range fields {
			if field.Tag == fieldTag {
				return field.Categories, nil
			}
		}
	}

	categories := []config.Category{}
	for _, field := range fields {
		for _, category := range field.Categories {
			categories = append(categories, category)
		}
	}

	return categories, nil
}
