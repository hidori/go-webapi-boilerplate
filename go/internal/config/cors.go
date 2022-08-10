package config

import (
	"strings"

	"github.com/hidori/go-webapi-boilerplate/go/pkg/env"
)

const keyCORSAllowOrigins = "CORS_ALLOW_ORIGINS"

// CORSConfig は、CORS 構成情報です。
type CORSConfig struct {
	// 許可オリジン
	AllowOrigins []string
}

// NewCORSConfig は、CORSConfig の新規インスタンスを返します。
func NewCORSConfig(getenv env.Getenv) (*CORSConfig, error) {
	allowOrigins, err := env.GetString(getenv, keyCORSAllowOrigins)
	if err != nil {
		logger.Errorf("fail to env.GetString(): err=%v", err)
		return nil, err
	}

	return &CORSConfig{
		AllowOrigins: strings.Split(allowOrigins, ","),
	}, nil
}
