package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prabhuomkar/arxiv-graphql/gql"

	"github.com/prabhuomkar/arxiv-graphql/config"

	"github.com/graphql-go/handler"
)

const (
	cfgFilePath = "arxiv_graphql_cfg.json"
)

func main() {
	err := config.Load(cfgFilePath)
	if err != nil {
		panic(err)
	}

	schema, err := gql.GetSchema()

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	log.Printf("Starting %s", config.Config.Service.Name)

	http.Handle("/graphql", h)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Printf("Cannot start %s", config.Config.Service.Name)
		panic(err)
	}
}
