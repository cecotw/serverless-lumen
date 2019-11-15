package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/flytedesk/foundation/services/graphql-go/internal/app/graphql/schema"
	"github.com/flytedesk/foundation/services/graphql-go/internal/pkg/handler"
	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("GO_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load() // The Original .env
}

func main() {
	var graphql = new(handler.GraphQl)
	graphql.BuildSchema(schema.MergeSchema(), &schema.QueryResolver{})
	lambda.Start(graphql.Lambda)
}
