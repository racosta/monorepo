#include "greeter.h"
#include <iostream>

int main(int argc, char *argv[]) {
  auto greeting = greet();
  std::cout << greeting << std::endl;
  return 0;
}
