////////////////////////////////////////////////////////////////////////////////
// Author:   Nikita Koryabkin
// Email:    Nikita@Koryabk.in
// Telegram: https://t.me/Apologiz
////////////////////////////////////////////////////////////////////////////////

package Logger

import (
	"alog/Alog"
	"alog/Config"
	"io"
	"sync"
	"time"
)

const (
	keyInfo = "ALOG_LOGGER_INFO"
	keyWrn  = "ALOG_LOGGER_WARNING"
	keyErr  = "ALOG_LOGGER_ERROR"
)

var logger struct {
	instance *Alog.Log
	once     sync.Once
}

func GetLogger() *Alog.Log {
	logger.once.Do(func() {
		logger.instance = Alog.Create(&Alog.Config{
			TimeFormat:  time.RFC3339Nano,
			LogFileLine: false,
			Loggers: Alog.LoggerMap{
				Alog.LoggerInfo: {
					Channel: make(chan string, 100),
					Strategies: []io.Writer{
						Alog.GetFileStrategy(Config.GetEnvStr(keyInfo)),
						Alog.GetDefaultStrategy(),
					},
				},
				Alog.LoggerWrn: {
					Channel: make(chan string, 100),
					Strategies: []io.Writer{
						Alog.GetFileStrategy(Config.GetEnvStr(keyWrn)),
						Alog.GetDefaultStrategy(),
					},
				},
				Alog.LoggerErr: {
					Channel: make(chan string, 100),
					Strategies: []io.Writer{
						Alog.GetFileStrategy(Config.GetEnvStr(keyErr)),
						Alog.GetDefaultStrategy(),
					},
				},
			},
		})
	})
	return logger.instance
}
