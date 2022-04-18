package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
)

func NewLocalServer(schema graphql.ExecutableSchema) Server {
	return localServer{
		engine: newEngine(schema),
	}
}

type localServer struct {
	engine *gin.Engine
}

func (s localServer) Start() {
	s.engine.Run()
}
