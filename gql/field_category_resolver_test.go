package gql

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/arxiv-graphql/config"
)

type fieldCategoryTest struct {
	Query    string
	Schema   graphql.Schema
	Expected interface{}
}

var fieldCategoryTests = []fieldCategoryTest{}

const (
	fcTestCfgFilePath = "arxiv_graphql_cfg.json"
)

func init() {
	_ = config.Load("../" + fcTestCfgFilePath)
	schema, _ := GetSchema()
	fieldCategoryTests = []fieldCategoryTest{
		{
			Query: `
				query {
					fields(tag: "astro-ph") {
						name 
						tag
						categories {
							name
							tag
						}
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"fields": []map[string]interface{}{
						{
							"categories": []map[string]interface{}{
								{
									"name": "Astrophysics of Galaxies",
									"tag":  "astro-ph.GA",
								},
								{
									"name": "Cosmology and Nongalactic Astrophysics",
									"tag":  "astro-ph.CO",
								},
								{
									"name": "Earth and Planetary Astrophysics",
									"tag":  "astro-ph.EP",
								},
								{
									"name": "High Energy Astrophysical Phenomena",
									"tag":  "astro-ph.HE",
								},
								{
									"name": "Instrumentation and Methods for Astrophysics",
									"tag":  "astro-ph.IM",
								},
								{
									"name": "Solar and Stellar Astrophysics",
									"tag":  "astro-ph.SR",
								},
							},
							"name": "Astrophysics",
							"tag":  "astro-ph",
						},
					},
				},
			},
		},
		{
			Query: `
				query {
					categories(field: "stat") {
						name
						tag
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"categories": []map[string]interface{}{
						{
							"name": "Applications",
							"tag":  "stat.AP",
						},
						{
							"name": "Computation",
							"tag":  "stat.CO",
						},
						{
							"name": "Machine Learning",
							"tag":  "stat.ML",
						},
						{
							"name": "Methodology",
							"tag":  "stat.ME",
						},
						{
							"name": "Other Statistics",
							"tag":  "stat.OT",
						},
						{
							"name": "Statistics Theory",
							"tag":  "stat.TH",
						},
					},
				},
			},
		},
	}
}
func TestFieldsAndCategoriesResolver(t *testing.T) {
	for _, test := range fieldCategoryTests {
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
		}
		testfieldAndCategory(test, params, t)
	}
}

func testfieldAndCategory(test fieldCategoryTest, p graphql.Params, t *testing.T) {
	result := graphql.Do(p)
	if len(result.Errors) > 0 {
		t.Fatalf("expected: no errors, got: %v", result.Errors)
	}
	resultJSON, _ := json.Marshal(result)
	expectedJSON, _ := json.Marshal(test.Expected)
	if !reflect.DeepEqual(resultJSON, expectedJSON) {
		t.Fatalf("expected: %v, got: %v", test.Expected, result)
	}
}
