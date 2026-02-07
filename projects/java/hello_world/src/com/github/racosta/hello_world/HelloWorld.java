package com.github.racosta.hello_world;

import com.github.racosta.greeter.Greeter;

public class HelloWorld {
  public static void main(String[] args) {
    Greeter greeter = new Greeter();
    System.out.println(greeter.greet());
  }
}
