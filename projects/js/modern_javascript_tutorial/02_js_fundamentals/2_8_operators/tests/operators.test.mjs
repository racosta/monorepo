import { add } from "../src/operators.mjs";

describe("Operators Module", () => {
  test("prefix and postfix increment", () => {
    let a = 5;
    expect(a++).toBe(5); // Postfix: returns 5, then a becomes 6
    expect(a).toBe(6); // a is now 6

    let b = 5;
    expect(++b).toBe(6); // Prefix: b becomes 6, then returns 6
    expect(b).toBe(6); // b is still 6
  });

  test("complex expression with multiple operators", () => {
    let a = 2;
    const x = 1 + (a *= 2);

    expect(a).toBe(4); // a is now 4
    expect(x).toBe(5); // a is 4, so 1 + 4 = 5
  });

  test("type conversions with operators", () => {
    expect("" + 1 + 0).toBe("10"); // "" + 1 = "1", "1" + 0 = "10"
    expect("" - 1 + 0).toBe(-1); // "" - 1 = -1, -1 + 0 = -1
    expect(true + false).toBe(1); // true = 1, false = 0, so 1 + 0 = 1
    expect(6 / "3").toBe(2); // "3" is converted to 3, so 6 / 3 = 2
    expect("2" * "3").toBe(6); // Both strings are converted to numbers, so 2 * 3 = 6
    expect(4 + 5 + "px").toBe("9px"); // (4 + 5) = 9, then "9" + "px" = "9px"
    expect("$" + 4 + 5).toBe("$45"); // "$" + 4 = "$4", then "$4" + 5 = "$45"
    expect("4" - 2).toBe(2); // "4" is converted to 4, so 4 - 2 = 2
    expect("4px" - 2).toBeNaN(); // "4px" cannot be converted to a number, so result is NaN
    expect("  -9  " + 5).toBe("  -9  5"); // String concatenation, no trimming
    expect("  -9  " - 5).toBe(-14); // "  -9  " is converted to -9, so -9 - 5 = -14
    expect(null + 1).toBe(1); // null is converted to 0, so 0 + 1 = 1
    expect(undefined + 1).toBeNaN(); // undefined cannot be converted to a number, so result is NaN
    expect(" \t \n" - 2).toBe(-2); // String with only whitespace is converted to 0, so 0 - 2 = -2
  });

  test("add function", () => {
    let a = "1";
    let b = "2";
    expect(add(a, b)).toBe("12");

    a = Number(a);
    b = Number(b);
    expect(add(a, b)).toBe(3);
  });
});
