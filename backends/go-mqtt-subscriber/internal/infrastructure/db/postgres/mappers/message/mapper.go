package mappers

import (
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/entities"
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/infrastructure/db/postgres/models"
)


func MapTempHumidityEntityToDBModel(entity *entities.TempHumidityEntity) *models.TempHumidityModel {
	if entity == nil {
		return nil
	}
	return &models.TempHumidityModel{
		DeviceMacAddress: entity.MacAddress,
		Temperature:      entity.Temperature,
		Humidity:         entity.Humidity,
	}
}

func MapDBModelToTempHumidityEntity(dbModel *models.TempHumidityModel) *entities.TempHumidityEntity {
	if dbModel == nil {
		return nil
	}
	return &entities.TempHumidityEntity{
		MacAddress: dbModel.DeviceMacAddress,
		Temperature: dbModel.Temperature,
		Humidity: dbModel.Humidity,
	}
}

func MapDBModelsToTempHumidityEntities(dbModels []*models.TempHumidityModel) []*entities.TempHumidityEntity {
	if dbModels == nil {
		return nil
	}
	var entities []*entities.TempHumidityEntity
	for _, dbModel := range dbModels {
		entities = append(entities, MapDBModelToTempHumidityEntity(dbModel))
	}
	return entities
}