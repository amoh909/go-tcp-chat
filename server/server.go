package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected\n", conn.RemoteAddr())
			break
		}
		fmt.Printf("Client %s: %s", conn.RemoteAddr(), message)
		conn.Write([]byte("Echo: " + message))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Failed to listen on port:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}
		fmt.Printf("New connection from %s\n", conn.RemoteAddr())
		go handleConnection(conn)
	}
}
