package interactor

import (
	"context"

	"github.com/hidori/go-webapi-boilerplate/go/internal/config"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/registry"
	"gorm.io/gorm"
)

// Interactor は、ユースケースの基本実装です。
type Interactor struct {
	config *config.Config
	db     *registry.Database
	rp     *registry.Repository
	sv     *registry.Service
}

// NewInteractor は、Interactor の新規インスタンスを返します。
func NewInteractor(config *config.Config, db *registry.Database, rp *registry.Repository, sv *registry.Service) Interactor {
	return Interactor{
		config: config,
		db:     db,
		rp:     rp,
		sv:     sv,
	}
}

// ReaderTransaction は、読み取り専用エンドポイントに対するトランザクションです。
func (uc *Interactor) ReaderTransaction(ctx context.Context, fc func(*gorm.DB) error) error {
	err := uc.db.Reader.Transaction(func(tx *gorm.DB) error {
		return fc(tx)
	})
	if err != nil {
		logger.Errorf("fail to uc.db.Reader.Transaction(): err=$v", err)
	}

	return err
}

// WriterTransaction は、書き込み可能エンドポイントに対するトランザクションです。
func (uc *Interactor) WriterTransaction(ctx context.Context, fc func(*gorm.DB) error) error {
	err := uc.db.Writer.Transaction(func(tx *gorm.DB) error {
		return fc(tx)
	})
	if err != nil {
		logger.Errorf("fail to uc.db.Writer.Transaction(): err=$v", err)
	}

	return err
}
