package com.raejz.firstbuf;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.Socket;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import transmission_object.TransmissionObjectOuterClass.TransmissionObject;

public class EchoClient {
  private static final Logger log = LoggerFactory.getLogger(EchoClient.class);

  public static void main(String args[]) {
    log.info("Spinning up the Echo Client in Java...");

    try {
      final Socket socketToServer;
      socketToServer = new Socket("localhost", 1234);

      // Note we don't need the second BufferedReader here.
      final BufferedReader commandLineInput = new BufferedReader(new InputStreamReader(System.in));
      log.info("Waiting on input from the user...");
      final String inputFromUser = commandLineInput.readLine();
      if (inputFromUser != null) {
        log.info("Received by Java: {}", inputFromUser);
        TransmissionObject transmissionObject = TransmissionObject.newBuilder()
            .setMessage(inputFromUser)
            .setValue(3.145f)
            .build();

        transmissionObject.writeTo(socketToServer.getOutputStream());
        TransmissionObject receivedObject = TransmissionObject.parseFrom(
            socketToServer.getInputStream());
        log.info("Received Message from server: {}", receivedObject);
      }
      socketToServer.close();
    } catch (IOException e) {
      log.error("Error: ", e);
    }
  }
}