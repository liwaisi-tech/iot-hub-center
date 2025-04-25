package usecases

import "github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/entities"

type TempHumidityUseCase interface {
	Execute(message *entities.TempHumidityEntity) (err error)
}