package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) NewUser(r *http.Request) ([]byte, error) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	user, err := h.services.Person.Create(bytes)
	if err != nil {
		return nil, err
	}
	h.Users = append(h.Users, user)
	jsonUser, err := json.Marshal(user)
	return jsonUser, nil
}

func (h *Handler) PostUser() error {
	countPersons := len(h.Users)

	for i := 0; i < countPersons; i++ {
		if err := h.services.Person.Add(h.Users[i]); err != nil {
			h.Users = append(h.Users[countPersons:])
			return err
		}
	}

	h.Users = append(h.Users[countPersons:])

	return nil
}

func (h *Handler) FindUsers(r *http.Request) ([]byte, error) {
	urlParams := h.GetURLParams(r)
	users, err := h.services.Person.Find(urlParams)
	if err != nil {
		return nil, err
	}
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return jsonUsers, nil
}
