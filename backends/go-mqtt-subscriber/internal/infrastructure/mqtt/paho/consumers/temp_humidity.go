package consumers

import (
	"context"
	"encoding/json"

	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/adapters/messages/services"
	ucTempHumidity "github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/application/usecases/messages"
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/entities"
	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/domain/messages/usecases"
	repositories "github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/infrastructure/db/postgres/repositories/messages"
	subscriber "github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/infrastructure/mqtt/paho"
	pkgGorm "github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/pkg/gorm/postgres"
	pkgZap "github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/pkg/zap"
	"go.uber.org/zap"
)

type TempHumidityConsumer struct {
	topic   string
	handler *subscriber.SubscriberHandler
	useCase usecases.TempHumidityUseCase
	logger  *zap.SugaredLogger
}

func newUseCase(logger *zap.SugaredLogger) usecases.TempHumidityUseCase {
	gormDB, err := pkgGorm.GetPostgresDB()
	if err != nil {
		logger.Fatalw("Failed to get Postgres DB", "error", err)
	}
	tempHumidityRepository := repositories.New(gormDB)
	tempHumidityService := services.New(tempHumidityRepository)
	return ucTempHumidity.New(tempHumidityService)
}
const (
	tempHumidityTopic = "liwaisi-iot/farm/temp_humidity"
)

func NewTempHumidityConsumer() *TempHumidityConsumer {
	logger := pkgZap.GetLogger()
	useCase := newUseCase(logger)
	consumer := &TempHumidityConsumer{
		topic:  tempHumidityTopic,
		useCase: useCase,
		logger: logger,
	}
	consumer.handler = subscriber.New(
		context.Background(),
		consumer.handleMessage,
	)
	return consumer
}

func (s *TempHumidityConsumer) Run() {
	s.logger.Infow("Running TempHumidityConsumer", "topic", s.topic)
	s.handler.RunConsumer(s.topic)
}

type TempHumidityMessage struct {
	Mac      string  `json:"mac"`
	Temp     float64 `json:"temperature"`
	Humidity float64 `json:"humidity"`
}

func (s *TempHumidityConsumer) handleMessage(ctx context.Context, eventMessage []byte) (err error) {
	var msg TempHumidityMessage
	if err = json.Unmarshal(eventMessage, &msg); err != nil {
		s.logger.Errorw("Failed to unmarshal temp humidity message", "error", err)
		return
	}
	entity := &entities.TempHumidityEntity{
		MacAddress:  msg.Mac,
		Temperature: msg.Temp,
		Humidity:    msg.Humidity,
	}
	if err = entity.Validate(); err != nil {
		s.logger.Errorw("Invalid temp humidity message", "error", err)
		return
	}
	if err = s.useCase.Execute(entity); err != nil {
		s.logger.Errorw("Failed to process temp humidity message", "error", err)
		return
	}
	s.logger.Info("Temp humidity message processed successfully")
	return
}