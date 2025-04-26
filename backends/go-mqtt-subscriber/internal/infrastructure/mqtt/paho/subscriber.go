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
}

func New(
	ctx context.Context,
	handlerMessage HandlerMessage,
) *SubscriberHandler {
	return &SubscriberHandler{
		ctx:            ctx,
		handlerMessage: handlerMessage,
	}
}

func (sh *SubscriberHandler) RunConsumer(topic string) {
	qos := byte(0)
	sh.choke = make(chan mqtt.Message)
	clientOptions := getClientOptions()
	clientOptions.SetDefaultPublishHandler(
		func(client mqtt.Client, msg mqtt.Message) {
			sh.choke <- msg
		},
	)
	client := mqtt.NewClient(clientOptions)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Fatalw("Error connecting to MQTT Broker", "error", token.Error())
		panic(token.Error())
	}

	if token := client.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		logger.Fatalw("Error subscribing to MQTT Broker", "error", token.Error())
		panic(token.Error())
	}
	logger.Infow("Subscribed to MQTT Broker", "topic", topic)
	for {
		incoming := <-sh.choke
		logger.Info("Received message", "topic", topic, "message", string(incoming.Payload()))
		err := sh.handlerMessage(sh.ctx, incoming.Payload())
		if err != nil {
			logger.Errorw("Failed processing MQTT message", "error", err)
			continue
		}
		logger.Info("MQTT message processed successfully")
	}
}

func getClientOptions() (clientOptions *mqtt.ClientOptions) {
	clientOptions = mqtt.NewClientOptions()
	clientOptions.AddBroker(os.Getenv("MQTT_BROKER"))
	clientOptions.SetClientID(fmt.Sprintf("%s-%s", os.Getenv("MQTT_CLIENT_ID"), uuid.New().String()))
	if os.Getenv("MQTT_USERNAME") != "" {
		clientOptions.SetUsername(os.Getenv("MQTT_USERNAME"))
		clientOptions.SetPassword(os.Getenv("MQTT_PASSWORD"))
	}
	clientOptions.SetCleanSession(true)
	return
}