package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func waitforkeyboard(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		_, err := fmt.Fprintln(conn, text)
		if err != nil {
			fmt.Println("Disconnected from server:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

func waitfornetwork(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		respone, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nDisconnected from server.")
			os.Exit(0)
		}
		fmt.Print(respone)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server successfully.")

	go waitforkeyboard(conn)

	waitfornetwork(conn)
}
