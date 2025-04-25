package mappers

import (
	"testing"
	"time"

	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/entities"
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/infrastructure/db/postgres/models"
	"github.com/stretchr/testify/assert"
)


func TestTempHumidityMappers(t *testing.T) {

	t.Run("mapping nil entity to db model", func(t *testing.T) {
		result := MapTempHumidityEntityToDBModel(nil)

		assert.Nil(t, result)
	})

	t.Run("mapping valid entity to db model", func(t *testing.T) {
		entity := &entities.TempHumidityEntity{
			MacAddress: "123456789012",
			Temperature: 25.5,
			Humidity: 50.2,
		}
		result := MapTempHumidityEntityToDBModel(entity)

		assert.NotNil(t, result)
		assert.Equal(t, entity.MacAddress, result.DeviceMacAddress)
		assert.Equal(t, entity.Temperature, result.Temperature)
		assert.Equal(t, entity.Humidity, result.Humidity)
		assert.Equal(t, result.CreatedAt, time.Time{})
	})

	t.Run("mapping nil db model to entity", func(t *testing.T) {
		result := MapDBModelToTempHumidityEntity(nil)

		assert.Nil(t, result)
	})

	t.Run("mapping valid db model to entity", func(t *testing.T) {
		dbModel := &models.TempHumidityModel{
			DeviceMacAddress: "123456789012",
			Temperature:      25.5,
			Humidity:         50.2,
		}
		result := MapDBModelToTempHumidityEntity(dbModel)

		assert.NotNil(t, result)
		assert.Equal(t, dbModel.DeviceMacAddress, result.MacAddress)
		assert.Equal(t, dbModel.Temperature, result.Temperature)
		assert.Equal(t, dbModel.Humidity, result.Humidity)
	})

}