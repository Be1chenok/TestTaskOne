package service

import (
	"testTaskOne/internal/domain"
	"testTaskOne/internal/repository"
)

type Service struct {
	Person
}

type Person interface {
	Create(bytes []byte) (domain.Person, error)
	Add(person domain.Person) error
	Find(params map[string]string) ([]domain.Person, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Person: NewPersonService(repos.Person),
	}
}
