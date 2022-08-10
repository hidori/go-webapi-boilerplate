package controller

import (
	"errors"
	"net/http"

	usecase "github.com/hidori/go-webapi-boilerplate/go/internal/usecase/errors"
	"github.com/labstack/echo/v4"
)

// HandlerFunc は、コントローラを内包したハンドラを返します。
func HandlerFunc(fc func(ctx echo.Context) (interface{}, error)) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		r, err := fc(ctx)
		if err == nil {
			return ctx.JSON(http.StatusOK, r)
		}

		r = &struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}

		{
			var e *usecase.BadRequestError
			if errors.As(err, &e) {
				return ctx.JSON(http.StatusBadRequest, r)
			}
		}

		{
			var e *usecase.UnauthorizedError
			if errors.As(err, &e) {
				return ctx.JSON(http.StatusUnauthorized, r)
			}
		}

		{
			var e *usecase.ForbiddenError
			if errors.As(err, &e) {
				return ctx.JSON(http.StatusForbidden, r)
			}
		}

		{
			var e *usecase.NotFoundError
			if errors.As(err, &e) {
				return ctx.JSON(http.StatusNotFound, r)
			}
		}

		return ctx.JSON(http.StatusInternalServerError, r)
	}
}
