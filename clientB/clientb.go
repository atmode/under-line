package main

import (
	"bufio"
	"fmt"
	"net"
)

const (
	ClientBID = "ClientB"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("error starting server:", err)
		return
	}
	defer listen.Close()

	fmt.Println(ClientBID, "is listening on port 8080")

	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("error accepting connection: ", err)
		return
	}
	defer conn.Close()
	// read message from clientA
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("error reading message:", err)
		return
	}
	fmt.Printf("[%s] received: %s", ClientBID, message)

	ackMessage := fmt.Sprintf("Hello, %s connection established with %s .\n", "clientA", ClientBID)
	_, err = conn.Write([]byte(ackMessage))
	if err != nil {
		fmt.Println("error sending acknowledgment:", err)
		return
	}
	fmt.Println("acknowledgment sent to clientA")
}
