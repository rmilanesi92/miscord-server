package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// Server struct defines a server with a port number
type Server struct {
    port int
}

// Start method starts the server and listens for incoming connections
func (server Server) Start() {
    // Creating the address string using the server port
    address := fmt.Sprintf(":%d", server.port)

    // Listen for incoming TCP connections on the given port
    listener, err := net.Listen("tcp", address)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Accept incoming connection
    connection, err := listener.Accept()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Listening on port %d", server.port)

    // Ensure the connection close at the end of the scope
    defer connection.Close()

    for {
        // Create a buffer to store inputs
        buffer := make([]byte, 1024)
        
        // Reading data from the connection into the buffer
        _, err := connection.Read(buffer)
        if err != nil {

            // Break the loop if the connection is closed
            if err == io.EOF {
                break
            }
            fmt.Println("ERR ", err.Error())
            os.Exit(1)
        }

        // Write a response back to the client
        connection.Write([]byte("+OK\r\n"))
    }

}




func main() {
    // Create Server listening on port 6379
    server := Server{port: 6379}

    // Starting the server
    server.Start()
}
