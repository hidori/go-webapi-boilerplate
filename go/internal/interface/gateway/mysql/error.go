package mysql

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// ErrorIsRecordNotFound は、err が gorm.ErrRecordNotFond の時は true を、それ以外の時は false を返します。
func ErrorIsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// ErrorIfRowsAffectedIsNotOne は、rowsAffected が 1 の時は true を、それ以外の時は false を返します。
func ErrorIfRowsAffectedIsNotOne(rowsAffected int64) error {
	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected is not 1: rowsAffected=%d", rowsAffected)
	}

	return nil
}
