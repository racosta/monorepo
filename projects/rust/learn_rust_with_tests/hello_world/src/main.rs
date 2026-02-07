#[allow(dead_code)]
fn main() {
  println!("Hello, world!");

  let hello_world = hello_world();
  println!("{}", hello_world);

  hello("Rusty", "");
}

pub const SPANISH_LANGUAGE: &str = "Spanish";
pub const FRENCH_LANGUAGE: &str = "French";

pub const ENGLISH_HELLO_PREFIX: &str = "Hello, ";
pub const SPANISH_HELLO_PREFIX: &str = "Hola, ";
pub const FRENCH_HELLO_PREFIX: &str = "Bonjour, ";

pub fn hello(name: &str, language: &str) -> String {
  let prefix = greeting_prefix(language);
  let final_name = name.is_empty().then_some("World").unwrap_or(&name);
  format!("{prefix}{final_name}!")
}

pub fn greeting_prefix(language: &str) -> String {
  match language {
    SPANISH_LANGUAGE => String::from(SPANISH_HELLO_PREFIX),
    FRENCH_LANGUAGE => String::from(FRENCH_HELLO_PREFIX),
    _ => String::from(ENGLISH_HELLO_PREFIX),
  }
}

pub fn hello_world() -> String {
  hello("World", "English")
}

#[cfg(test)]
mod tests {
  use super::hello;
  use super::hello_world;
  use super::FRENCH_LANGUAGE;
  use super::SPANISH_LANGUAGE;

  #[test]
  fn test_hello_world() {
    let want = String::from("Hello, World!");
    let got = hello_world();
    assert_eq!(want, got);
  }

  #[test]
  fn test_hello_to_a_person() {
    let want = String::from("Hello, Rusty!");
    let name = "Rusty";
    let got = hello(name, "");
    assert_eq!(want, got);
  }

  #[test]
  fn test_hello_empty_string() {
    let want = String::from("Hello, World!");
    let got = hello("", "");
    assert_eq!(want, got);
  }

  #[test]
  fn test_hello_in_spanish() {
    let want = String::from("Hola, Elodie!");
    let got = hello("Elodie", SPANISH_LANGUAGE);
    assert_eq!(want, got);
  }

  #[test]
  fn test_hello_in_french() {
    let want = String::from("Bonjour, Lauren!");
    let got = hello("Lauren", FRENCH_LANGUAGE);
    assert_eq!(want, got);
  }
}
