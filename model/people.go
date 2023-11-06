package model

import epul "graph-bp/services/epul/grpc"

type PeopleRepository interface {
	Peoples() (*epul.Peoples, error)
	People(name string) (*epul.People, error)
	AddPeople(epul.People) (*epul.People, error)
}
type PeopleUsecase interface {
	Peoples() (*epul.Peoples, error)
	People(name string) (*epul.People, error)
	AddPeople(epul.People) (*epul.People, error)
}
