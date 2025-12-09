package main

import (
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

var mqttClient mqtt.Client

func main() {
	initConfig()
	logger = initLogger(config.Log)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.MQTT.Broker)
	opts.SetClientID(config.MQTT.ClientID)
	opts.SetUsername(config.MQTT.User)
	opts.SetPassword(config.MQTT.Password)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)

	client := mqtt.NewClient(opts)
	mqttClient = client
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Fatalf("failed to connect to mqtt broker: %v", token.Error())
		os.Exit(1)
	}

	startDIWatchers()
	startDOSubscribers()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	logger.Infof("receive shutdown signal, shutting down...")
	mqttClient.Disconnect(250)
}
