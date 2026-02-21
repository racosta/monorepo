import { createInterface } from "readline";
import { runGreeting } from "./greeter.mjs";
import { process } from "node:process";

const rl = createInterface({
  input: process.stdin,
  output: process.stdout,
});

runGreeting(rl).then(() => {
  process.exit(0);
});
