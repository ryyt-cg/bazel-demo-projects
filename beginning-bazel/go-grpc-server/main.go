package main

import (
	"context"
	"fmt"
	"gitlab.com/aionx/bazel-demo-projects/beginning-bazel/first-protobuf/transceiver"
	"gitlab.com/aionx/bazel-demo-projects/beginning-bazel/first-protobuf/transmission"
	"google.golang.org/grpc"
	"log"
	"net"
)

type EchoServer struct{}

func (es *EchoServer) Echo(context context.Context, request *transceiver.EchoRequest) (*transceiver.EchoResponse, error) {
	log.Println("Message = " + (*request).FromClient.GetMessage())
	log.Println("Value = " +
		fmt.Sprintf("%f", (*request).FromClient.GetValue()))
	serverMessage := "Received from client: " +
		(*request).FromClient.GetMessage()
	serverValue := (*request).FromClient.Value * 2
	fromServer := transmission.TransmissionObject{
		Message: serverMessage,
		Value:   serverValue,
	}

	return &transceiver.EchoResponse{
		FromServer: &fromServer,
	}, nil
}

func main() {
	log.Println("Spinning up the Echo Server in Go...")
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panicln("Unable to listen: " + err.Error())
	}

	defer listen.Close()
	defer log.Println("Connection now closed.")
	grpc_server := grpc.NewServer()

	transceiver.RegisterTransceiverServer(grpc_server, &EchoServer{})

	err = grpc_server.Serve(listen)
	if err != nil {
		log.Panicln("Unable to start serving! Error: " + err.Error())
	}
}
