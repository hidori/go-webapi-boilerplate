package mysql

import "gorm.io/gorm"

// byContactID は、contactID による絞り込み条件を db に適用します。
func byContactID(db *gorm.DB, contactID int) *gorm.DB {
	return db.Where("contact_id = ?", contactID)
}
