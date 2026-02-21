import { formatGreeting, runGreeting } from "../src/greeter.mjs";
import { jest } from "@jest/globals";

describe("Greeter Module", () => {
  // Test 1: The Pure Function
  test("formatGreeting creates a valid string", () => {
    const result = formatGreeting("Alice");
    expect(result).toBe("Hello, Alice!");
  });

  // Test 2: The Interaction Logic
  test("runGreeting asks for name and logs output", async () => {
    // 1. Create a Mock Logger (so we don't actually spam the console)
    const mockLogger = jest.fn();

    // 2. Create a Mock Readline interface
    // We simulate the 'question' method by immediately calling the callback
    const mockRl = {
      question: jest.fn((_query, callback) => callback("Dave")),
      close: jest.fn(),
      on: jest.fn((event, callback) => {
        if (event === "close") callback();
      }),
    };

    // 3. Execute the function
    await runGreeting(mockRl, mockLogger);

    // 4. Assertions
    // Did it ask the right question?
    expect(mockRl.question).toHaveBeenCalledWith(
      "What is your name? ",
      expect.any(Function),
    );

    // Did it log the right greeting?
    expect(mockLogger).toHaveBeenCalledWith("Hello, Dave!");

    // Did it say goodbye?
    expect(mockLogger).toHaveBeenCalledWith("Goodbye!");

    // Did it close the interface?
    expect(mockRl.close).toHaveBeenCalled();
  });
});
