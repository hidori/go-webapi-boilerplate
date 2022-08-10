package env

import (
	"testing"

	"github.com/hidori/go-webapi-boilerplate/go/test/xtest"
	"github.com/stretchr/testify/assert"
)

func TestGetInt(t *testing.T) {
	type testCase = xtest.Case[any, any, testArgs, int]
	tests := []testCase{
		{
			Name: "正常系:key に対応する文字列を真理値に変換して返す。",
			Args: testArgs{
				getenv: func(s string) string {
					return "123"
				},
				key: "KEY",
			},
			Want: 123,
		},
		{
			Name: "異常系:文字列が見つからないか空文字列の時はエラーを返す。",
			Args: testArgs{
				getenv: func(s string) string {
					return ""
				},
				key: "KEY",
			},
			WantError: true,
			Error:     "getenv(key) returns empty string: key=KEY",
		},
		{
			Name: "異常系:変換に失敗した時はエラーを返す。",
			Args: testArgs{
				getenv: func(s string) string {
					return "AAA"
				},
				key: "KEY",
			},
			WantError: true,
			Error:     "strconv.Atoi: parsing \"AAA\": invalid syntax",
		},
	}
	for _, tt := range tests {
		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			got, err := GetInt(tt.Args.getenv, tt.Args.key)
			if err != nil {
				if tt.WantError {
					assert.Contains(t, err.Error(), tt.Error)
				} else {
					t.Errorf("fail to GetInt(): err=%v, tt.args=%v", err, tt.Args)
				}
			} else {
				xtest.Equal(t, tt.Want, got)
			}
		})
	}
}
