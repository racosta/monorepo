#[allow(dead_code)]
fn main() {
  println!("{}", add(2, 2));
}

pub fn add(x: i32, y: i32) -> i32 {
  x + y
}

#[cfg(test)]
mod tests {
  use super::add;

  #[test]
  fn test_add() {
    let want = 4;
    let got = add(2, 2);
    assert_eq!(want, got);
  }

  #[test]
  fn test_variable_default_init() {
    let want: i64 = 0;
    let got = 0;
    assert_eq!(want, got);
  }

  #[test]
  fn test_mut_variable() {
    let want = 42;
    let mut got = 0;
    assert_ne!(want, got);
    got = 42;
    assert_eq!(want, got);
  }

  #[test]
  fn test_shadow_variable() {
    let int_want = 42;
    let got = 42;
    assert_eq!(int_want, got);

    let str_want = "FORTY-TWO";
    let got = "FORTY-TWO";
    assert_eq!(str_want, got);
  }
}
