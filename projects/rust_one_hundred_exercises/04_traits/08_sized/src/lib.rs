pub fn example() {
  // Trying to get the size of a str (or any other DST)
  // via `std::mem::size_of` will result in a compile-time error.

  // This fails to compile:
  //let _ = std::mem::size_of::<str>();

  // However, we can get the size of a reference to a str,
  // since references have a known size.
  let _ = std::mem::size_of::<&str>();
}
