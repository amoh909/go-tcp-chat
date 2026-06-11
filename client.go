package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Failed to send message:", err)
			return
		}
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Failed to receive message:", err)
			return
		}
		fmt.Print("Server echo: " + response)
	}
}
