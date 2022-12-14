package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Open は、新規の DB 接続を返します。
func Open(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), opts...)
	if err != nil {
		logger.Errorf("fail to gorm.Open(): err=%v", err)
		return nil, err
	}

	return db, nil
}
