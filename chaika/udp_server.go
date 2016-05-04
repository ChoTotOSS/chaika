package chaika

import (
	"fmt"
	"net"
	"strconv"

	"github.com/duythinht/chaika/courier"
)

func RunServer(port int64, consulHost string, consulPort int64) {

	listenAddr := ":" + strconv.FormatInt(port, 10)
	ServerAddr, err := net.ResolveUDPAddr("udp", listenAddr)

	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	fmt.Println("Server is up and listen on " + listenAddr)

	CheckError(err)

	defer ServerConn.Close()

	courier.Setup(consulHost, consulPort)

	// Buffer for 4KB
	buffer := make([]byte, 4096)

	for {
		// n, add, err
		length, _, err := ServerConn.ReadFromUDP(buffer)

		log := ParseLog(buffer[:length])
		fmt.Println(log.Service, ":", log.Message)

		g := courier.Get(log.Service)
		g.Send(log.Service, log.Catalog, log.Level, log.Message)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
