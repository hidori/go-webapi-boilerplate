package config

import (
	"testing"

	"github.com/hidori/go-webapi-boilerplate/go/test/xtest"
	"github.com/stretchr/testify/assert"
)

func TestNewDBConfig(t *testing.T) {
	type testCase = xtest.Case[any, any, testArgs, *DBConfig]
	tests := []testCase{
		{
			Name: "正常系:DB 構成情報を返す。",
			Args: testArgs{
				getenv: func(s string) string {
					return map[string]string{
						"DB_READER_DSN": "readerDSN",
						"DB_WRITER_DSN": "writerDSN",
					}[s]
				},
			},
			Want: &DBConfig{
				ReaderDSN: "readerDSN",
				WriterDSN: "writerDSN",
			},
		},
	}
	for _, tt := range tests {
		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			got, err := NewDBConfig(tt.Args.getenv)
			if err != nil {
				if tt.WantError {
					assert.Contains(t, err.Error(), tt.Error)
				} else {
					t.Errorf("fail to NewDBConfig(): err=%v, tt.args=%v", err, tt.Args)
				}
			} else {
				xtest.Equal(t, tt.Want, got)
			}
		})
	}
}
