package repository

import (
	"graph-bp/model"
	epul "graph-bp/services/epul/grpc"
)

type PeopleRepository struct {
}

func NewPeopleRepository() model.PeopleRepository {
	return &PeopleRepository{}
}

var P = epul.Peoples{
	Peoples: []*epul.People{
		{
			Name:  "iqbal",
			Age:   17,
			Phone: "081211",
			Address: &epul.Address{
				Province: "Jawa Barat",
				City:     "Bogor",
			},
		},
		{
			Name:  "restu",
			Age:   17,
			Phone: "081211",
			Address: &epul.Address{
				Province: "Jawa Barat",
				City:     "Bogor",
			},
		},
	},
}

func (r *PeopleRepository) Peoples() (peoples *epul.Peoples, err error) {

	return &P, nil
}

func (r *PeopleRepository) People(name string) (people *epul.People, err error) {
	for _, v := range P.Peoples {
		if v.Name == name {
			people = &epul.People{
				Name:    v.Name,
				Age:     v.Age,
				Phone:   v.Phone,
				Address: v.Address,
			}
			break
		}
	}
	return people, nil
}

func (r *PeopleRepository) AddPeople(in epul.People) (peoples *epul.People, err error) {
	P.Peoples = append(P.Peoples, &epul.People{
		Name:  in.Name,
		Age:   in.Age,
		Phone: in.Phone,
		Address: &epul.Address{
			Province: in.Address.Province,
			City:     in.Address.City,
		},
	})
	l := len(P.Peoples) - 1
	return P.Peoples[l], nil
}
