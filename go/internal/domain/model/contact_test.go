package model

import (
	"testing"

	"github.com/hidori/go-webapi-boilerplate/go/test/xtest"
)

func TestContact_TableName(t *testing.T) {
	type testArgs struct {
		model Contact
	}
	type testCase = xtest.Case[any, any, testArgs, string]
	tests := []testCase{
		{
			Name: "正常系:物理テーブル名を返す。",
			Args: testArgs{
				model: Contact{},
			},
			Want: "contacts",
		},
	}
	for _, tt := range tests {
		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			got := tt.Args.model.TableName()
			xtest.Equal(t, tt.Want, got)
		})
	}
}
