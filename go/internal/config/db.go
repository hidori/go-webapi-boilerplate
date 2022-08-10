package config

import "github.com/hidori/go-webapi-boilerplate/go/pkg/env"

const (
	keyDBReaderDSN = "DB_READER_DSN"
	keyDBWriterDSN = "DB_WRITER_DSN"
)

// DBConfig は、DB 構成情報です。
type DBConfig struct {
	// 読み取り専用 DB エンドポイントの接続情報
	ReaderDSN string

	// 書き込み可能 DB エンドポイントの接続情報
	WriterDSN string
}

// NewDBConfig は、DBConfig の新規インスタンスを返します。
func NewDBConfig(getenv env.Getenv) (*DBConfig, error) {
	readerDSN, err := env.GetString(getenv, keyDBReaderDSN)
	if err != nil {
		logger.Errorf("fail to env.GetString(): err=%v", err)
		return nil, err
	}

	writerDSN, err := env.GetString(getenv, keyDBWriterDSN)
	if err != nil {
		logger.Errorf("fail to env.GetString(): err=%v", err)
		return nil, err
	}

	return &DBConfig{
		ReaderDSN: readerDSN,
		WriterDSN: writerDSN,
	}, nil
}
