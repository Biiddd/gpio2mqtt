package main

func sendMqttMsg(topic string, msg string, qos byte, retained bool) {
	if mqttClient == nil {
		logger.Errorf("MQTT client not initialized, cannot publish topic=%s", topic)
		return
	}
	token := mqttClient.Publish(topic, qos, retained, msg)
	token.Wait()
	if token.Error() != nil {
		logger.Errorf("MQTT publish failed! topic=%s err=%v", topic, token.Error())
	}
}
