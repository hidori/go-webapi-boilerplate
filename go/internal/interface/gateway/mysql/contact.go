package mysql

import (
	"errors"

	"github.com/hidori/go-webapi-boilerplate/go/internal/domain/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ContactRepository は、Contact レポジトリの実装です。
type ContactRepository struct{}

// NewContactRepository は、Contact レポジトリの新規インスタンスを返します。
func NewContactRepository() *ContactRepository {
	return &ContactRepository{}
}

// GetList は、すべてのレコードを取得します。レコードが存在しない時は空の配列を返します。
func (rp *ContactRepository) GetList(db *gorm.DB) ([]model.Contact, error) {
	var list []model.Contact

	err := db.Find(&list).Error
	if err != nil {
		logger.Errorf("fail to db.Find(): err=%v", err)
		return nil, err
	}

	return list, nil
}

// GetByID は、指定された条件にマッチする１件のレコードを取得します。条件にマッチするレコードが存在しない時は nil を返します。
func (rp *ContactRepository) GetByID(db *gorm.DB, contactID int) (*model.Contact, error) {
	var contact model.Contact

	db = byContactID(db, contactID)

	err := db.First(&contact).Error
	if err != nil {
		if ErrorIsRecordNotFound(err) {
			return nil, nil
		}

		logger.Errorf("fail to db.Find(): err=%v", err)
		return nil, err
	}

	return &contact, nil
}

// AddOrUpdate は、１件のレコードを追加または更新して、そのレコードを返します。
func (rp *ContactRepository) AddOrUpdate(db *gorm.DB, contact *model.Contact) (*model.Contact, error) {
	err := db.Clauses(clause.OnConflict{UpdateAll: true}).
		Create(contact).Error
	if err != nil {
		logger.Errorf("fail to db.Create(): err=%v", err)
		return nil, err
	}

	return contact, nil
}

// DeleteByID は、指定された条件にマッチする１件のレコードを削除します。条件にマッチするレコードが存在しない時はエラーを返します。
func (rp *ContactRepository) DeleteByID(db *gorm.DB, contactID int) error {
	db = byContactID(db, contactID)

	err := db.Delete(&model.Contact{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Errorf("fail to db.Delete(): err=%v", err)
		return err
	}

	err = ErrorIfRowsAffectedIsNotOne(db.RowsAffected)
	if err != nil {
		logger.Errorf("fail to errorIfRowsAffectedIsNotOne(): err=%v", err)
		return err
	}

	return nil
}
