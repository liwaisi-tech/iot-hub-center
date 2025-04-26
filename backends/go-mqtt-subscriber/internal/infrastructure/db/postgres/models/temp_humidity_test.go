package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestTempHumidityModel(t *testing.T) {

	t.Run("should return the table name", func(t *testing.T) {
		model := TempHumidityModel{}
		tableName := model.TableName()
		assert.Equal(t, tableName, "temp_humidity")
	})

}
