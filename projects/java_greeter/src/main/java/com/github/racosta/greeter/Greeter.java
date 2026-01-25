package com.github.racosta.greeter;

public class Greeter {
  public String greet() {
    return greet("World");
  }

  public String greet(String name) {
    return "Hello " + name + "!";
  }
}
