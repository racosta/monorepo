import { greeter } from "./stringUtils.mjs";

const name = "Ilya";

console.log(greeter(1)); // hello 1

console.log(greeter("name")); // hello name

console.log(greeter(name)); // hello Ilya
