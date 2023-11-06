package main

import (
	"fmt"
	_peopleHandler "graph-bp/services/epul/handler"
	_Repo "graph-bp/services/epul/repository"
	_Usecase "graph-bp/services/epul/usecase"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	peopleRepo := _Repo.NewPeopleRepository()
	peopleUc := _Usecase.NewPeopleUsecase(peopleRepo)
	_peopleHandler.NewPeopleHandler(server, peopleUc)
	conn, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Service run in port 8000")
	log.Print(server.Serve(conn))
}
