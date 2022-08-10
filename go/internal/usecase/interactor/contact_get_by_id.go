package interactor

import (
	"context"
	"fmt"

	"github.com/hidori/go-webapi-boilerplate/go/internal/config"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/registry"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase/errors"
	"gorm.io/gorm"
)

// ContactGetByIDInteractor は、ContactGetByIDUsecase ユースケースの実装です。
type ContactGetByIDInteractor struct {
	// 基本実装
	Interactor
}

// NewContactGetByIDInteractor は、ContactGetByIDInteractor の新規インスタンスを返します。
func NewContactGetByIDInteractor(config *config.Config, db *registry.Database, rp *registry.Repository, sv *registry.Service) *ContactGetByIDInteractor {
	return &ContactGetByIDInteractor{
		Interactor: NewInteractor(config, db, rp, sv),
	}
}

// Execute は、「連絡先を取得する」ユースケースを実行します。
func (uc *ContactGetByIDInteractor) Execute(ctx context.Context, input *usecase.ContactGetByIDInputPort) (*usecase.ContactGetByIDOutputPort, error) {
	var output *usecase.ContactGetByIDOutputPort

	err := uc.ReaderTransaction(ctx, func(tx *gorm.DB) error {
		out, err := uc.rp.Contacts.GetByID(tx, input.ContactID)
		if err != nil {
			logger.Errorf("fail to it.rp.Contacts.GetByID() :err=%v, input=%v", err, input)
		}
		if out == nil {
			return errors.NewNotFoundError(fmt.Errorf("no record for input=%v", input))
		}

		output = &usecase.ContactGetByIDOutputPort{
			Model: out,
		}

		return nil
	})

	return output, err
}
