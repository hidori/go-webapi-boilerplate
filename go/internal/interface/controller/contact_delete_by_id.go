package controller

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/api/param"
	"github.com/hidori/go-webapi-boilerplate/go/internal/interface/presenter"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
	"github.com/labstack/echo/v4"
)

// NewContactDeleteByIDController は、ContactDeleteByIDController の新規インスタンスを返します。
func NewContactDeleteByIDController(uc usecase.ContactDeleteByIDUsecase) echo.HandlerFunc {
	return HandlerFunc(func(ctx echo.Context) (interface{}, error) {
		input, err := ToContactDeleteByIDInputPort(ctx)
		if err != nil {
			logger.Errorf("fail to ToContactDeleteByIDInputPort(): err=%v, input=%v", err, input)
			return nil, err
		}

		output, err := uc.Execute(ctx.Request().Context(), input)
		if err != nil {
			logger.Errorf("fail to uc.Execute(): err=%v, input=%v", err, input)
			return nil, err
		}

		response, err := presenter.FromContactDeleteByIDOutputPort(output)
		if err != nil {
			logger.Errorf("fail presenter.FromContactDeleteByIDOutputPort(): err=%v, input=%v", err, input)
			return nil, err
		}

		return response, nil
	})
}

// ToContactDeleteByIDInputPort は、ContactDeleteByIDInputPort の新規インスタンスを返します。
func ToContactDeleteByIDInputPort(ctx echo.Context) (*usecase.ContactDeleteByIDInputPort, error) {
	contactID, err := param.Int(ctx, "id")
	if err != nil {
		logger.Errorf("fail to paramInt(): err=%v, name=id", err)
		return nil, err
	}

	return &usecase.ContactDeleteByIDInputPort{
		ContactID: contactID,
	}, nil
}
