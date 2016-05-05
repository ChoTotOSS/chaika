package main

import (
	"flag"

	"github.com/duythinht/chaika/chaika"
	"github.com/duythinht/chaika/config"
)

func init() {
	port := flag.Int64("p", 2435, "Port for agent run on")

	consulHost := flag.String("consul-host", "localhost", "Consul hostname")
	consulPort := flag.Int64("consul-port", 8500, "Consul port")

	grayHost := flag.String("graylog-host", "localhost", "Default graylog host or domain")
	grayPort := flag.Int64("graylog-port", 12201, "Default graylog port")
	flag.Parse()
	config.InitConfig(&config.Config{
		Port:        *port,
		ConsulHost:  *consulHost,
		ConsulPort:  *consulPort,
		GraylogHost: *grayHost,
		GraylogPort: *grayPort,
	})
}

func main() {
	chaika.RunServer()
}
