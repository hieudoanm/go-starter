package schema

import (
	"github.com/graphql-go/graphql"

	tasks_resolver "graphql-starter/src/graphql/tasks/resolver"
	tasks_schema "graphql-starter/src/graphql/tasks/schema"
)

func CreateGraphQlSchema() (graphql.Schema, error) {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"tasks": &graphql.Field{
				Name:    "tasks",
				Type:    graphql.NewList(tasks_schema.Task),
				Resolve: tasks_resolver.GetTasks,
			},
			"task": &graphql.Field{
				Name: "task",
				Type: tasks_schema.Task,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: tasks_resolver.GetTask,
			},
		},
	})

	// above schema definition code
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
