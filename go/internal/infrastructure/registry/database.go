package registry

import (
	"gorm.io/gorm"
)

// Database は、DB 接続のレジストリです。
type Database struct {
	// 読み取り専用エンドポイントの DB 接続
	Reader *gorm.DB

	// 書き込み可能エンドポイントの DB 接続
	Writer *gorm.DB
}
