package zap

import (
	"os"
	"sync"

	"go.uber.org/zap"
)

var( 
	instance *zap.SugaredLogger
	once sync.Once
)

func GetLogger() *zap.SugaredLogger {
	if instance != nil {
		return instance
	}
	once.Do(func() {
		var logger *zap.Logger
		if os.Getenv("ENV") == "prod" {
			logger, _ = zap.NewProductionConfig().Build()
		} else {
			logger, _ = zap.NewDevelopmentConfig().Build()
		}
		defer logger.Sync()
		instance = logger.Sugar()
	})
	return instance
}