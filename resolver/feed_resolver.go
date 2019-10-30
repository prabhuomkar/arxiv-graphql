package resolver

import (
	"encoding/json"
	"strings"

	"github.com/prabhuomkar/arxiv-graphql/model"

	"github.com/clbanning/mxj"

	"github.com/graphql-go/graphql"
)

const (
	// SortOrderAscending ...
	SortOrderAscending = "ascending"
	// SortOrderDescending ...
	SortOrderDescending = "descending"

	// SortByRelevance ...
	SortByRelevance = "relevance"
	// SortByLastUpdatedDate ...
	SortByLastUpdatedDate = "lastUpdatedDate"
	// SortBySubmittedDate ...
	SortBySubmittedDate = "submittedDate"

	resultsLimit = 10

	// APIPath ...
	APIPath = "/query?search_query="
)

// FeedResolver : Resolver for query { feed }
var FeedResolver = func(p graphql.ResolveParams) (interface{}, error) {
	limit, ok := p.Args["limit"].(int)
	if !ok || limit > resultsLimit || limit < 1 {
		limit = resultsLimit
	}
	offset, ok := p.Args["offset"].(int)
	if !ok || offset < 1 {
		offset = 0
	}
	sortBy, ok := p.Args["sortBy"].(string)
	if !ok || strings.TrimSpace(sortBy) == "" {
		sortBy = SortBySubmittedDate
	}
	sortOrder, ok := p.Args["sortOrder"].(string)
	if !ok || strings.TrimSpace(sortOrder) == "" {
		sortOrder = SortOrderDescending
	}
	category, ok := p.Args["category"].(string)
	if !ok || strings.TrimSpace(category) == "" {
		category = "cs.AI"
	}

	feedEntries, err := fetchArxivFeed(category, sortBy, sortOrder, limit, offset)
	if err != nil {
		return nil, err
	}

	return feedEntries, nil
}

// fetches arXiv feed using their API
func fetchArxivFeed(category, sortBy, sortOrder string, limit, offset int) ([]model.Article, error) {
	apiURL := BuildAPIURL(category, sortBy, sortOrder, offset, limit)
	res, err := FetchResponse(apiURL)
	if err != nil {
		return nil, err
	}

	mv, err := mxj.NewMapXml(res)
	if err != nil {
		return nil, err
	}

	jsonRes, err := mv.Json()
	if err != nil {
		return nil, err
	}

	var response model.Response
	json.Unmarshal(jsonRes, &response)

	return response.Feed.Entry, nil
}
