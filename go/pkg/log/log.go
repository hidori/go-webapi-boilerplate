package log

import (
	"sync"

	"github.com/labstack/gommon/log"
)

var logger *log.Logger

var mutex sync.Mutex

// GetLogger は、Logger のキャシュまたは新規インスタンスを返します。
func GetLogger() *log.Logger {
	mutex.Lock()
	defer mutex.Unlock()

	if logger == nil {
		logger = log.New("-")
	}

	return logger
}
