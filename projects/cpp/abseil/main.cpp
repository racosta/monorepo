#include "projects/cpp/abseil/abseil_example.h"
#include <spdlog/cfg/argv.h>
#include <spdlog/cfg/env.h>
#include <spdlog/sinks/stdout_color_sinks.h>
#include <spdlog/spdlog.h>

int main(int argc, char *argv[]) {
  // SPDLOG_LEVEL=error
  spdlog::cfg::load_env_levels();
  // ./abseil_example SPDLOG_LEVEL=error
  spdlog::cfg::load_argv_levels(argc, argv);

  std::vector<std::string> items = {"apple", "banana", "cherry"};
  std::string result = projects::cpp::abseil::JoinItems(items);
  spdlog::info("Joined string: {}", result);
  return 0;
}
