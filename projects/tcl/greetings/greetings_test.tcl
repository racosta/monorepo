package require tcltest
package require greetings

namespace import ::tcltest::*

test greet-world {
  Test: [greetings::greet "World"] == "Hello World"
} -body {
    greetings::greet "World"
} -result "Hello World"

test greet-alice {
  Test: [greetings::greet "Alice"] == "Hello Alice"
} -body {
  greetings::greet "Alice"
} -result "Hello Alice"

cleanupTests
