package env

import (
	"fmt"
	"regexp"
)

var doubleQuotedValuePattern = regexp.MustCompile(`^".+"$`)

// GetString は、key に対応する文字列を返します。文字列が見つからないか空文字列の時はエラーを返します。
func GetString(getenv Getenv, key string) (string, error) {
	s := getenv(key)
	if s == "" {
		return "", fmt.Errorf("getenv(key) returns empty string: key=%s", key)
	}

	if doubleQuotedValuePattern.MatchString(s) {
		s = s[1 : len(s)-1]
	}

	return s, nil
}
