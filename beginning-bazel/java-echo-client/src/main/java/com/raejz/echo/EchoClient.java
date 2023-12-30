package com.raejz.echo;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class EchoClient {
  private static final Logger log = LoggerFactory.getLogger(EchoClient.class);

  public static void main(String[] args) {
    log.info("Spinning up the Echo Client in Java...");

    try {
      final Socket socketToServer = new Socket("localhost", 1234);
      final BufferedReader inputFromServer = new BufferedReader(
          new InputStreamReader(socketToServer.getInputStream()));
      final BufferedReader commandLineInput = new BufferedReader(
          new InputStreamReader(System.in));
      log.info("Waiting on input from the user...");
      final String inputFromUser = commandLineInput.readLine();

      if (inputFromUser != null) {
        log.info("Received by Java: " + inputFromUser);
        final PrintWriter outputToServer =
            new PrintWriter(socketToServer.getOutputStream(), true);
        outputToServer.println(inputFromUser);
        log.info(inputFromServer.readLine());
      }
      socketToServer.close();
    } catch (Exception e) {
      log.error("Error: ", e);
    }
  }
}