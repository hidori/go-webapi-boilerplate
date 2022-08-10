package interactor

import (
	"context"

	"github.com/hidori/go-webapi-boilerplate/go/internal/config"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/registry"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
	"gorm.io/gorm"
)

// ContactAddOrUpdateInteractor は、ContactAddOrUpdateUsecase ユースケースの実装です。
type ContactAddOrUpdateInteractor struct {
	// 基本実装
	Interactor
}

// NewContactAddOrUpdateInteractor は、ContactAddOrUpdateInteractor の新規インスタンスを返します。
func NewContactAddOrUpdateInteractor(config *config.Config, db *registry.Database, rp *registry.Repository, sv *registry.Service) *ContactAddOrUpdateInteractor {
	return &ContactAddOrUpdateInteractor{
		Interactor: NewInteractor(config, db, rp, sv),
	}
}

// Execute は、「連絡先を追加または更新する」ユースケースを実行します。
func (uc *ContactAddOrUpdateInteractor) Execute(ctx context.Context, input *usecase.ContactAddOrUpdateInputPort) (*usecase.ContactAddOrUpdateOutputPort, error) {
	var output *usecase.ContactAddOrUpdateOutputPort

	err := uc.WriterTransaction(ctx, func(tx *gorm.DB) error {
		out, err := uc.rp.Contacts.AddOrUpdate(tx, input.Model)
		if err != nil {
			logger.Errorf("fail to it.rp.Contacts.AddOrUpdate() :err=%v, input=%v", err, input)
		}

		output = &usecase.ContactAddOrUpdateOutputPort{
			Model: out,
		}

		return nil
	})

	return output, err
}
