export const formatGreeting = (name) => `Hello, ${name}!`;

export const runGreeting = (rl, logger = console.log) => {
  return new Promise((resolve) => {
    rl.question("What is your name? ", (name) => {
      logger(formatGreeting(name));
      rl.close();
    });

    rl.on("close", () => {
      logger("Goodbye!");
      resolve();
    });
  });
};
