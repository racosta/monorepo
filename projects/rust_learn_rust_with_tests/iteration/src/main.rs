#[allow(dead_code)]
fn main() {}

const REPEAT_COUNT: usize = 5;

pub fn repeat(c: char) -> String {
  let s = c.to_string();
  s.repeat(REPEAT_COUNT)
}

#[cfg(test)]
mod tests {
  use super::repeat;

  #[test]
  fn test_repeat() {
    let want = String::from("aaaaa");
    let got = repeat('a');
    assert_eq!(want, got);
  }
}
