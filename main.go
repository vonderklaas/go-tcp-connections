package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// Helper
func printError(err error) {
	fmt.Println("Error:", err)
}

func startServer() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		printError(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening at http://localhost:8080")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			printError(err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			printError(err)
			return
		}

		// Process the data
		fmt.Printf("Received: %s\n", buffer[:n])
	}
}

func main() {
	// Start the server in a separate goroutine
	go startServer()

	// Wait for interrupt signal (e.g., Ctrl+C) to shut down the server
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	fmt.Println("\nShutting down...")
}
