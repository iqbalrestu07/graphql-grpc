package usecase

import (
	"graph-bp/model"
	epul "graph-bp/services/epul/grpc"
)

type PeopleUsecase struct {
	peopleRepo model.PeopleRepository
}

func NewPeopleUsecase(peopleRepo model.PeopleRepository) model.PeopleUsecase {
	return &PeopleUsecase{
		peopleRepo: peopleRepo,
	}
}

func (u *PeopleUsecase) Peoples() (*epul.Peoples, error) {
	return u.peopleRepo.Peoples()
}

func (u *PeopleUsecase) People(name string) (*epul.People, error) {
	return u.peopleRepo.People(name)
}

func (u *PeopleUsecase) AddPeople(in epul.People) (*epul.People, error) {
	return u.peopleRepo.AddPeople(in)
}
