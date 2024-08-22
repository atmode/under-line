package main

import (
	"bufio"
	"fmt"
	"net"
)

const (
	ClientAID = "ClientA"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("error connection to server:", err)
		return
	}
	defer conn.Close()

	//send connection message to clientB
	connectionMessage := fmt.Sprintf("hello, %s this is %s connection to you. \n", "ClientB", ClientAID)
	_, err = conn.Write([]byte(connectionMessage))
	if err != nil {
		fmt.Println("error sending message:", err)
		return
	}

	//read acknowledgment from clientB
	ackMessage, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("error reading acknowledgment:", err)
		return
	}

	fmt.Printf("[%s] received: %s", ClientAID, ackMessage)
}
