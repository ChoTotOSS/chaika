package chaika

import (
	"fmt"
	"net"
	"strconv"

	"github.com/duythinht/chaika/config"
	"github.com/duythinht/chaika/courier"
)

func RunServer() {
	cfg := config.GetConfig()

	listenAddr := ":" + strconv.FormatInt(cfg.Port, 10)
	ServerAddr, err := net.ResolveUDPAddr("udp", listenAddr)

	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	fmt.Println("Server is up and listen on " + listenAddr)

	CheckError(err)

	defer ServerConn.Close()

	courier.Setup()
	// Buffer for 4KB
	buffer := make([]byte, 4096)

	for {
		// n, add, err
		length, _, err := ServerConn.ReadFromUDP(buffer)

		log := ParseLog(buffer[:length])
		g := courier.Get(log.Service)

		fmt.Println(log.Service, ":", log.Message)
		g.Send(log.Service, log.Catalog, log.Level, log.Message)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
