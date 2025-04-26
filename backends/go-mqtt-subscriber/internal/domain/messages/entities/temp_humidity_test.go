package entities

import (
	"testing"

	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/errors"

	"github.com/stretchr/testify/assert"
)


func TestTempHumidityTestSuite(t *testing.T) {
	t.Run("no errors due to valid entity", func(t *testing.T) {
		entity := TempHumidityEntity{
			MacAddress:  "AA:BB:CC:DD:EE:FF",
			Temperature: 25.0,
			Humidity:    50.0,
		}
		err := entity.Validate()
		assert.NoError(t, err)
	})

    t.Run("error due to invalid mac address", func(t *testing.T) {
        entity := TempHumidityEntity{
            MacAddress:  "INVALID-MAC",
            Temperature: 25.0,
            Humidity:    50.0,
        }
        err := entity.Validate()
        assert.Error(t, err)
        assert.Equal(t, errors.ErrInvalidMacAddress, err)
    })

    t.Run("error due to invalid temperature", func(t *testing.T) {
        entity := TempHumidityEntity{
            MacAddress:  "AA:BB:CC:DD:EE:FF",
            Temperature: -50.0,
            Humidity:    50.0,
        }
        err := entity.Validate()
        assert.Error(t, err)
        assert.Equal(t, errors.ErrInvalidTemperature, err)
    })

    t.Run("error due to invalid humidity", func(t *testing.T) {
        entity := TempHumidityEntity{
            MacAddress:  "AA:BB:CC:DD:EE:FF",
            Temperature: 25.0,
            Humidity:    150.0,
        }
        err := entity.Validate()
        assert.Error(t, err)
        assert.Equal(t, errors.ErrInvalidHumidity, err)
    })
}