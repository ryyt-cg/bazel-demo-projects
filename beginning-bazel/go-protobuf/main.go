package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"gitlab.com/aionx/bazel-demo-projects/beginning-bazel/first-protobuf/transmission"
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
		log.Panicln(
			"Unable to read from the buffer! Error: " + err.Error())
	}
	data := buffer[:size]

	tObject := &transmission.TransmissionObject{}
	err = proto.Unmarshal(data, tObject)
	if err != nil {
		log.Panicln(
			"Unable to unmarshal the buffer! Error: " + err.Error())
	}
	log.Println("Message = " + tObject.GetMessage())
	log.Println("Value = " + fmt.Sprintf("%f", tObject.GetValue()))
	tObject.Message = "Echoed from Go: " + tObject.GetMessage()
	tObject.Value = 2 * tObject.GetValue()
	message, err := proto.Marshal(tObject)
	if err != nil {
		log.Panicln("Unable to marshal the object! Error: " + err.Error())
	}
	connection.Write(message)
}
