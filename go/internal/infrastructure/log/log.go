package log

import (
	"sync"

	"github.com/labstack/gommon/log"
)

var mutex sync.Mutex

var logger *log.Logger

// GetLogger は、Logger のキャッシュまたは新規インスタンスを返します。
func GetLogger() *log.Logger {
	mutex.Lock()
	defer mutex.Unlock()

	if logger == nil {
		logger = log.New("webapi")
	}

	return logger
}
