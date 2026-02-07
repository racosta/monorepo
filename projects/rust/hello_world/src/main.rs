extern crate greeter;

use greeter::greeter::Greeter;

fn main() -> anyhow::Result<()> {
  let hello = Greeter::new("Hello");
  println!("{}", hello.greet("world"));
  Ok(())
}
