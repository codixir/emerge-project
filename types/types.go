package types

import "github.com/graphql-go/graphql"

type Types struct{}

func (t Types) GetTypes() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "Patient",
			Description: "This is a patient type",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"phone": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
