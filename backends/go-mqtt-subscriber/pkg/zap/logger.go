package zap

import (
	"os"

	"go.uber.org/zap"
)

var intance *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if intance != nil {
		return intance
	}
	var logger *zap.Logger
	if os.Getenv("ENV") == "prod" {
		logger, _ = zap.NewProductionConfig().Build()
	} else {
		logger, _ = zap.NewDevelopmentConfig().Build()
	}
	defer logger.Sync()
	intance = logger.Sugar()
	return intance
}