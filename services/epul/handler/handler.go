package handler

import (
	"context"
	"graph-bp/model"
	epul "graph-bp/services/epul/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type peopleHandler struct {
	peopleUc model.PeopleUsecase
	epul.PeopleServiceServer
}

func NewPeopleHandler(gr *grpc.Server, peopleUc model.PeopleUsecase) {
	peopleHandler := peopleHandler{
		peopleUc: peopleUc,
	}
	epul.RegisterPeopleServiceServer(gr, &peopleHandler)
}

func (h *peopleHandler) ListPeople(ctx context.Context, in *empty.Empty) (*epul.Peoples, error) {
	return h.peopleUc.Peoples()
}

func (h *peopleHandler) DetailPeople(ctx context.Context, in *epul.People) (*epul.People, error) {
	return h.peopleUc.People(in.Name)
}

func (h *peopleHandler) AddPeople(ctx context.Context, in *epul.People) (*epul.People, error) {
	return h.peopleUc.AddPeople(*in)
}
