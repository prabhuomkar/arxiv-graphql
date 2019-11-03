package gql

import (
	"github.com/graphql-go/graphql"
)

// ArgsFieldConfig : GraphQL Field Config for arXiv Feed Arguments
var ArgsFieldConfig = graphql.FieldConfigArgument{
	"limit": &graphql.ArgumentConfig{
		Type:        graphql.Int,
		Description: "Number of feed items to return in results",
	},
	"offset": &graphql.ArgumentConfig{
		Type:        graphql.Int,
		Description: "Number of feed items to skip in results",
	},
	"sortBy": &graphql.ArgumentConfig{
		Type:        graphql.NewEnum(sortByEnumConfig),
		Description: "Sorting by field of the feed",
	},
	"sortOrder": &graphql.ArgumentConfig{
		Type:        graphql.NewEnum(sortOrderEnumConfig),
		Description: "Sorting order of the feed",
	},
	"category": &graphql.ArgumentConfig{
		Type:        graphql.String,
		Description: "Category of the arXiv feed",
	},
}

// sortByEnumConfig ...
var sortByEnumConfig = graphql.EnumConfig{
	Name: "sortBy",
	Values: graphql.EnumValueConfigMap{
		"relevance":       &graphql.EnumValueConfig{Value: SortByRelevance},
		"lastUpdatedDate": &graphql.EnumValueConfig{Value: SortByLastUpdatedDate},
		"submittedDate":   &graphql.EnumValueConfig{Value: SortBySubmittedDate},
	},
}

// sortOrderEnumConfig ...
var sortOrderEnumConfig = graphql.EnumConfig{
	Name: "sortOrder",
	Values: graphql.EnumValueConfigMap{
		"ascending":  &graphql.EnumValueConfig{Value: SortOrderAscending},
		"descending": &graphql.EnumValueConfig{Value: SortOrderDescending},
	},
}
