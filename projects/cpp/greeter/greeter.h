#pragma once

#include <string>

namespace projects::cpp::greeter {

auto greet(const std::string &name = "World") -> std::string;

} // namespace projects::cpp::greeter
