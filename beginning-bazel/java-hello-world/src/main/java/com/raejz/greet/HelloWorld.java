package com.raejz.greet;

import com.raejz.multi.IntMultiplier;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class HelloWorld {

  private static final Logger log = LoggerFactory.getLogger(HelloWorld.class);
  public static void main(String[] args) {
    log.info("Java: Hello, World!");
    IntMultiplier in = new IntMultiplier(3, 4);
    log.info("result: {}", in.getProduct());
  }
}