package usecases

import (
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/entities"
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/services"
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/usecases"
)



type TempHumidityUseCaseAdapter struct {
	tempHumidityService services.TempHumidityService
}

func New(
	tempHumidityService services.TempHumidityService,
) usecases.TempHumidityUseCase {
	return &TempHumidityUseCaseAdapter{
		tempHumidityService: tempHumidityService,
	}
}

func (s *TempHumidityUseCaseAdapter) Execute(message *entities.TempHumidityEntity) (err error) {
	err = s.tempHumidityService.ProcessTempHumidityMessage(message)
	return
}
