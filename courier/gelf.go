package courier

import (
	"fmt"
	"strconv"
	"time"

	"github.com/robertkowalski/graylog-golang"
)

type Gelf struct {
	Client      *gelf.Gelf
	ServiceName string
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
	logData := `{
    "host": "[` + serviceName + "][" + catalog + `]",
    "timestamp": ` + strconv.FormatInt(time.Now().Unix(), 10) + `,
    "message": "` + message + `"
  }`
	g.Client.Log(logData)
}
