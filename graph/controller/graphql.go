package controller

import (
	graph "github.com/magicalcosmos/goblogssr/graph/generated"
	"github.com/magicalcosmos/goblogssr/graph/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler Defining the Graphql handler
func GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &service.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler  Defining the Playground handler
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
