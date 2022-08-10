package controller

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/interface/presenter"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
	"github.com/labstack/echo/v4"
)

// NewContactGetListController は、ContactGetListController の新規インスタンスを返します。
func NewContactGetListController(uc usecase.ContactGetListUsecase) echo.HandlerFunc {
	return HandlerFunc(func(ctx echo.Context) (interface{}, error) {
		input, err := ToContactGetListInputPort(ctx)
		if err != nil {
			logger.Errorf("fail to ToContactGetListInputPort(): err=%v, input=%v", err, input)
			return nil, err
		}

		output, err := uc.Execute(ctx.Request().Context(), input)
		if err != nil {
			logger.Errorf("fail to uc.Execute(): err=%v, input=%v", err, input)
			return nil, err
		}

		response, err := presenter.FromContactGetListOutputPort(output)
		if err != nil {
			logger.Errorf("fail presenter.FromContactGetListOutputPort(): err=%v, input=%v", err, input)
			return nil, err
		}

		return response, nil
	})
}

// ToContactGetListInputPort は、ContactGetListInputPort の新規インスタンスを返します。
func ToContactGetListInputPort(ctx echo.Context) (*usecase.ContactGetListInputPort, error) {
	return &usecase.ContactGetListInputPort{}, nil
}
