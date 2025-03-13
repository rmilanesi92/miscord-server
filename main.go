package main

import (
	"fmt"
	"net"
    "github.com/rmilanesi92/miscord-server/resp"
    "github.com/rmilanesi92/miscord-server/command"
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
        // Create the RESP reader
        reader := resp.NewRespReader(connection)

        // Read inputs
        value := reader.Read()
       
        // Handle Command
        commandResponse := command.Handle(value) 

        // Write the RespValue response to client
        connection.Write(commandResponse.ToBytes())
    }

}




func main() {
    // Create Server listening on port 6379
    server := Server{port: 6379}

    // Starting the server
    server.Start()
}
