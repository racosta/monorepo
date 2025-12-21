extern crate hello_lib;

use hello_lib::greeter::Greeter;

fn main() -> anyhow::Result<()> {
    let hello = Greeter::new("Hello");
    println!("{}", hello.greet("world"));
    Ok(())
}
