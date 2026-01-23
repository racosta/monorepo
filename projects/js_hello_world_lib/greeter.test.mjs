import { greet } from "./greeter.mjs";

test("should greet properly", () => {
  expect(greet("Alice")).toBe("Hello, Alice!");
});
