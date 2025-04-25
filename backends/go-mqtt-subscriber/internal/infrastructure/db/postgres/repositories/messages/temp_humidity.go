package repositories

import (
	"gorm.io/gorm"

	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/entities"
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/repositories"
	errors "github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/infrastructure/db/postgres/errors/messages"
	mappers "github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/infrastructure/db/postgres/mappers/message"
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/infrastructure/db/postgres/models"
)

type TempHumidityPostgresRepository struct {
	db *gorm.DB
}

func New(
	db *gorm.DB,
) repositories.TempHumidityRepository {
	return &TempHumidityPostgresRepository{
		db: db,
	}
}

// Save implements repositories.TempHumidityRepository
func (r *TempHumidityPostgresRepository) Save(entity *entities.TempHumidityEntity) error {
	if entity == nil {
		return errors.ErrInvalidTempHumidityEntity
	}
	dbModel := mappers.MapTempHumidityEntityToDBModel(entity)
	if err := r.db.Create(dbModel).Error; err != nil {
		return err
	}
	return nil
}

// GetLatestRecord implements repositories.TempHumidityRepository
func (r *TempHumidityPostgresRepository) GetLatestRecord(macAddress string) (*entities.TempHumidityEntity, error) {
	var dbModel *models.TempHumidityModel
	if err := r.db.Where("device_mac_address = ?", macAddress).Order("created_at desc").First(&dbModel).Error; err != nil {
		return nil, err
	}
	return mappers.MapDBModelToTempHumidityEntity(dbModel), nil
}

// FindPaginated implements repositories.TempHumidityRepository
func (r *TempHumidityPostgresRepository) FindPaginated(macAddress string, limit int, offset int) ([]*entities.TempHumidityEntity, error) {
	var dbModels []*models.TempHumidityModel
	if err := r.db.Where("device_mac_address = ?", macAddress).Order("created_at desc").Offset(offset).Limit(limit).Find(&dbModels).Error; err != nil {
		return nil, err
	}
	return mappers.MapDBModelsToTempHumidityEntities(dbModels), nil
}

