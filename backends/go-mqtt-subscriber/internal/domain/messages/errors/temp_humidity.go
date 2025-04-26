package errors

import (
	"errors"
)

var(
	ErrInvalidMacAddress = errors.New("Invalid MAC address")
	ErrInvalidTemperature = errors.New("Invalid temperature")
	ErrInvalidHumidity = errors.New("Invalid humidity")
)