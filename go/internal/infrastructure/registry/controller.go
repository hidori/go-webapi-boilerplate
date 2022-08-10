package registry

import (
	"github.com/labstack/echo/v4"
)

// Controller は、コントローラーのレジストリです。
type Controller struct {
	ContactAddOrUpdate echo.HandlerFunc
	ContactDeleteByID  echo.HandlerFunc
	ContactGetByID     echo.HandlerFunc
	ContactGetList     echo.HandlerFunc
}
