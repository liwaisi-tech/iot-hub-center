package entities

import (
	"regexp"
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/errors"
)

type TempHumidityEntity struct {
	MacAddress string
	Temperature float64
	Humidity float64
}


func (t *TempHumidityEntity) Validate() error {
	macRegex := regexp.MustCompile(`^([0-9A-Fa-f]{2}:){5}[0-9A-Fa-f]{2}$`)
	if !macRegex.MatchString(t.MacAddress) {
		return errors.ErrInvalidMacAddress
	}
	if t.Temperature < -40 || t.Temperature > 125 {
		return errors.ErrInvalidTemperature
	}
	if t.Humidity < 0 || t.Humidity > 100 {
		return errors.ErrInvalidHumidity
	}
	return nil
}
