package interactor

import (
	"context"

	"github.com/hidori/go-webapi-boilerplate/go/internal/config"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/registry"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
	"gorm.io/gorm"
)

// ContactDeleteByIDInteractor は、ContactDeleteByIDUsecase ユースケースの実装です。
type ContactDeleteByIDInteractor struct {
	// 基本実装
	Interactor
}

// NewContactDeleteByIDInteractor は、ContactDeleteByIDInteractor の新規インスタンスを返します。
func NewContactDeleteByIDInteractor(config *config.Config, db *registry.Database, rp *registry.Repository, sv *registry.Service) *ContactDeleteByIDInteractor {
	return &ContactDeleteByIDInteractor{
		Interactor: NewInteractor(config, db, rp, sv),
	}
}

// Execute は、「連絡先を削除する」ユースケースを実行します。
func (uc *ContactDeleteByIDInteractor) Execute(ctx context.Context, input *usecase.ContactDeleteByIDInputPort) (*usecase.ContactDeleteByIDOutputPort, error) {
	var output *usecase.ContactDeleteByIDOutputPort

	err := uc.WriterTransaction(ctx, func(tx *gorm.DB) error {
		err := uc.rp.Contacts.DeleteByID(tx, input.ContactID)
		if err != nil {
			logger.Errorf("fail to it.rp.Contacts.DeleteByID() :err=%v, input=%v", err, input)
		}

		output = &usecase.ContactDeleteByIDOutputPort{}

		return nil
	})

	return output, err
}
