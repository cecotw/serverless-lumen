package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/graph-gophers/graphql-go"
)

// GraphQl graphql handler
type GraphQl struct {
	Schema *graphql.Schema
}

// BuildSchema builds schema
func (g *GraphQl) BuildSchema(schema string, resolver interface{}) {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	g.Schema = graphql.MustParseSchema(schema, resolver, opts...)
}

var (
	// ErrQueryNameNotProvided is thrown when a name is not provided
	ErrQueryNameNotProvided = errors.New("no query was provided in the HTTP body")
)

// Lambda is the Lambda function handler
func (g *GraphQl) Lambda(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	// If no query is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrQueryNameNotProvided
	}

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		log.Print("Could not decode body", err)
	}

	response := g.Schema.Exec(context, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Print("Could not decode body")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*", // TODO use env var
			"Access-Control-Allow-Methods": "POST, GET, OPTIONS",
			"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		},
	}, nil
}
