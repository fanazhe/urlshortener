package router

import (
	"urlshortener/config"
	"urlshortener/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func registerRoutes(e *echo.Echo, h *handler.Handler) {
	e.POST("/shorten", h.Shorten)
	e.GET("/:id", h.Redirect)
	e.GET("/:id/info", h.GetInfo)
}

func New(h *handler.Handler, config *config.Config) *echo.Echo {
	e := echo.New()

	if config.APP_ENV == "production" {
		e.HideBanner = true
		e.Static("/", config.WebPath)
		e.File("/favicon.ico", config.WebPath+"/favicon.ico")
		e.Logger.SetLevel(log.INFO)
	} else {
		e.Logger.SetLevel(log.DEBUG)
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))
	}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Validator = NewValidator()
	registerRoutes(e, h)
	return e
}
