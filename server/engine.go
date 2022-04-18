package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func newEngine(schema graphql.ExecutableSchema) *gin.Engine {
	engine := gin.Default()
	engine.GET("/ping", pongHandler)
	engine.GET("/playground", playgroundHandler)
	engine.POST("/query", newOperationsHandler(schema))

	return engine
}

func pongHandler(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"message": "pong",
	})
}

func playgroundHandler(ctx *gin.Context) {
	playground.Handler("GraphQL playground", "/query").ServeHTTP(ctx.Writer, ctx.Request)
}

func newOperationsHandler(schema graphql.ExecutableSchema) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler.NewDefaultServer(schema).ServeHTTP(ctx.Writer, ctx.Request)
	}
}
