package handler

import (
	"urlshortener/storage"
)

type Handler struct {
	service storage.Service
}

func NewHandler(service storage.Service) *Handler {
	return &Handler{
		service: service,
	}
}
