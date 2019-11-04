package gql

import (
	"github.com/graphql-go/graphql"
)

var (
	fields = graphql.Fields{
		"feed": &graphql.Field{
			Description: "arXiv feed of papers",
			Type:        graphql.NewList(FeedType),
			Args:        ArgsFieldConfig,
			Resolve:     FeedResolver,
		},
		"fields": &graphql.Field{
			Description: "list of arXiv fields",
			Type:        graphql.NewList(FieldsType),
			Args:        FieldArgsFieldConfig,
			Resolve:     FieldsResolver,
		},
		"categories": &graphql.Field{
			Description: "list of arXiv categories",
			Type:        graphql.NewList(CategoriesType),
			Args:        CategoryArgsFieldConfig,
			Resolve:     CategoriesResolver,
		},
	}

	schemaConfig = graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Query",
			Fields:      fields,
			Description: "Query Type for arXiv GraphQL API",
		}),
	}
)

// GetSchema : Returns global schema
func GetSchema() (graphql.Schema, error) {
	return graphql.NewSchema(schemaConfig)
}
