namespace eval greetings {
  proc greet {name} {
    return "Hello $name"
  }

  package provide greetings 1.0
}
