package config

import (
	"github.com/hidori/go-webapi-boilerplate/go/pkg/env"
)

const (
	keyServerPort = "SERVER_PORT"
)

// ServerConfig は、サーバー構成情報です。
type ServerConfig struct {
	// ポート番号
	Port int
}

// NewServerConfig は、ServerConfig の新規インスタンスを返します。
func NewServerConfig(getenv env.Getenv) (*ServerConfig, error) {
	port, err := env.GetInt(getenv, keyServerPort)
	if err != nil {
		logger.Errorf("fail to env.GetInt(): err=%v", err)
		return nil, err
	}

	return &ServerConfig{
		Port: port,
	}, nil
}
