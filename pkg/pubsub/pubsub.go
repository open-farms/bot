package pubsub

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/open-farms/bot/internal/logger"
)

const (
	PublicBroker = "broker.emqx.io"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	logger.Log.Info().Str("topic", msg.Topic()).Bytes("payload", msg.Payload()).Msg("received message")
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	logger.Log.Info().Msgf("connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	logger.Log.Error().Err(err).Msg("connection lost")
}

type Client struct {
	client mqtt.Client
}

type PublishHandler func(topic string, payloads ...interface{})

type SubscribeHandler func(topic string, handle mqtt.MessageHandler)

func NewClient(broker string, port int) (*Client, error) {
	opts := mqtt.NewClientOptions()
	conn := fmt.Sprintf("tcp://%s:%d", broker, port)
	logger.Log.Info().Str("broker", conn).Send()
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	opts.AddBroker(conn)
	opts.SetClientID(fmt.Sprintf("%s_%s", "go_mqtt_client", id.String()))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Client{client: client}, nil
}

func (c *Client) Disconnect(delay uint, done chan bool) {
	c.client.Disconnect(delay)
	done <- true
}

func (c *Client) Publish(topic string, payloads ...interface{}) {
	for _, payload := range payloads {
		token := c.client.Publish(topic, 0, false, payload)
		token.Wait()
		logger.Log.Info().Str("topic", topic).Interface("payload", payload).Msg("sent")
	}
}

func (c *Client) Subscribe(topic string, handle mqtt.MessageHandler) {
	token := c.client.Subscribe(topic, 1, handle)
	token.Wait()
	logger.Log.Info().Str("topic", topic).Msgf("subscribed")
}
