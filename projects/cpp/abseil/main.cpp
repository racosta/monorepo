#include "absl/flags/flag.h"
#include "absl/flags/parse.h"
#include "projects/cpp/abseil/abseil_example.h"
#include <spdlog/cfg/argv.h>
#include <spdlog/cfg/env.h>
#include <spdlog/sinks/stdout_color_sinks.h>
#include <spdlog/spdlog.h>

// Define flags: ABSL_FLAG(type, name, default_value, description);
ABSL_FLAG(std::string, name, "World", "The name to greet");
ABSL_FLAG(int, count, 1, "Number of times to print the greeting");
ABSL_FLAG(std::string, log_level, "info",
          "Log level (trace, debug, info, warn, err, critical, off)");

int main(int argc, char *argv[]) {
  absl::ParseCommandLine(argc, argv);

  // Set log level from the flag
  std::string log_level_str = absl::GetFlag(FLAGS_log_level);
  auto level = spdlog::level::from_str(log_level_str);
  spdlog::set_level(level);

  spdlog::trace("This is a trace message");
  spdlog::debug("This is a debug message");
  spdlog::info("This is an info message");
  spdlog::warn("This is a warning message");
  spdlog::error("This is an error message");
  spdlog::critical("This is a critical message");

  std::vector<std::string> items = {"apple", "banana", "cherry"};
  std::string result = projects::cpp::abseil::JoinItems(items);
  spdlog::info("Joined string: {}", result);
  return 0;
}
