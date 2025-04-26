package services

import(
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/entities"
)

type TempHumidityService interface {
	ProcessTempHumidityMessage(message *entities.TempHumidityEntity) error
}
