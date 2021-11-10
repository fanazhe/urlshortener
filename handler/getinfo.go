package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetInfo(c echo.Context) error {
	id := c.Param("id")
	item, err := h.service.GetInfo(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, item)
}
