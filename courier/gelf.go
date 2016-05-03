package courier

import (
	"strconv"
	"time"

	"github.com/robertkowalski/graylog-golang"
)

type GelfClient struct {
	Client      *gelf.Gelf
	ServiceName string
}

func CreateGelf(serviceName string, graylogHost string, graylogPort int) GelfClient {
	client := GelfClient{
		Client: gelf.New(gelf.Config{
			GraylogPort:     graylogPort,
			GraylogHostname: graylogHost,
		}),
		ServiceName: serviceName,
	}
	return client
}

func (g *GelfClient) Send(serviceName string, catalog string, level string, message string) {
	g.Client.Log(`{
    "host": "[` + serviceName + "][" + catalog + `]",
    "timestamp": ` + strconv.FormatInt(time.Now().Unix(), 10) + `,
    "level": "` + level + `",
    "message": "` + message + `"
  }`)
}
