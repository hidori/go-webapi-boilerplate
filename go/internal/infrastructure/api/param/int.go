package param

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// Int は、パスパラメータを整数値に変換して返します。
func Int(ctx echo.Context, name string) (int, error) {
	s := ctx.Param(name)
	n, err := strconv.Atoi(s)
	if err != nil {
		logger.Errorf("fail to strconv.Atoi(): err=%v", err)
		return 0, err
	}

	return n, nil
}
