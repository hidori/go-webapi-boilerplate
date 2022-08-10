package router

import (
	"net/http"

	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/registry"
	"github.com/labstack/echo/v4"
)

// Attach は、コントローラーをパスにアタッチします。
func Attach(group *echo.Group, controllers *registry.Controller) {
	group.GET("health", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})

	AttachContacts(group, controllers)
}

// AttachContacts は、連絡先のコントローラーをパスにアタッチします。
func AttachContacts(group *echo.Group, ct *registry.Controller) {
	group.GET("contacts", ct.ContactGetList)
	group.GET("contacts/:id", ct.ContactGetByID)
	group.POST("contacts", ct.ContactAddOrUpdate)
	group.DELETE("contacts/:id", ct.ContactDeleteByID)
}
