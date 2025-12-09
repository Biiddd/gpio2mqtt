package main

import (
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func startDOSubscribers() {
	for name, d := range config.App.Do {
		name := name
		d := d

		// 订阅 CmdTopic
		if d.CmdTopic != "" {
			logger.Infof("DO[%s] 订阅高电平 topic=%s payload=%s", name, d.CmdTopic, d.HighPayload)
			tok := mqttClient.Subscribe(d.CmdTopic, 0, func(c mqtt.Client, m mqtt.Message) {
				payload := strings.TrimSpace(string(m.Payload()))
				if d.HighPayload == "" || payload == d.HighPayload {
					logger.Infof("DO[%s] 收到高电平指令: topic=%s payload=%s", name, m.Topic(), payload)
					if err := updateDo(d.Path, 1); err != nil {
						logger.Errorf("DO[%s] 置高失败: %v", name, err)
					} else {
						sendMqttMsg(d.StatusTopic, "1", 0, false)
					}
				}

				if d.LowPayload == "" || payload == d.LowPayload {
					logger.Infof("DO[%s] 收到低电平指令: topic=%s payload=%s", name, m.Topic(), payload)
					if err := updateDo(d.Path, 0); err != nil {
						logger.Errorf("DO[%s] 置低失败: %v", name, err)
					} else {
						sendMqttMsg(d.StatusTopic, "0", 0, false)
					}
				}
			})
			tok.Wait()
			if tok.Error() != nil {
				logger.Errorf("订阅 DO[%s].CmdTopic 失败: %v", name, tok.Error())
			}
		}
	}
}
