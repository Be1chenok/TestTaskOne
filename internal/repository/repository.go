package repository

import (
	"database/sql"
	"testTaskOne/internal/domain"
	"testTaskOne/internal/repository/postgres"
)

type Repository struct {
	Person *postgres.PersonPostgres
}

type Person interface {
	Create(bytes []byte) (domain.Person, error)
	Add(person domain.Person) error
	Find(params map[string]string) ([]domain.Person, error)
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Person: postgres.NewPersonPostgres(db),
	}
}
