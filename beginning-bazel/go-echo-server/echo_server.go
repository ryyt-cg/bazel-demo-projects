package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Spinning up the Echo Server in Go...")
	listen, error := net.Listen("tcp", ":1234")
	if error != nil {
		log.Panicln("Unable to listen: " + error.Error())
	}
	defer listen.Close()
	connection, error := listen.Accept()
	if error != nil {
		log.Panicln("Cannot accept a connection! Error: " + error.Error())
	}
	log.Println("Receiving on a new connection")
	defer connection.Close()
	defer log.Println("Connection now closed.")
	buffer := make([]byte, 2048)
	size, error := connection.Read(buffer)
	if error != nil {
		log.Println("Cannot read from the buffer! Error: " + error.Error())
	}
	data := string(buffer[:size])
	log.Println("Received data: " + data)
	connection.Write([]byte("Echoed from Go: " + data))
}
