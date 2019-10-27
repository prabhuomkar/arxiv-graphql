package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/prabhuomkar/arXiv/schema"
)

func main() {
	schema, err := schema.GetSchema()

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
		log.Printf("Cannot start %s", "arXiv GraphQL API")
	} else {
		log.Printf("Started %s", "arXiv GraphQL API")
		log.Printf("Listening on: %s", "localhost:8080")
	}
}
