package env

import "strconv"

// GetInt は、key に対応する文字列を整数値に変換して返します。文字列が見つからないか空文字列の時はエラーを返します。
func GetInt(getenv Getenv, key string) (int, error) {
	s, err := GetString(getenv, key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(s)
}
