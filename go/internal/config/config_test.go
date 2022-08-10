package config

import (
	"os"
	"testing"

	"github.com/hidori/go-webapi-boilerplate/go/test/xtest"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	type testCase = xtest.Case[any, any, testArgs, *Config]
	tests := []testCase{
		{
			Name: "正常系:構成情報を返す。",
			Args: testArgs{
				getenv: os.Getenv,
			},
			Want: func() *Config {
				server, _ := NewServerConfig(os.Getenv)
				db, _ := NewDBConfig(os.Getenv)
				cors, _ := NewCORSConfig(os.Getenv)
				return &Config{
					Server: server,
					DB:     db,
					CORS:   cors,
				}
			}(),
		},
	}
	for _, tt := range tests {
		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			got, err := NewConfig(tt.Args.getenv)
			if err != nil {
				if tt.WantError {
					assert.Contains(t, err.Error(), tt.Error)
				} else {
					t.Errorf("fail to NewConfig(): err=%v, tt.args=%v", err, tt.Args)
				}
			} else {
				xtest.Equal(t, tt.Want, got)
			}
		})
	}
}
