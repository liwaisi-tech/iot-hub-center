package services

import (
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/entities"
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/repositories"
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/services"
)

type TempHumidityServiceAdapter struct {
	tempHumidityRepository repositories.TempHumidityRepository
}

func New(
	tempHumidityRepository repositories.TempHumidityRepository,
) services.TempHumidityService {
	return &TempHumidityServiceAdapter{
		tempHumidityRepository: tempHumidityRepository,
	}
}

func (s *TempHumidityServiceAdapter) ProcessTempHumidityMessage(message *entities.TempHumidityEntity) (err error) {
	err = s.tempHumidityRepository.Save(message)
	return
}
