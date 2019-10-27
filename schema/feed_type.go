package schema

import "github.com/graphql-go/graphql"

// FeedType : GraphQL type for arXiv Feed
var FeedType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "feed",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.String,
				Description: "ID of the arXiv entry",
			},
			"updated": &graphql.Field{
				Type:        graphql.String,
				Description: "Update DateTime for arXiv entry",
			},
			"published": &graphql.Field{
				Type:        graphql.String,
				Description: "Update DateTime for arXiv entry",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "Title of arXiv entry",
			},
			"summary": &graphql.Field{
				Type:        graphql.String,
				Description: "Summary of arXiv entry",
			},
			"author": &graphql.Field{
				Type:        graphql.NewList(authorType),
				Description: "Authors of arXiv entry",
			},
			"link": &graphql.Field{
				Type:        graphql.NewList(linkType),
				Description: "Links of arXiv entry",
			},
			"primaryCategory": &graphql.Field{
				Type:        categoryType,
				Description: "Primary Category of arXiv entry",
			},
			"category": &graphql.Field{
				Type:        graphql.NewList(categoryType),
				Description: "Categories of arXiv entry",
			},
		},
	},
)

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "author",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of the author",
			},
		},
	},
)

var linkType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "link",
		Fields: graphql.Fields{
			"href": &graphql.Field{
				Type:        graphql.String,
				Description: "Hyperlink of the arXiv entry",
			},
			"rel": &graphql.Field{
				Type:        graphql.String,
				Description: "Relationship between link and arXiv entry",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "Title of the link",
			},
		},
	},
)

var categoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "category",
		Fields: graphql.Fields{
			"term": &graphql.Field{
				Type:        graphql.String,
				Description: "Category Term of the arXiv entry",
			},
			"scheme": &graphql.Field{
				Type:        graphql.String,
				Description: "Schema Definition of category",
			},
		},
	},
)
