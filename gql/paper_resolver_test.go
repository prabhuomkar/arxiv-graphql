package gql

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/arxiv-graphql/config"
)

type paperTest struct {
	Query    string
	Schema   graphql.Schema
	Expected interface{}
}

var paperTests = []paperTest{}

func init() {
	_ = config.Load("../" + cfgFilePath)
	schema, _ := GetSchema()
	paperTests = []paperTest{
		{
			Query: `
				query {
					paper(id: "2004.14356") {
						id
						title
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"paper": map[string]interface{}{
						"id":    "http://arxiv.org/abs/2004.14356v1",
						"title": "AxCell: Automatic Extraction of Results from Machine Learning Papers",
					},
				},
			},
		},
	}
}
func TestPaperResolver(t *testing.T) {
	for _, test := range paperTests {
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
		}
		testPaper(test, params, t)
	}
}

func testPaper(test paperTest, p graphql.Params, t *testing.T) {
	result := graphql.Do(p)
	if len(result.Errors) > 0 {
		t.Fatalf("expected: no errors, got: %v", result.Errors)
	}
	resultJSON, _ := json.Marshal(result)
	expectedJSON, _ := json.Marshal(test.Expected)
	if !reflect.DeepEqual(len(resultJSON), len(expectedJSON)) {
		t.Fatalf("expected: %v, got: %v", test.Expected, result)
	}
}
