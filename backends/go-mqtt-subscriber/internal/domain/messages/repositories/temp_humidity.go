package repositories

import (
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/entities"
)

type TempHumidityRepository interface {
	Save(entity *entities.TempHumidityEntity) error
	GetLatestRecord(macAddress string) (*entities.TempHumidityEntity, error)
	FindPaginated(macAddress string, limit int, offset int) ([]*entities.TempHumidityEntity, error)
}
