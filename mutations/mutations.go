package mutations

import (
	"database/sql"
	"emerge-project/models"
	"log"

	"github.com/graphql-go/graphql"
)

type Mutations struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (q Mutations) GetMutations(patientType *graphql.Object, db *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutations",
			Fields: graphql.Fields{
				"create": &graphql.Field{
					Type:        patientType,
					Description: "Creates a new patient",
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"email": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"phone": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						patient := models.Patient{}
						var lastInsertID int

						name, _ := params.Args["name"].(string)
						email, _ := params.Args["email"].(string)
						phone, _ := params.Args["phone"].(string)

						stmt := "insert into patients(name, email, phone) values($1, $2, $3) returning id;"
						err := db.QueryRow(stmt, name, email, phone).Scan(&lastInsertID)

						logFatal(err)

						patient.ID = lastInsertID
						patient.Name = name
						patient.Email = email
						patient.Phone = phone

						return patient, nil
					},
				},
				"update": &graphql.Field{
					Type:        patientType,
					Description: "Updates an existing patient",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"email": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"phone": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						patient := models.Patient{}

						id, _ := params.Args["id"].(int)
						name, _ := params.Args["name"].(string)
						email, _ := params.Args["email"].(string)
						phone, _ := params.Args["phone"].(string)

						stmt, err := db.Prepare("update patients set name = $1, email = $2, phone = $3 where id = $4")
						logFatal(err)

						_, err = stmt.Exec(name, email, phone, id)
						logFatal(err)

						patient.ID = id
						patient.Name = name
						patient.Email = email
						patient.Phone = phone

						return patient, nil
					},
				},
				"delete": &graphql.Field{
					Type:        patientType,
					Description: "Deletes a patient",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						id, _ := params.Args["id"].(int)

						stmt, err := db.Prepare("delete from patients where id = $1")
						logFatal(err)

						_, err = stmt.Exec(id)
						logFatal(err)

						return nil, nil
					},
				},
			},
		},
	)
}
