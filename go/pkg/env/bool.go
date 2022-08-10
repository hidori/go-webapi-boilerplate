package env

import "strconv"

// GetBool は、key に対応する文字列を真理値に変換して返します。文字列が見つからないか空文字列の時はエラーを返します。
func GetBool(getenv Getenv, key string) (bool, error) {
	s, err := GetString(getenv, key)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(s)
}
