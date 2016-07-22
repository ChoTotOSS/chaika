package courier

import (
	"encoding/json"
	"fmt"
	"github.com/robertkowalski/graylog-golang"
	"time"
)

type Gelf struct {
	Client      *gelf.Gelf
	ServiceName string
}

var levels = map[string]int32{
	"DEBUG":   0,
	"INFO":    1,
	"WARN":    2,
	"ERROR":   3,
	"FATAL":   4,
	"UNKNOWN": 5,
}

func CreateGelf(serviceName string, graylogHost string, graylogPort int64) Courier {
	fmt.Println("Next logs of", serviceName, "will be ship to", graylogHost, graylogPort)
	client := Gelf{
		Client: gelf.New(gelf.Config{
			GraylogPort:     int(graylogPort),
			GraylogHostname: graylogHost,
		}),
		ServiceName: serviceName,
	}
	return client
}

func (g Gelf) Send(serviceName string, catalog string, level string, message string) {
	logObj := map[string]interface{}{
		"host":      "[" + serviceName + "][" + catalog + "]",
		"timestamp": time.Now().Unix(),
		"message":   message,
	}

	if lvlNumber, ok := levels[level]; ok {
		logObj["level"] = lvlNumber
	} else {
		logObj["level"] = 5
	}

	logBuff, err := json.Marshal(logObj)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		g.Client.Log(string(logBuff))
	}
}
