package server

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

func NewLambdaServer(schema graphql.ExecutableSchema) lambdaServer {
	engine := newEngine(schema)
	return lambdaServer{
		ginLambda: ginadapter.New(engine),
	}
}

type lambdaServer struct {
	ginLambda *ginadapter.GinLambda
}

func (s lambdaServer) Start() {
	lambda.Start(s.handleRequest)
}

func (s lambdaServer) handleRequest(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	return s.ginLambda.ProxyWithContext(ctx, request)
}
