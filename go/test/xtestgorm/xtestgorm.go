package xtestgorm

import (
	"fmt"

	"gorm.io/gorm"
)

// DeleteAll deletes all records in the specified tables.
func DeleteAll(db *gorm.DB, tableNames []string) error {
	for i := range tableNames {
		sql := fmt.Sprintf("DELETE FROM `%s` WHERE 1=1;", tableNames[i])
		err := db.Exec(sql).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// InsertAll inserts all records into the specified tables.
func InsertAll(db *gorm.DB, models []interface{}) error {
	for i := range models {
		err := db.Create(models[i]).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// SelectAll reads all records from the specified tables.
func SelectAll[TModel any](db *gorm.DB) ([]TModel, error) {
	var out []TModel

	err := db.Find(&out).Error
	if err != nil {
		return nil, err
	}

	return out, nil
}
