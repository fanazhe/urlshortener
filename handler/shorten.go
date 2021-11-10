package handler

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

type (
	Shorten struct {
		Url  string `json:"url" validate:"required,url"`
		User string `json:"user"`
	}
)

func isValidUrl(e string) bool {
	urlRegex := regexp.MustCompile(`^(https?|ftp)://(-\.)?([^\s/?\.#-]+\.?)+(/[^\s]*)?$`)
	return urlRegex.MatchString(e)
}

func (h *Handler) Shorten(c echo.Context) error {
	u := new(Shorten)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if !isValidUrl(u.Url) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid url"})
	}

	// Don't use this validator now
	// if err := c.Validate(u); err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }

	item, err := h.service.Add(u.Url)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, item)
}
