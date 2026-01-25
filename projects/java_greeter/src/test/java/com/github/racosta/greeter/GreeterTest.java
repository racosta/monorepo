package com.github.racosta.greeter;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

@RunWith(JUnit4.class)
public final class GreeterTest {
  @Test
  public void testDefaultGreet() {
    Greeter greeter = new Greeter();
    org.junit.Assert.assertEquals("Hello World!", greeter.greet());
  }

  @Test
  public void testGreetWithName() {
    Greeter greeter = new Greeter();
    org.junit.Assert.assertEquals("Hello Bob!", greeter.greet("Bob"));
    org.junit.Assert.assertEquals("Hello Charlie!", greeter.greet("Charlie"));
  }
}
