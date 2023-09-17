package handler

import (
	"testTaskOne/internal/domain"
	"testTaskOne/internal/service"
)

type Handler struct {
	services *service.Service
	Users    []domain.Person
}

func NewHandler(users []domain.Person, services *service.Service) *Handler {
	return &Handler{
		Users:    users,
		services: services,
	}
}
