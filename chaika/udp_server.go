package chaika

import (
	"fmt"
	"net"
)

func UDPServerRun() {
	ServerAddr, err := net.ResolveUDPAddr("udp", ":2435")

	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)

	CheckError(err)

	defer ServerConn.Close()

	// Buffer for 4KB
	buffer := make([]byte, 4096)

	for {
		// n, add, err
		length, _, err := ServerConn.ReadFromUDP(buffer)
		fmt.Println("JSON ", string(buffer[0:length]))
		log := ParseLog(buffer[:length])
		fmt.Println(log.Service)
		fmt.Println(log.Message)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
