package main

import (
	"graph-bp/gateway/services"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	server := gin.Default()
	server.GET("/", playgrounHandler())
	server.POST("/query", graphqlHandler())
	server.Run(":" + defaultPort)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))

}

func playgrounHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func graphqlHandler() gin.HandlerFunc {
	srv := services.Setup()
	// h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	// return func(ctx *gin.Context) {
	// 	h.ServeHTTP(ctx.Writer, ctx.Request)
	// }
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{
		MaxUploadSize: 128000,
		MaxMemory:     1280,
	})
	srv.Use(extension.Introspection{})
	return func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	}

}
