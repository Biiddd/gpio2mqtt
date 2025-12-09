package main

import (
	"fmt"
	"time"
)

func startDIWatchers() {
	for name, di := range config.App.Di {
		name := name
		di := di

		go func() {
			logger.Infof("DI [%s] start monitor path=%s interval=%dms topic=%s",
				name, di.Path, di.Interval, di.StatusTopic)

			var last = -1
			interval := time.Duration(di.Interval) * time.Millisecond

			for {
				cur, err := pullStatus(di.Path)
				if err != nil {
					logger.Errorf("pull DI[%s] status failed: %v", name, err)
					time.Sleep(time.Second)
					continue
				}
				if cur != last {
					logger.Infof("DI[%s] changed: %d -> %d", name, last, cur)
					last = cur

					sendMqttMsg(di.StatusTopic, fmt.Sprintf("%d", cur), 0, false)
					logger.Infof("sending DI[%s] status=%d to topic=%s", name, cur, di.StatusTopic)
				}
				time.Sleep(interval)
			}
		}()
	}
}
