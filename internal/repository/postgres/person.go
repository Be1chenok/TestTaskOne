package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"testTaskOne/internal/domain"
)

type PersonPostgres struct {
	db *sql.DB
}

func NewPersonPostgres(db *sql.DB) *PersonPostgres {
	return &PersonPostgres{db: db}
}

func (r *PersonPostgres) Create(bytes []byte) (domain.Person, error) {
	var person domain.Person

	if err := json.Unmarshal(bytes, &person); err != nil {
		return person, err
	}

	return person, nil
}

func (r *PersonPostgres) Add(person domain.Person) error {
	_, err := strconv.Atoi(person.Age)
	if err != nil {
		return err
	}
	queryUser := fmt.Sprintf("INSERT INTO %s (name, surname, age, phone) VALUES ($1, $2, $3, $4) RETURNING addres_id", peopleTable)
	row := r.db.QueryRow(
		queryUser,
		person.Name,
		person.Surname,
		person.Age,
		person.Phone)
	if row.Err() != nil {
		return row.Err()
	}

	var addres_id int
	if err = row.Scan(&addres_id); err != nil {
		return err
	}

	querryAddres := fmt.Sprintf("INSERT INTO %s (id, country, city, street, house, apartment) VALUES ($1, $2, $3, $4, $5, $6)", addresTable)
	_, err = r.db.Exec(
		querryAddres,
		addres_id,
		person.Addres.Country,
		person.Addres.City,
		person.Addres.Street,
		person.Addres.House,
		person.Addres.Apartment)
	if err != nil {
		return err
	}

	return nil
}

func (r *PersonPostgres) Find(params map[string]string) ([]domain.Person, error) {

	for k, v := range params {
		if v == "" {
			params[k] = "%"
		}
	}

	query := fmt.Sprintf("SELECT People.name, People.surname, People.age, Addres.country, Addres.city, Addres.street, Addres.house, Addres.apartment, People.phone"+
		"\nFROM %s"+
		"\nJOIN %s"+
		"\nON People.addres_id = Addres.id"+
		"\nWHERE (CAST(age as TEXT) LIKE '%v' AND country LIKE '%v' AND city LIKE '%v' AND street LIKE '%v')",
		peopleTable,
		addresTable,
		params["age"],
		params["country"],
		params["city"],
		params["street"])
	_, err := strconv.Atoi(params["size"])
	if err == nil {
		query += fmt.Sprintf("\n\tLIMIT %v", params["size"])
	}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var people []domain.Person

	for rows.Next() {
		var person domain.Person
		if err := rows.Scan(
			&person.Name,
			&person.Surname,
			&person.Age,
			&person.Addres.Country,
			&person.Addres.City,
			&person.Addres.Street,
			&person.Addres.House,
			&person.Addres.Apartment,
			&person.Phone); err != nil {
			return nil, err
		}
		people = append(people, person)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return people, nil
}
