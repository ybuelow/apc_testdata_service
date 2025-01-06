package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

type client struct {
	mqttClient mqtt.Client
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	//fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func newClient() (*client, error) {
	err := godotenv.Load(".env")
	check(err)

	broker := os.Getenv("MQTT_BROKER")
	envport := os.Getenv("PORT")
	port, porterr := strconv.ParseInt(envport, 10, 16)

	opts := mqtt.NewClientOptions()
	check(porterr)

	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(os.Getenv("CLIENT_ID"))
	opts.SetUsername(os.Getenv("USER_NAME"))
	opts.SetPassword(os.Getenv("USER:PASSWORD"))
	opts.SetDefaultPublishHandler(messagePubHandler)
	//opts.WebsocketOptions.Proxy() //TODO:

	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return &client{mqttClient}, nil
}

func (c *client) Publish(topic string, interval int) error {
	for e := 0; e < 12; e++ {
		telemetry, err := json.Marshal(insertTelemetry(e, 11))
		check(err)
		token := c.mqttClient.Publish(topic, 1, false, telemetry)
		token.Wait()
		time.Sleep(time.Duration(interval) * time.Second)
	}
	return nil
}

func (c *client) sub(topic string, f mqtt.MessageHandler) error {
	if token := c.mqttClient.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	fmt.Println("token")
	return nil
}
