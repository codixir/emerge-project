package queries

import (
	"database/sql"
	"emerge-project/models"
	"log"

	"github.com/graphql-go/graphql"
)

type Queries struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (q Queries) GetQueries(patientType *graphql.Object, db *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getPatient": &graphql.Field{
					Type:        patientType,
					Description: "Get a patient by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id, _ := p.Args["id"].(int)
						patient := &models.Patient{}

						err := db.QueryRow("select id, name, email, phone from patients where id=$1", id).
							Scan(&patient.ID, &patient.Name, &patient.Email, &patient.Phone)

						logFatal(err)

						return patient, nil
					},
				},
				"getPatients": &graphql.Field{
					Type:        graphql.NewList(patientType),
					Description: "Gets a list of patients",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						var patients []*models.Patient

						rows, err := db.Query("select * from patients")
						logFatal(err)

						for rows.Next() {
							patient := &models.Patient{}

							rows.Scan(&patient.ID, &patient.Name, &patient.Email, &patient.Phone)
							patients = append(patients, patient)
						}

						return patients, nil
					},
				},
			},
		},
	)
}
