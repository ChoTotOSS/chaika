package main

import (
	"flag"

	"github.com/duythinht/chaika/chaika"
)

func main() {
	port := flag.Int64("p", 2435, "Port for agent run on")
	consulHost := flag.String("consul-host", "localhost", "Consul hostname")
	consulPort := flag.Int64("consul-port", 8500, "Consul port")
	flag.Parse()
	//flag.Parse()
	chaika.RunServer(*port, *consulHost, *consulPort)
}
