package repository

import (
	"github.com/hidori/go-webapi-boilerplate/go/internal/domain/model"
	"gorm.io/gorm"
)

// ContactRepository は、Contact レポジトリのインターフェースです。
type ContactRepository interface {
	// GetList は、すべてのレコードを取得します。レコードが存在しない時は空の配列を返します。
	GetList(db *gorm.DB) ([]model.Contact, error)

	// GetByID は、指定された条件にマッチする１件のレコードを取得します。条件にマッチするレコードが存在しない時は nil を返します。
	GetByID(db *gorm.DB, contactID int) (*model.Contact, error)

	// AddOrUpdate は、１件のレコードを追加または更新して、そのレコードを返します。
	AddOrUpdate(db *gorm.DB, contact *model.Contact) (*model.Contact, error)

	// DeleteByID は、指定された条件にマッチする１件のレコードを削除します。条件にマッチするレコードが存在しない時はエラーを返します。
	DeleteByID(db *gorm.DB, contactID int) error
}
