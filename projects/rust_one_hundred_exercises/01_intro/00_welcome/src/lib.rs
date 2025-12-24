pub fn greeting() -> &'static str {
  "I'm ready to learn Rust!"
}

#[cfg(test)]
mod tests {
  use crate::greeting;

  #[test]
  fn test_welcome() {
    assert_eq!(greeting(), "I'm ready to learn Rust!");
  }
}
