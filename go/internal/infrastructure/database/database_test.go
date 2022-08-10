package database

import (
	"os"
	"testing"

	"github.com/hidori/go-webapi-boilerplate/go/internal/config"
	"github.com/hidori/go-webapi-boilerplate/go/test/xtest"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestOpen(t *testing.T) {
	config, err := config.NewConfig(os.Getenv)
	if err != nil {
		t.Errorf("fail to config.NewConfig(): err=%v", err)
		return
	}
	type testArgs struct {
		dsn string
	}
	type testCase = xtest.Case[any, any, testArgs, *gorm.DB]
	tests := []testCase{
		{
			Name: "正常系:新規 DB 接続を返す。",
			Args: testArgs{
				dsn: config.DB.ReaderDSN,
			},
		},
		{
			Name: "異常系:DB 接続に失敗した時はエラーを返す。",
			Args: testArgs{
				dsn: "tcp(no-host)/no-database",
			},
			WantError: true,
			Error:     "lookup no-host",
		},
	}
	for _, tt := range tests {
		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			got, err := Open(tt.Args.dsn)
			if err != nil {
				if tt.WantError {
					assert.Contains(t, err.Error(), tt.Error)
				} else {
					t.Errorf("fail to NewConfig(): err=%v, tt.args=%v", err, tt.Args)
				}
			} else {
				assert.NotNil(t, got)
			}
		})
	}
}
