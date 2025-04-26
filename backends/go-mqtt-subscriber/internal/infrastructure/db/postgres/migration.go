package postgres

import (
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/infrastructure/db/postgres/models"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.TempHumidityModel{},
	)
}
