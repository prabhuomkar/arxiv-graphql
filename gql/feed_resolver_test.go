package gql

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/arxiv-graphql/config"
)

type feedTest struct {
	Query    string
	Schema   graphql.Schema
	Expected interface{}
}

var feedTests = []feedTest{}

const (
	cfgFilePath = "arxiv_graphql_cfg.json"
)

func init() {
	_ = config.Load("../" + cfgFilePath)
	schema, _ := GetSchema()
	feedTests = []feedTest{
		{
			Query: `
				query {
					feed(category:"cs.LG", limit: 2) {
						id
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"feed": []map[string]interface{}{
						{
							"id": "http://arxiv.org/abs/1910.14673v1",
						},
						{
							"id": "http://arxiv.org/abs/1910.14670v1",
						},
					},
				},
			},
		},
		{
			Query: `
				query {
					feed {
						id
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"feed": []map[string]interface{}{
						{"id": "http://arxiv.org/abs/1910.14673v1"},
						{"id": "http://arxiv.org/abs/1910.14670v1"},
						{"id": "http://arxiv.org/abs/1910.14673v1"},
						{"id": "http://arxiv.org/abs/1910.14670v1"},
						{"id": "http://arxiv.org/abs/1910.14673v1"},
						{"id": "http://arxiv.org/abs/1910.14670v1"},
						{"id": "http://arxiv.org/abs/1910.14673v1"},
						{"id": "http://arxiv.org/abs/1910.14670v1"},
						{"id": "http://arxiv.org/abs/1910.14673v1"},
						{"id": "http://arxiv.org/abs/1910.14670v1"},
					},
				},
			},
		},
	}
}
func TestFeedResolver(t *testing.T) {
	for _, test := range feedTests {
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
		}
		testfeed(test, params, t)
	}
}

func testfeed(test feedTest, p graphql.Params, t *testing.T) {
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
