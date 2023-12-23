package com.raejz.greet;

import com.raejz.multi.IntMultiplier;

public class HelloWorld {

  public static void main(String[] args) {
    System.out.println("Java: Hello, World!");
    IntMultiplier in = new IntMultiplier(3, 4);
    System.out.println(in.GetProduct());
  }
}