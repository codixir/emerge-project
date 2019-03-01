package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"emerge-project/driver"
	"emerge-project/mutations"
	"emerge-project/queries"
	"emerge-project/types"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	err := godotenv.Load()
	logFatal(err)
}

var db *sql.DB

func main() {
	db = driver.Connect()

	defer db.Close()

	types := types.Types{}
	queries := queries.Queries{}
	mutations := mutations.Mutations{}

	//step 1, a patientType

	var patientType = types.GetTypes()

	//step 2, a queryType --- queries the database / does not modify/mutate the data

	var queryType = queries.GetQueries(patientType, db)

	//step 3, a mutationType --- queries the database / but it changes/mutates the data

	var mutationType = mutations.GetMutations(patientType, db)

	//step 4, a schema -- an object that has the queryType and mutationType

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)

	//step 5, a graphql method called Do, that takes schema and a requestString and
	//returns a result..

	r := mux.NewRouter()
	r.HandleFunc("/patient", func(w http.ResponseWriter, r *http.Request) {

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})

		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", r)
}
