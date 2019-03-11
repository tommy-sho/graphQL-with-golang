package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
)

var q graphql.ObjectConfig = graphql.ObjectConfig{
	Name: "query",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			// ここで引数部分を作成
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolveID,
		},
		"name": &graphql.Field{
			Type:    graphql.String,
			Resolve: resolveName,
		},
	},
}

var schemaConfig graphql.SchemaConfig = graphql.SchemaConfig{
	Query: graphql.NewObject(q),
}

// ここでスキーマを定義
var schema, _ = graphql.NewSchema(schemaConfig)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if r.HasErrors() {
		fmt.Fprintf(os.Stderr, "An error is occured : %v", r.Errors)
	}

	j, _ := json.Marshal(r)
	fmt.Printf("%s \n", j)

	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)
	query := bufBody.String()

	result := executeQuery(query, schema)
	json.NewEncoder(w).Encode(result)
}

func main() {
	query := "{ id(id: 100), name }"
	executeQuery(query, schema)

	query = "{ id,name}"
	executeQuery(query, schema)

}

func resolveID(p graphql.ResolveParams) (interface{}, error) {
	return p.Args["id"], nil
}

func resolveName(p graphql.ResolveParams) (interface{}, error) {
	return "hoge", nil
}
