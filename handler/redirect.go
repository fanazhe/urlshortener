package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Redirect to full url if found id
func (h *Handler) Redirect(c echo.Context) error {
	id := c.Param("id")
	redirect, err := h.service.Visit(id)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return c.Redirect(http.StatusTemporaryRedirect, redirect)
}
