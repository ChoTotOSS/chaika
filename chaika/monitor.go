package chaika

import (
	"fmt"
	"net"
	"os"
)

var handlers []net.Conn

func RunMonitor() {
	serverListen, err := net.Listen("tcp", ":2436")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer serverListen.Close()

	for {
		conn, err := serverListen.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("A client has been connected")
		handlers = append(handlers, conn)
	}
}

func SendOverMonitor(message string) {
	for i := 0; i < len(handlers); {
		_, err := handlers[i].Write([]byte(message))
		if err != nil {
			handlers = append(handlers[:i], handlers[i+1:]...)
			continue
		}
		i++
	}
	fmt.Println(handlers)
}
