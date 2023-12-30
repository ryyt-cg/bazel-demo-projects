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
		log.Panicln(
			"Unable to read from the buffer! Error: " + error.Error())
	}
	data := buffer[:size]

	tObject := &transmission.TransmissionObject{}
	error = proto.Unmarshal(data, tObject)
	if error != nil {
		log.Panicln(
			"Unable to unmarshal the buffer! Error: " + error.Error())
	}
	log.Println("Message = " + tObject.GetMessage())
	log.Println("Value = " +
		fmt.Sprintf("%f", tObject.GetValue()))
	tObject.Message = "Echoed from Go: " +
		tObject.GetMessage()
	tObject.Value = 2 * tObject.GetValue()
	message, err := proto.Marshal(tObject)
	if err != nil {
		log.Panicln("Unable to marshal the object! Error: " + error.Error())
	}
	connection.Write(message)
}
