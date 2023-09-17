package service

import (
	"testTaskOne/internal/domain"
	"testTaskOne/internal/repository"
)

type PersonService struct {
	repos repository.Person
}

func NewPersonService(repos repository.Person) *PersonService {
	return &PersonService{repos: repos}
}

func (s *PersonService) Create(bytes []byte) (domain.Person, error) {
	return s.repos.Create(bytes)
}

func (s *PersonService) Add(person domain.Person) error {
	return s.repos.Add(person)
}

func (s *PersonService) Find(params map[string]string) ([]domain.Person, error) {
	return s.repos.Find(params)
}
