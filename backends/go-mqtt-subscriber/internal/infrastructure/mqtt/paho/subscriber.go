package handlers

import (
	"context"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	pkgZap "github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/pkg/zap"
)

var logger = pkgZap.GetLogger()

type HandlerMessage func(ctx context.Context, eventMessage []byte) (err error)

type SubscriberHandler struct {
	ctx            context.Context
	handlerMessage HandlerMessage
	choke          chan mqtt.Message
	workers        int
	client         mqtt.Client
	cancel         context.CancelFunc
}

// New creates a new SubscriberHandler.
// workers: number of concurrent message processors (default 4 if 0)
func New(
	ctx context.Context,
	handlerMessage HandlerMessage,
	workers int,
) *SubscriberHandler {
	if workers <= 0 {
		workers = 4
	}
	ctx, cancel := context.WithCancel(ctx)
	return &SubscriberHandler{
		ctx:            ctx,
		handlerMessage: handlerMessage,
		workers:        workers,
		choke:          make(chan mqtt.Message, 100), // buffered channel
		cancel:         cancel,
	}
}

// RunConsumer connects to the MQTT broker, subscribes, and processes messages with a worker pool.
func (sh *SubscriberHandler) RunConsumer(topic string) error {
	qos := byte(0)
	clientOptions, err := sh.getClientOptions()
	if err != nil {
		logger.Fatalw("Invalid MQTT configuration", "error", err)
		return err
	}
	clientOptions.SetDefaultPublishHandler(
		func(client mqtt.Client, msg mqtt.Message) {
			select {
			case sh.choke <- msg:
				// delivered
			default:
				logger.Warnw("Message dropped due to full channel buffer", "topic", msg.Topic())
			}
		},
	)
	client := mqtt.NewClient(clientOptions)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Fatalw("Error connecting to MQTT Broker", "error", token.Error())
		return token.Error()
	}
	sh.client = client
	if token := client.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		logger.Fatalw("Error subscribing to MQTT Broker", "error", token.Error())
		client.Disconnect(250)
		return token.Error()
	}
	logger.Infow("Subscribed to MQTT Broker",
		"topic", topic,
		"ClientID", clientOptions.ClientID)

	// Start worker pool
	for i := 0; i < sh.workers; i++ {
		go sh.workerLoop(topic, i)
	}

	// Wait for context cancellation
	<-sh.ctx.Done()
	logger.Info("Context cancelled, shutting down MQTT consumer...")
	client.Unsubscribe(topic)
	client.Disconnect(250)
	return nil
}

// workerLoop processes messages from the channel
func (sh *SubscriberHandler) workerLoop(topic string, workerID int) {
	for {
		select {
		case <-sh.ctx.Done():
			logger.Infow("Worker exiting", "workerID", workerID)
			return
		case incoming := <-sh.choke:
			logger.Infow("Received message", "workerID", workerID, "topic", topic, "message", string(incoming.Payload()))
			err := sh.handlerMessage(sh.ctx, incoming.Payload())
			if err != nil {
				logger.Errorw("Failed processing MQTT message", "workerID", workerID, "error", err)
				continue
			}
			logger.Infow("MQTT message processed successfully", "workerID", workerID)
		}
	}
}

// getClientOptions builds and validates MQTT client options from environment variables.
func (sh *SubscriberHandler) getClientOptions() (*mqtt.ClientOptions, error) {
	broker := os.Getenv("MQTT_BROKER")
	clientID := os.Getenv("MQTT_CLIENT_ID")
	if broker == "" || clientID == "" {
		return nil, fmt.Errorf("missing MQTT_BROKER or MQTT_CLIENT_ID environment variable")
	}
	clientOptions := mqtt.NewClientOptions()
	clientOptions.AddBroker(broker)
	clientOptions.SetClientID(fmt.Sprintf("%s-%s", clientID, uuid.New().String()))
	username := os.Getenv("MQTT_USERNAME")
	if username != "" {
		clientOptions.SetUsername(username)
		clientOptions.SetPassword(os.Getenv("MQTT_PASSWORD"))
	}
	clientOptions.SetCleanSession(true)
	return clientOptions, nil
}