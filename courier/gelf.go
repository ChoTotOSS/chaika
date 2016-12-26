package courier

import (
	"fmt"
	"time"

	"github.com/duythinht/gelf"
	"github.com/duythinht/gelf/client"
)

type Gelf struct {
	Client      *client.Gelf
	ServiceName string
	Host        string
	Port        int
}

var levels = map[string]int{
	"DEBUG":   0,
	"INFO":    1,
	"WARN":    2,
	"ERROR":   3,
	"FATAL":   4,
	"UNKNOWN": 5,
}

func CreateGelf(serviceName string, graylogHost string, graylogPort int) Courier {
	fmt.Println("Next logs of", serviceName, "will be ship to", graylogHost, graylogPort)
	client := Gelf{
		Client: client.New(client.Config{
			GraylogHost: graylogHost,
			GraylogPort: graylogPort,
		}),
		ServiceName: serviceName,
		Host:        graylogHost,
		Port:        graylogPort,
	}
	return client
}

func (g Gelf) GetHost() string {
	return g.Host
}

func (g Gelf) GetPort() int {
	return g.Port
}

func (g Gelf) Close() {
	g.Client.Close()
}

func (g Gelf) Send(serviceName string, catalog string, level string, message string) {

	host := fmt.Sprintf("[%s][%s]", serviceName, catalog)

	timestamp := time.Now().Unix()

	gm := gelf.Create(message).SetHost(host).SetTimestamp(timestamp)

	if lvlNumber, ok := levels[level]; ok {
		gm.SetLevel(lvlNumber)
	} else {
		gm.SetLevel(5)
	}

	g.Client.Send(gm.ToJSON())
}
