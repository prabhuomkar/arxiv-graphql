package gql

import (
	"reflect"
	"testing"

	"github.com/prabhuomkar/arxiv-graphql/config"
)

type buildAPIURLTest struct {
	idList     string
	category   string
	sortBy     string
	sortOrder  string
	start      int
	maxResults int
	expected   string
}

var buildAPIURLTests = []buildAPIURLTest{}

const (
	testCfgFilePath = "arxiv_graphql_cfg.json"
)

func init() {
	_ = config.Load("../" + testCfgFilePath)
	buildAPIURLTests = []buildAPIURLTest{
		{
			category:   "cs.LG",
			sortBy:     SortByRelevance,
			sortOrder:  SortOrderAscending,
			start:      1,
			maxResults: 5,
			expected:   "http://export.arxiv.org/api/query?search_query=cat:cs.LG&sortBy=relevance&sortOrder=ascending&start=1&max_results=5",
		},
		{
			category:   "cs.LG",
			sortBy:     SortByLastUpdatedDate,
			start:      0,
			maxResults: 10,
			expected:   "http://export.arxiv.org/api/query?search_query=cat:cs.LG&sortBy=lastUpdatedDate&sortOrder=descending&start=0&max_results=10",
		},
		{
			category:  "cs.LG",
			sortBy:    SortBySubmittedDate,
			sortOrder: SortOrderAscending,
			expected:  "http://export.arxiv.org/api/query?search_query=cat:cs.LG&sortBy=submittedDate&sortOrder=ascending&start=0&max_results=10",
		},
		{
			idList:   "2004.14356",
			expected: "http://export.arxiv.org/api/query?id_list=2004.14356",
		},
	}
}

func TestUtils(t *testing.T) {
	for _, test := range buildAPIURLTests {
		found := BuildAPIURL(test.idList, test.category, test.sortBy, test.sortOrder, test.start, test.maxResults)
		if !reflect.DeepEqual(found, test.expected) {
			t.Fatalf("expected: %s, got: %s", test.expected, found)
		}
		_, err := FetchResponse(found)
		if err != nil {
			t.Fatalf("expected: no error, got: %v", err)
		}
	}
}
