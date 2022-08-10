package config

import (
	"testing"

	"github.com/hidori/go-webapi-boilerplate/go/test/xtest"
	"github.com/stretchr/testify/assert"
)

func TestNewServerConfig(t *testing.T) {
	type testCase = xtest.Case[any, any, testArgs, *ServerConfig]
	tests := []testCase{
		{
			Name: "正常系:サーバー構成情報を返す。",
			Args: testArgs{
				getenv: func(s string) string {
					return map[string]string{
						"SERVER_PORT": "1234",
					}[s]
				},
			},
			Want: &ServerConfig{
				Port: 1234,
			},
		},
	}
	for _, tt := range tests {
		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			got, err := NewServerConfig(tt.Args.getenv)
			if err != nil {
				if tt.WantError {
					assert.Contains(t, err.Error(), tt.Error)
				} else {
					t.Errorf("fail to NewServerConfig(): err=%v, tt.args=%v", err, tt.Args)
				}
			} else {
				xtest.Equal(t, tt.Want, got)
			}
		})
	}
}
