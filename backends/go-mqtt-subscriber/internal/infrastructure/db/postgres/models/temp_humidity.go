package models

import "time"

type TempHumidityModel struct {
	ID        uint `gorm:"primarykey;autoIncrement;comment:'Primary key'"`
	DeviceMacAddress string `gorm:"not null;index;comment:'Source device MAC address'"`
	Temperature      float64 `gorm:"not null;comment:'Sensored temperature'"`
	Humidity         float64 `gorm:"not null;comment:'Sensored humidity'"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:'Created time'"`
}

func (TempHumidityModel) TableName() string {
	return "temp_humidity"
}
