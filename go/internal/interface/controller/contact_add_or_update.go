package controller

import (
	"fmt"

	"github.com/hidori/go-webapi-boilerplate/go/internal/domain/model"
	"github.com/hidori/go-webapi-boilerplate/go/internal/interface/presenter"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase/errors"
	"github.com/labstack/echo/v4"
)

// NewContactAddOrUpdateController は、ContactAddOrUpdateController の新規インスタンスを返します。
func NewContactAddOrUpdateController(uc usecase.ContactAddOrUpdateUsecase) echo.HandlerFunc {
	return HandlerFunc(func(ctx echo.Context) (interface{}, error) {
		input, err := ToContactAddOrUpdateInputPort(ctx)
		if err != nil {
			logger.Errorf("fail to ToContactAddOrUpdateInputPort(): err=%v, input=%v", err, input)
			return nil, errors.NewBadRequestError(err)
		}

		output, err := uc.Execute(ctx.Request().Context(), input)
		if err != nil {
			logger.Errorf("fail to uc.Execute(): err=%v, input=%v", err, input)
			return nil, err
		}

		response, err := presenter.FromContactAddOrUpdateOutputPort(output)
		if err != nil {
			logger.Errorf("fail presenter.FromContactAddOrUpdateOutputPort(): err=%v, input=%v", err, input)
			return nil, err
		}

		return response, nil
	})
}

// ToContactAddOrUpdateInputPort は、ContactAddOrUpdateInputPort の新規インスタンスを返します。
func ToContactAddOrUpdateInputPort(ctx echo.Context) (*usecase.ContactAddOrUpdateInputPort, error) {
	var model model.Contact

	err := ctx.Bind(&model)
	if err != nil {
		logger.Errorf("fail to ctx.Bind(): err=%v", err)
		return nil, err
	}

	results, err := model.Validate()
	if err != nil {
		logger.Errorf("fail to ctx.Bind(): err=%v", err)
		return nil, err
	}
	if len(results) > 0 {
		// TODO: ミドルウェアで HTTP StatusCode 404 に変換可能なエラーを返す
		err := fmt.Errorf("validation error: %v", results)
		return nil, err
	}

	return &usecase.ContactAddOrUpdateInputPort{
		Model: &model,
	}, nil
}
