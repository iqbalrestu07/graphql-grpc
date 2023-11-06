package services

import (
	"graph-bp/gateway/graph/generated"
	graph "graph-bp/gateway/graph/resolvers"
	epul "graph-bp/services/epul/grpc"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"google.golang.org/grpc"
)

func Setup() *handler.Server {
	epulServiceConn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		epul.NewPeopleServiceClient(epulServiceConn),
	}}))
}
