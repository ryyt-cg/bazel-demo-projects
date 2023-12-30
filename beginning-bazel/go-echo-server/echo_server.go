package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Spinning up the Echo Server in Go...")
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panicln("Unable to listen: " + err.Error())
	}

	defer listen.Close()
	connection, err := listen.Accept()
	if err != nil {
		log.Panicln("Cannot accept a connection! Error: " + err.Error())
	}

	log.Println("Receiving on a new connection")
	defer connection.Close()
	defer log.Println("Connection now closed.")
	buffer := make([]byte, 2048)

	size, err := connection.Read(buffer)
	if err != nil {
		log.Println("Cannot read from the buffer! Error: " + err.Error())
	}
	data := string(buffer[:size])

	log.Println("Received data: " + data)
	connection.Write([]byte("Echoed from Go: " + data))
}
