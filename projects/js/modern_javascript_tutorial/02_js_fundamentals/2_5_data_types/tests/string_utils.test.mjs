import { greeter } from "../src/stringUtils.mjs";

describe("Greeter Interpolation", () => {
  test("interpolates a number correctly", () => {
    expect(greeter(1)).toBe("hello 1");
  });

  test("interpolates a string literal correctly", () => {
    expect(greeter("name")).toBe("hello name");
  });

  test("interpolates a variable value correctly", () => {
    const name = "Ilya";
    expect(greeter(name)).toBe("hello Ilya");
  });
});
