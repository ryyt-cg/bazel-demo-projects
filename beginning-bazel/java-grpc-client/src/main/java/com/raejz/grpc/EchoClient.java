package com.raejz.grpc;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import transceiver.TransceiverGrpc;
import transceiver.TransceiverOuterClass.EchoRequest;
import transceiver.TransceiverOuterClass.EchoResponse;
import transmission_object.TransmissionObjectOuterClass.TransmissionObject;

public class EchoClient {
  private static final Logger log = LoggerFactory.getLogger(EchoClient.class);
  public static void main(String[] args) {
    log.info("Spinning up the Echo Client in Java...");

    try {
      final BufferedReader commandLineInput = new BufferedReader(new InputStreamReader(System.in));
      log.info("Waiting on input from the user...");
      final String inputFromUser = commandLineInput.readLine();

      if (inputFromUser != null) {
        ManagedChannel channel =
            ManagedChannelBuilder
                .forAddress("localhost", 1234)
                .usePlaintext()
                .build();
        TransceiverGrpc.TransceiverBlockingStub stub =
            TransceiverGrpc.newBlockingStub(channel);
        EchoRequest request =  EchoRequest.newBuilder()
            .setFromClient(
                TransmissionObject.newBuilder()
                    .setMessage(inputFromUser)
                    .setValue(3.145f)
                    .build())
            .build();
        EchoResponse response = stub.echo(request);
        log.info("Received Message from server: ");
        log.info(response.toString());
        channel.shutdownNow();
      }
    } catch (Exception e) {
      log.error("Error: ", e);
    }
  }
}