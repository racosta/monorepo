fn greeting(name: &str) -> String {
  format!("Hello, {name}!")
}

#[allow(dead_code)]
fn main() {
  let name = "Ryan";
  println!("{}", greeting(name));
}

#[cfg(test)]
mod test {
  use super::greeting;

  #[test]
  fn test_greeting() {
    let expected_msg = "Hello, Dolly!";
    let hello_msg = greeting("Dolly");
    assert_eq!(expected_msg, hello_msg);
  }
}
