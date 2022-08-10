package log

import (
	"testing"

	"github.com/hidori/go-webapi-boilerplate/go/test/xtest"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestGetLogger(t *testing.T) {
	type testCase = xtest.Case[any, any, any, *log.Logger]
	tests := []testCase{
		{
			Name: "正常系:Logger のインスタンスを返す。",
		},
	}
	for _, tt := range tests {
		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			got := GetLogger()
			assert.NotNil(t, got)
		})
	}
}
