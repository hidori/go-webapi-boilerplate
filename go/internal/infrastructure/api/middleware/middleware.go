package middleware

import (
	"net/http"

	"github.com/hidori/go-webapi-boilerplate/go/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Attach は、ミドルウェアをアタッチします。
func Attach(app *echo.Echo, config *config.Config) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	AttachCORS(app, config)
}

// AttachCORS は、CORS ミドルウェアをアタッチします。
func AttachCORS(app *echo.Echo, config *config.Config) {
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{
			echo.HeaderAccept,
			echo.HeaderAcceptEncoding,
			echo.HeaderContentLength,
			echo.HeaderContentType,
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
		},
		AllowOrigins: config.CORS.AllowOrigins,
	}))

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			origin := c.Request().Header.Get(echo.HeaderOrigin)
			for _, o := range config.CORS.AllowOrigins {
				if origin == o {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, map[string]string{"message:": "forbidden"})
		}
	})
}
