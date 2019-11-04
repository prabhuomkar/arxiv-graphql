package gql

import "github.com/graphql-go/graphql"

// FieldsType : GraphQL type for arXiv Field
var FieldsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "fields",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of the Field",
			},
			"tag": &graphql.Field{
				Type:        graphql.String,
				Description: "Tag of the Field",
			},
			"categories": &graphql.Field{
				Type:        graphql.NewList(CategoriesType),
				Description: "Categories available in this Field",
			},
		},
	},
)

// CategoriesType : GraphQL type for arXiv Category
var CategoriesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "categories",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of the Category",
			},
			"tag": &graphql.Field{
				Type:        graphql.String,
				Description: "Tag of the Category",
			},
		},
	},
)
