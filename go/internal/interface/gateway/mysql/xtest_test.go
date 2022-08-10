package mysql

import (
	"fmt"
	"os"

	"github.com/hidori/go-webapi-boilerplate/go/internal/domain/model"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/database"
	"github.com/hidori/go-webapi-boilerplate/go/pkg/env"
	"github.com/hidori/go-webapi-boilerplate/go/test/xtestgorm"
	"gorm.io/gorm"
)

func newTestDB() *gorm.DB {
	dsn, err := env.GetString(os.Getenv, "TEST_DSN_REPOSITORY")
	if err != nil {
		panic(err)
	}

	db, err := database.Open(dsn)
	if err != nil {
		panic(err)
	}

	return db
}

type testContext struct {
	db *gorm.DB
}

func newTestContext() testContext {
	return testContext{
		db: newTestDB(),
	}
}

func e(s string) string {
	return fmt.Sprintf("E%s", s)
}

func deleteAll(db *gorm.DB) error {
	return xtestgorm.DeleteAll(db, []string{
		model.Contact{}.TableName(),
	})
}
