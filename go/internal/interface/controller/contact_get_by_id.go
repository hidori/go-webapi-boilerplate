package controller

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/api/param"
	"github.com/hidori/go-webapi-boilerplate/go/internal/interface/presenter"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase/errors"
	"github.com/labstack/echo/v4"
)

// NewContactGetByIDController は、ContactGetByIDController の新規インスタンスを返します。
func NewContactGetByIDController(uc usecase.ContactGetByIDUsecase) echo.HandlerFunc {
	return HandlerFunc(func(ctx echo.Context) (interface{}, error) {
		input, err := ToContactGetByIDInputPort(ctx)
		if err != nil {
			logger.Errorf("fail to ToContactGetByIDInputPort(): err=%v", err)
			return nil, err
		}

		output, err := uc.Execute(ctx.Request().Context(), input)
		if err != nil {
			logger.Errorf("fail to uc.Execute(): err=%v, input=%v", err, input)
			return nil, err
		}

		response, err := presenter.FromContactGetByIDOutputPort(output)
		if err != nil {
			logger.Errorf("fail presenter.FromContactGetByIDOutputPort(): err=%v, input=%v", err, input)
			return nil, err
		}

		return response, nil
	})
}

// ToContactGetByIDInputPort は、ContactGetByIDInputPort の新規インスタンスを返します。
func ToContactGetByIDInputPort(ctx echo.Context) (*usecase.ContactGetByIDInputPort, error) {
	contactID, err := param.Int(ctx, "id")
	if err != nil {
		logger.Errorf("fail to paramInt(): err=%v, name=id", err)
		return nil, errors.NewBadRequestError(err)
	}

	return &usecase.ContactGetByIDInputPort{
		ContactID: contactID,
	}, nil
}
