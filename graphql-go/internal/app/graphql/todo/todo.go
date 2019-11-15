package todo

import (
	"fmt"

	"github.com/flytedesk/foundation/services/graphql-go/internal/pkg/db"
	"github.com/graph-gophers/graphql-go"
)

// Todo Todo
type Todo struct {
	ID         graphql.ID
	Message    *string
	IsComplete *bool `db:"is_complete"`
}

// Input Todo Input
type Input struct {
	ID         *graphql.ID
	Message    string
	IsComplete bool `db:"is_complete"`
}

// Resolver Todo Resolver
type Resolver struct{}

// Todos : Resolver function for the "Todo" query
func (r *Resolver) Todos() *[]*Todo {
	db := db.Connect()
	defer db.Close()
	todos := []*Todo{}
	db.Select(&todos, "SELECT * FROM todos")
	return &todos
}

// Todo : Resolver function for the "Todo" query
func (r *Resolver) Todo(args struct{ ID graphql.ID }) *Todo {
	db := db.Connect()
	defer db.Close()
	todo := &Todo{}
	db.Get(&todo, "SELECT * FROM todos WHERE id=$1", args.ID)
	return todo
}

// CreateTodo : Create a todo
func (r *Resolver) CreateTodo(args struct{ Input Input }) *Todo {
	db := db.Connect()
	defer db.Close()
	row := db.QueryRowx(
		"INSERT INTO todos (id, message, is_complete) VALUES (gen_random_uuid(), $1, $2) RETURNING *",
		&args.Input.Message,
		&args.Input.IsComplete,
	)
	todo := &Todo{}
	err := row.StructScan(todo)
	if err != nil {
		fmt.Println(err.Error())
	}
	return todo
}
