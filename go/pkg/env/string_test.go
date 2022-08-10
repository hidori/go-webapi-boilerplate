package env

import (
	"testing"

	"github.com/hidori/go-webapi-boilerplate/go/test/xtest"
	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	type testCase = xtest.Case[any, any, testArgs, string]
	tests := []testCase{
		{
			Name: "正常系:key に対応する文字列を返す。",
			Args: testArgs{
				getenv: func(s string) string {
					return map[string]string{
						"KEY1": "AAA",
						"KEY2": "BBB",
						"KEY3": "CCC",
					}[s]
				},
				key: "KEY2",
			},
			Want: "BBB",
		},
		{
			Name: "正常系:key に対応する文字列がダブルクォーテーションで括られて時はダブルクォーテーションを外して返す。",
			Args: testArgs{
				getenv: func(s string) string {
					return map[string]string{
						"KEY1": "AAA",
						"KEY2": "\"BBB\"",
						"KEY3": "CCC",
					}[s]
				},
				key: "KEY2",
			},
			Want: "BBB",
		},
		{
			Name: "異常系:文字列が見つからないか空文字列の時はエラーを返す。",
			Args: testArgs{
				getenv: func(s string) string {
					return map[string]string{
						"KEY1": "AAA",
						"KEY2": "",
						"KEY3": "CCC",
					}[s]
				},
				key: "KEY",
			},
			WantError: true,
			Error:     "getenv(key) returns empty string: key=KEY",
		},
	}
	for _, tt := range tests {
		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			got, err := GetString(tt.Args.getenv, tt.Args.key)
			if err != nil {
				if tt.WantError {
					assert.Contains(t, err.Error(), tt.Error)
				} else {
					t.Errorf("fail to GetString(): err=%v, tt.args=%v", err, tt.Args)
				}
			} else {
				xtest.Equal(t, tt.Want, got)
			}
		})
	}
}
