package graph

import epul "graph-bp/services/epul/grpc"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	EpulClient epul.PeopleServiceClient
}
