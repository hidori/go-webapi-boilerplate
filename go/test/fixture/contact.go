package fixture

import "github.com/hidori/go-webapi-boilerplate/go/internal/domain/model"

// NewContact は、Contact のフィクスチャを返します。
func NewContact(contactID int) *model.Contact {
	return &model.Contact{
		ContactID:      contactID,
		FamilyName:     "姓（漢字）",
		FirstName:      "名（漢字）",
		FamilyNameKana: "姓（カナ）",
		FirstNameKana:  "名（カナ）",
		PhoneNumber:    "12345678901",
		PostalCode:     "1234567",
		PrefectureCode: "13",
		CityCode:       "01",
		AddressLine1:   "住所１",
		AddressLine2:   "住所２",
	}
}
