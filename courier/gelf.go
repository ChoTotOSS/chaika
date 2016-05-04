package courier

import (
	"strconv"
	"time"

	"github.com/robertkowalski/graylog-golang"
)

type Gelf struct {
	Client      *gelf.Gelf
	ServiceName string
}

func CreateGelf(serviceName string, graylogHost string, graylogPort int64) Courier {
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
	g.Client.Log(`{
    "host": "[` + serviceName + "][" + catalog + `]",
    "timestamp": ` + strconv.FormatInt(time.Now().Unix(), 10) + `,
    "level": "` + level + `",
    "message": "` + message + `"
  }`)
}
