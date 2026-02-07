const MSG = "I'm JavaScript!";

// Check if running in a browser environment
if (typeof window !== "undefined") {
  alert(MSG);
} else {
  console.log(MSG);
}
