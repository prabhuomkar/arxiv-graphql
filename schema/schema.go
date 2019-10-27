package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/arxiv-graphql/resolver"
)

var (
	fields = graphql.Fields{
		"feed": &graphql.Field{
			Description: "arXiv feed of papers",
			Type:        graphql.NewList(FeedType),
			Args:        ArgsFieldConfig,
			Resolve:     resolver.FeedResolver,
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
