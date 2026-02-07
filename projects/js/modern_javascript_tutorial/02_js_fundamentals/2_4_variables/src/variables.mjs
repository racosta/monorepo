const NAME = "John";
const ADMIN = NAME;

// Check if running in a browser environment
if (typeof window !== "undefined") {
  alert(ADMIN);
} else {
  console.log(ADMIN);
}
