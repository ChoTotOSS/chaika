package main

import (
	"flag"
	"fmt"

	"github.com/duythinht/chaika/chaika"
	"github.com/duythinht/chaika/config"
)

var (
	requestVersionInfo *bool
)

func init() {
	port := flag.Int64("p", 2435, "Port for agent run on")

	consulHost := flag.String("consul-host", "localhost", "Consul hostname")
	consulPort := flag.Int64("consul-port", 8500, "Consul port")

	grayHost := flag.String("graylog-host", "localhost", "Default graylog host or domain")
	grayPort := flag.Int64("graylog-port", 12201, "Default graylog port")
	requestVersionInfo = flag.Bool("version", false, "Version")
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
	if *requestVersionInfo {
		fmt.Println("v" + VERSION)
		return
	}
	chaika.RunServer()
}
