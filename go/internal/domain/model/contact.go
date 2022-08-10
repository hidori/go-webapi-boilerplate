package model

import "time"

// Contact は、連絡先のドメインモデルです。
type Contact struct {
	// 連絡先 ID
	ContactID int `gorm:"primaryKey"`

	// 姓（漢字）
	FamilyName string

	// 名（漢字）
	FirstName string

	// 姓（カナ）
	FamilyNameKana string

	// 名（カナ）
	FirstNameKana string

	// 電話番号
	PhoneNumber string

	// 郵便番号
	PostalCode string

	// 都道府県コード
	PrefectureCode string

	// 市区町村コード
	CityCode string

	// 町丁目・番地
	AddressLine1 string

	// 建物名・部屋番号
	AddressLine2 string

	// レコード作成時刻
	CreatedAt time.Time `gorm:"type:datetime(6)"`

	// レコード更新時刻
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}

// TableName は、物理テーブル名を返します。
func (m Contact) TableName() string {
	return "contacts"
}

// Validate は、モデルを検証します。
func (m Contact) Validate() ([]string, error) {
	message := []string{}

	// TODO:

	return message, nil
}
