package config

import (
	"github.com/hidori/go-webapi-boilerplate/go/pkg/env"
)

// Config は、構成情報です。
type Config struct {
	// サーバー構成情報
	Server *ServerConfig

	// DB 構成情報
	DB *DBConfig

	// CORS 構成情報
	CORS *CORSConfig
}

// NewConfig() は、Config の新規インスタンスを返します。
func NewConfig(getenv env.Getenv) (*Config, error) {
	server, err := NewServerConfig(getenv)
	if err != nil {
		logger.Errorf("fail to NewServerConfig(): err=%v", err)
		return nil, err
	}

	db, err := NewDBConfig(getenv)
	if err != nil {
		logger.Errorf("fail to NewDBConfig(): err=%v", err)
		return nil, err
	}

	cors, err := NewCORSConfig(getenv)
	if err != nil {
		logger.Errorf("fail to NewCORSConfig(): err=%v", err)
		return nil, err
	}

	return &Config{
		Server: server,
		DB:     db,
		CORS:   cors,
	}, nil
}
