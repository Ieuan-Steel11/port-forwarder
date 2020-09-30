package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":80")
	// listening socket for getting connections

	if err != nil {
		fmt.Println("Error: Could not set up server socket")
		os.Exit(1)
	}

	for {
		connection, err := listener.Accept()
		// accepts new client connections

		if err != nil {
			fmt.Println("Error: Could not accept connection")
			continue
		}
		go handle(connection)
	}
}

func handle(connection net.Conn) {
	destination_conn, err := net.Dial("tcp", "127.0.0.1:5000")

	if err != nil {
		fmt.Println("Error: could not conenct to destination server")
	}

	defer destination_conn.Close()
	// after func returns close server connection

	go func() {
		_, err := io.Copy(destination_conn, connection)
		// copies data from client to destination

		if err != nil {
			fmt.Println("Error: could not send data")
		}
	}()

	_, err2 := io.Copy(connection, destination_conn)
	// sends data from destination to the client

	if err2 != nil {
		fmt.Println("Error: Could not get data")
	}
}
