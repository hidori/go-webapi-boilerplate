package interactor

import (
	"context"

	"github.com/hidori/go-webapi-boilerplate/go/internal/config"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/registry"
	"github.com/hidori/go-webapi-boilerplate/go/internal/usecase"
	"gorm.io/gorm"
)

// ContactGetListInteractor は、ContactGetListUsecase ユースケースの実装です。
type ContactGetListInteractor struct {
	// 基本実装
	Interactor
}

// NewContactGetListInteractor は、ContactGetListInteractor の新規インスタンスを返します。
func NewContactGetListInteractor(config *config.Config, db *registry.Database, rp *registry.Repository, sv *registry.Service) *ContactGetListInteractor {
	return &ContactGetListInteractor{
		Interactor: NewInteractor(config, db, rp, sv),
	}
}

// Execute は、「連絡先の一覧を取得する」ユースケースを実行します。
func (uc *ContactGetListInteractor) Execute(ctx context.Context, input *usecase.ContactGetListInputPort) (*usecase.ContactGetListOutputPort, error) {
	var output *usecase.ContactGetListOutputPort

	err := uc.ReaderTransaction(ctx, func(tx *gorm.DB) error {
		out, err := uc.rp.Contacts.GetList(tx)
		if err != nil {
			logger.Errorf("fail to it.rp.Contacts.GetList() :err=%v, input=%v", err, input)
		}

		output = &usecase.ContactGetListOutputPort{
			List: out,
		}

		return nil
	})

	return output, err
}
