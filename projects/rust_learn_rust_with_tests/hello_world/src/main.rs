fn main() {
  println!("Hello, world!");

  let hello = hello_world();
  println!("{}", hello);

  let name = String::from("Rusty");
  greeting(name);
}

fn hello_world() -> String {
  let greeting = String::from("Hello, World!");
  greeting
}

fn greeting(name: String) -> String {
  let hello = String::from("Hello, ");
  let greeting = format!("{hello}{name}!");
  greeting
}

#[cfg(test)]
mod tests {
  use super::greeting;
  use super::hello_world;

  #[test]
  fn hello_world_test() {
    let want = String::from("Hello, World!");
    let result = hello_world();
    assert_eq!(want, result);
  }

  #[test]
  fn greeting_test() {
    let want = String::from("Hello, Rusty!");
    let name = String::from("Rusty");
    let result = greeting(name);
    assert_eq!(want, result);
  }
}
